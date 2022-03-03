package context

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/baetyl/baetyl-go/v2/context/model"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/http"
	"github.com/baetyl/baetyl-go/v2/log"
	"github.com/baetyl/baetyl-go/v2/mqtt"
	"github.com/baetyl/baetyl-go/v2/pki"
	"github.com/baetyl/baetyl-go/v2/utils"
	"gorm.io/driver/mysql"
)

var (
	ErrSystemCertInvalid  = errors.New("system certificate is invalid")
	ErrSystemCertNotFound = errors.New("system certificate is not found")
)

// Context of service
type Context interface {
	// NodeName returns node name from data.
	NodeName() string
	// AppName returns app name from data.
	AppName() string
	// AppVersion returns application version from data.
	AppVersion() string
	// ServiceName returns service name from data.
	ServiceName() string

	// SystemConfig returns the config of baetyl system from data.
	SystemConfig() *SystemConfig

	//MysqlConnUrl return database's url
	MysqlConnUrl() string

	// Log returns logger interface.
	Log() *log.Logger

	// Wait waits until exit, receiving SIGTERM and SIGINT signals.
	Wait()
	// WaitChan returns wait channel.
	WaitChan() <-chan os.Signal

	// Load returns the value stored in the map for a key, or nil if no value is present.
	// The ok result indicates whether value was found in the map.
	Load(key interface{}) (value interface{}, ok bool)
	// Store sets the value for a key.
	Store(key, value interface{})
	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	// Delete deletes the value for a key.
	Delete(key interface{})

	// CheckSystemCert checks system certificate, if certificate is not found or invalid, returns an error.
	CheckSystemCert() error
	// LoadCustomConfig loads custom config.
	// If 'files' is empty, will load config from default path,
	// else the first file path will be used to load config from.
	LoadSystemCustomConfig() (*SystemConfig, error)
	// NewFunctionHttpClient creates a new function http client.
	NewFunctionHttpClient() (*http.Client, error)
	// NewSystemBrokerClientConfig creates the system config of broker
	NewSystemBrokerClientConfig() (mqtt.ClientConfig, error)
	// NewBrokerClient creates a new broker client.
	NewBrokerClient(mqtt.ClientConfig) (*mqtt.Client, error)
	// NewSystemBrokerClient creates a new system broker client.
	NewSystemBrokerClient([]mqtt.QOSTopic) (*mqtt.Client, error)
}

type ctx struct {
	sync.Map // global cache
	log      *log.Logger
}

// NewContext creates a new context
func NewContext(mysqlConnUrl string) Context {
	if mysqlConnUrl == "" {
		mysqlConnUrl = os.Getenv(KeyMysqlConnUrl)
	}
	c := &ctx{}
	c.Store(KeyMysqlConnUrl, mysqlConnUrl)
	c.Store(KeyNodeName, os.Getenv(KeyNodeName))
	c.Store(KeyAppName, os.Getenv(KeyAppName))
	c.Store(KeyAppVersion, os.Getenv(KeyAppVersion))
	c.Store(KeySvcName, os.Getenv(KeySvcName))

	var lfs []log.Field
	if c.NodeName() != "" {
		lfs = append(lfs, log.Any("node", c.NodeName()))
	}
	if c.AppName() != "" {
		lfs = append(lfs, log.Any("app", c.AppName()))
	}
	if c.ServiceName() != "" {
		lfs = append(lfs, log.Any("service", c.ServiceName()))
	}
	c.log = log.With(lfs...)
	c.log.Info("to load config from db")

	sc, err := c.LoadSystemCustomConfig()
	if err != nil {
		c.log.Error("failed to load system config, to use default config", log.Error(err))
		utils.UnmarshalYAML(nil, sc)
	}
	// populate configuration
	// if not set in config file, to use value from env.
	// if not set in env, to use default value.
	if sc.Function.Address == "" {
		sc.Function.Address = getFunctionAddress()
	}
	if sc.Function.CA == "" {
		sc.Function.CA = sc.Certificate.CA
	}
	if sc.Function.Key == "" {
		sc.Function.Key = sc.Certificate.Key
	}
	if sc.Function.Cert == "" {
		sc.Function.Cert = sc.Certificate.Cert
	}

	if sc.Broker.Address == "" {
		sc.Broker.Address = getBrokerAddress()
	}
	// auto subscribe link topic for service if service name not nil.
	if sc.Broker.Subscriptions == nil {
		sc.Broker.Subscriptions = []mqtt.QOSTopic{}
	}
	if c.ServiceName() != "" {
		if sc.Broker.ClientID == "" {
			sc.Broker.ClientID = "baetyl-link-" + c.ServiceName()
		}
		sc.Broker.Subscriptions = append(sc.Broker.Subscriptions, mqtt.QOSTopic{QOS: 1, Topic: "$link/" + c.ServiceName()})
	}
	if sc.Broker.CA == "" {
		sc.Broker.CA = sc.Certificate.CA
	}
	if sc.Broker.Key == "" {
		sc.Broker.Key = sc.Certificate.Key
	}
	if sc.Broker.Cert == "" {
		sc.Broker.Cert = sc.Certificate.Cert
	}
	c.Store(KeySysConf, sc)

	_log, err := log.Init(sc.Logger, lfs...)
	if err != nil {
		c.log.Error("failed to init logger", log.Error(err))
	}
	c.log = _log
	c.log.Debug("context is created", log.Any("conf", sc))
	return c
}

func (c *ctx) NodeName() string {
	v, ok := c.Load(KeyNodeName)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) AppName() string {
	v, ok := c.Load(KeyAppName)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) AppVersion() string {
	v, ok := c.Load(KeyAppVersion)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) ServiceName() string {
	v, ok := c.Load(KeySvcName)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) MysqlConnUrl() string {
	v, ok := c.Load(KeyMysqlConnUrl)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) SystemConfig() *SystemConfig {
	v, ok := c.Load(KeySysConf)
	if !ok {
		return nil
	}
	return v.(*SystemConfig)
}

func (c *ctx) Log() *log.Logger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
}

func (c *ctx) WaitChan() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	return sig
}

func (c *ctx) CheckSystemCert() error {
	cfg := c.SystemConfig().Certificate
	if !utils.FileExists(cfg.CA) || !utils.FileExists(cfg.Key) || !utils.FileExists(cfg.Cert) {
		return errors.Trace(ErrSystemCertNotFound)
	}
	crt, err := ioutil.ReadFile(cfg.Cert)
	if err != nil {
		return errors.Trace(err)
	}
	info, err := pki.ParseCertificates(crt)
	if err != nil {
		return errors.Trace(err)
	}
	if len(info) != 1 || len(info[0].Subject.OrganizationalUnit) != 1 ||
		info[0].Subject.OrganizationalUnit[0] != KeyBaetyl {
		return errors.Trace(ErrSystemCertInvalid)
	}
	return nil
}

func (c *ctx) LoadSystemCustomConfig() (*SystemConfig, error) {
	//todo need to merge db inStance
	db, err := gorm.Open(mysql.Open(c.MysqlConnUrl()), &gorm.Config{})
	if err != nil {
		return nil, errors.Trace(err)
	}
	systemConfig, err := GetSystemConfig(db)
	if err != nil {
		return nil, errors.Trace(err)
	} else {
		return &systemConfig, nil
	}
}

func (c *ctx) NewFunctionHttpClient() (*http.Client, error) {
	err := c.CheckSystemCert()
	if err != nil {
		return nil, errors.Trace(err)
	}
	ops, err := c.SystemConfig().Function.ToClientOptions()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return http.NewClient(ops), nil
}

func (c *ctx) NewSystemBrokerClientConfig() (mqtt.ClientConfig, error) {
	err := c.CheckSystemCert()
	if err != nil {
		return mqtt.ClientConfig{}, errors.Trace(err)
	}
	config := c.SystemConfig().Broker

	config.Subscriptions = make([]mqtt.QOSTopic, 0)
	copy(config.Subscriptions, c.SystemConfig().Broker.Subscriptions)

	return config, nil
}

func (c *ctx) NewBrokerClient(config mqtt.ClientConfig) (*mqtt.Client, error) {
	ops, err := config.ToClientOptions()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return mqtt.NewClient(ops), nil
}

func (c *ctx) NewSystemBrokerClient(subTopics []mqtt.QOSTopic) (*mqtt.Client, error) {
	config, err := c.NewSystemBrokerClientConfig()
	if err != nil {
		return nil, err
	}
	if len(subTopics) > 0 {
		config.Subscriptions = append(config.Subscriptions, subTopics...)
	}
	client, err := c.NewBrokerClient(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetSystemConfig(db *gorm.DB) (SystemConfig, error) {
	systemConfList, systemConfListErr := model.SystemConfMgr(db).GetByOptions()
	for key, errNode := range map[string]error{
		"systemConfListErr": systemConfListErr} {
		if errNode != nil {
			return SystemConfig{}, fmt.Errorf("%s get list error,info:%s", key, errNode.Error())
		}
	}
	if len(systemConfList) <= 0 {
		return SystemConfig{}, errors.New("system config can not be empty,please config")
	}
	//todo  specific version
	latestSystemConfig := systemConfList[len(systemConfList)-1]
	SubscriptionsList := make([]mqtt.QOSTopic, 0)
	if err := json.Unmarshal([]byte(latestSystemConfig.BrokerSubscriptions), &SubscriptionsList); err != nil {
		return SystemConfig{}, fmt.Errorf("latestSystemConfig.BrokerSubscriptions is not a json formate,err:%s", err.Error())
	}
	return SystemConfig{
		Certificate: utils.Certificate{
			CA:                 latestSystemConfig.CertCa,
			Key:                latestSystemConfig.CertKey,
			Cert:               latestSystemConfig.CertCert,
			Name:               latestSystemConfig.CertName,
			InsecureSkipVerify: latestSystemConfig.CertInsecureSkipVerify,
			ClientAuthType:     tls.ClientAuthType(latestSystemConfig.CertClientAuthType),
		},
		Function: http.ClientConfig{
			Address:               latestSystemConfig.FunctionAddress,
			Timeout:               time.Second * time.Duration(latestSystemConfig.FunctionTimeout),
			KeepAlive:             time.Second * time.Duration(latestSystemConfig.FunctionKeepalive),
			MaxIdleConns:          latestSystemConfig.FunctionMaxIDleConns,
			IdleConnTimeout:       time.Second * time.Duration(latestSystemConfig.FunctionIDleConnTimeout),
			TLSHandshakeTimeout:   time.Second * time.Duration(latestSystemConfig.FunctionTLSHandshakeTimeout),
			ExpectContinueTimeout: time.Second * time.Duration(latestSystemConfig.FunctionExpectContinueTimeout),
			Certificate: utils.Certificate{
				CA:                 latestSystemConfig.FunctionCert,
				Key:                latestSystemConfig.FunctionKey,
				Cert:               latestSystemConfig.FunctionCert,
				Name:               latestSystemConfig.FunctionName,
				InsecureSkipVerify: latestSystemConfig.FunctionInsecureSkipVerify,
				ClientAuthType:     tls.ClientAuthType(latestSystemConfig.FunctionClientAuthType),
			},
		},
		Broker: mqtt.ClientConfig{
			Address:              latestSystemConfig.BrokerAddress,
			Username:             latestSystemConfig.BrokerUsername,
			Password:             latestSystemConfig.BrokerPassword,
			ClientID:             latestSystemConfig.BrokerClientID,
			CleanSession:         latestSystemConfig.BrokerCleanSession,
			Timeout:              time.Second * time.Duration(latestSystemConfig.BrokerTimeout),
			KeepAlive:            time.Second * time.Duration(latestSystemConfig.BrokerKeepalive),
			MaxReconnectInterval: time.Second * time.Duration(latestSystemConfig.BrokerMaxReconnectInterval),
			MaxCacheMessages:     latestSystemConfig.BrokerMaxCacheMessages,
			DisableAutoAck:       latestSystemConfig.BrokerDisableAutoAck,
			Subscriptions:        SubscriptionsList,
			Certificate: utils.Certificate{
				CA:                 latestSystemConfig.BrokerCa,
				Key:                latestSystemConfig.BrokerKey,
				Cert:               latestSystemConfig.BrokerCert,
				Name:               latestSystemConfig.BrokerName,
				InsecureSkipVerify: latestSystemConfig.BrokerInsecureSkipVerify,
				ClientAuthType:     tls.ClientAuthType(latestSystemConfig.BrokerClientAuthType),
			},
		},
		Logger: log.Config{
			Level:       latestSystemConfig.LoggerLevel,
			Encoding:    latestSystemConfig.LoggerEncoding,
			Filename:    latestSystemConfig.LoggerFilename,
			Compress:    latestSystemConfig.LoggerCompress,
			MaxAge:      latestSystemConfig.LoggerMaxAge,
			MaxSize:     latestSystemConfig.LoggerMaxSize,
			MaxBackups:  latestSystemConfig.LoggerMaxBackups,
			EncodeTime:  latestSystemConfig.LoggerEncodeTime,
			EncodeLevel: latestSystemConfig.LoggerEncodeLevel,
		},
	}, nil
}
