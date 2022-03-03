package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SystemConfMgr struct {
	*_BaseMgr
}

// SystemConfMgr open func
func SystemConfMgr(db *gorm.DB) *_SystemConfMgr {
	if db == nil {
		panic(fmt.Errorf("SystemConfMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemConfMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_conf"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemConfMgr) GetTableName() string {
	return "system_conf"
}

// Reset 重置gorm会话
func (obj *_SystemConfMgr) Reset() *_SystemConfMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SystemConfMgr) Get() (result SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemConfMgr) Gets() (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemConfMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一自增主键id
func (obj *_SystemConfMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCertCa cert_ca获取 ca证书
func (obj *_SystemConfMgr) WithCertCa(certCa string) Option {
	return optionFunc(func(o *options) { o.query["cert_ca"] = certCa })
}

// WithCertKey cert_key获取 证书的key
func (obj *_SystemConfMgr) WithCertKey(certKey string) Option {
	return optionFunc(func(o *options) { o.query["cert_key"] = certKey })
}

// WithCertCert cert_cert获取 cert证书
func (obj *_SystemConfMgr) WithCertCert(certCert string) Option {
	return optionFunc(func(o *options) { o.query["cert_cert"] = certCert })
}

// WithCertName cert_name获取 证书名称
func (obj *_SystemConfMgr) WithCertName(certName string) Option {
	return optionFunc(func(o *options) { o.query["cert_name"] = certName })
}

// WithCertInsecureSkipVerify cert_insecure_skip_verify获取 跳过证书验证
func (obj *_SystemConfMgr) WithCertInsecureSkipVerify(certInsecureSkipVerify bool) Option {
	return optionFunc(func(o *options) { o.query["cert_insecure_skip_verify"] = certInsecureSkipVerify })
}

// WithCertClientAuthType cert_client_auth_type获取 授权类型
func (obj *_SystemConfMgr) WithCertClientAuthType(certClientAuthType int) Option {
	return optionFunc(func(o *options) { o.query["cert_client_auth_type"] = certClientAuthType })
}

// WithFunctionAddress function_address获取
func (obj *_SystemConfMgr) WithFunctionAddress(functionAddress string) Option {
	return optionFunc(func(o *options) { o.query["function_address"] = functionAddress })
}

// WithFunctionTimeout function_timeout获取
func (obj *_SystemConfMgr) WithFunctionTimeout(functionTimeout int) Option {
	return optionFunc(func(o *options) { o.query["function_timeout"] = functionTimeout })
}

// WithFunctionKeepalive function_keepalive获取
func (obj *_SystemConfMgr) WithFunctionKeepalive(functionKeepalive int) Option {
	return optionFunc(func(o *options) { o.query["function_keepalive"] = functionKeepalive })
}

// WithFunctionMaxIDleConns function_max_idle_conns获取
func (obj *_SystemConfMgr) WithFunctionMaxIDleConns(functionMaxIDleConns int) Option {
	return optionFunc(func(o *options) { o.query["function_max_idle_conns"] = functionMaxIDleConns })
}

// WithFunctionIDleConnTimeout function_idle_conn_timeout获取
func (obj *_SystemConfMgr) WithFunctionIDleConnTimeout(functionIDleConnTimeout int) Option {
	return optionFunc(func(o *options) { o.query["function_idle_conn_timeout"] = functionIDleConnTimeout })
}

// WithFunctionTLSHandshakeTimeout function_tls_handshake_timeout获取
func (obj *_SystemConfMgr) WithFunctionTLSHandshakeTimeout(functionTLSHandshakeTimeout int) Option {
	return optionFunc(func(o *options) { o.query["function_tls_handshake_timeout"] = functionTLSHandshakeTimeout })
}

// WithFunctionExpectContinueTimeout function_expect_continue_timeout获取
func (obj *_SystemConfMgr) WithFunctionExpectContinueTimeout(functionExpectContinueTimeout int) Option {
	return optionFunc(func(o *options) { o.query["function_expect_continue_timeout"] = functionExpectContinueTimeout })
}

// WithFunctionCa function_ca获取
func (obj *_SystemConfMgr) WithFunctionCa(functionCa string) Option {
	return optionFunc(func(o *options) { o.query["function_ca"] = functionCa })
}

// WithFunctionKey function_key获取
func (obj *_SystemConfMgr) WithFunctionKey(functionKey string) Option {
	return optionFunc(func(o *options) { o.query["function_key"] = functionKey })
}

// WithFunctionCert function_cert获取
func (obj *_SystemConfMgr) WithFunctionCert(functionCert string) Option {
	return optionFunc(func(o *options) { o.query["function_cert"] = functionCert })
}

// WithFunctionName function_name获取
func (obj *_SystemConfMgr) WithFunctionName(functionName string) Option {
	return optionFunc(func(o *options) { o.query["function_name"] = functionName })
}

// WithFunctionInsecureSkipVerify function_insecure_skip_verify获取
func (obj *_SystemConfMgr) WithFunctionInsecureSkipVerify(functionInsecureSkipVerify bool) Option {
	return optionFunc(func(o *options) { o.query["function_insecure_skip_verify"] = functionInsecureSkipVerify })
}

// WithFunctionClientAuthType function_client_auth_type获取
func (obj *_SystemConfMgr) WithFunctionClientAuthType(functionClientAuthType int) Option {
	return optionFunc(func(o *options) { o.query["function_client_auth_type"] = functionClientAuthType })
}

// WithBrokerAddress broker_address获取
func (obj *_SystemConfMgr) WithBrokerAddress(brokerAddress string) Option {
	return optionFunc(func(o *options) { o.query["broker_address"] = brokerAddress })
}

// WithBrokerUsername broker_username获取
func (obj *_SystemConfMgr) WithBrokerUsername(brokerUsername string) Option {
	return optionFunc(func(o *options) { o.query["broker_username"] = brokerUsername })
}

// WithBrokerPassword broker_password获取
func (obj *_SystemConfMgr) WithBrokerPassword(brokerPassword string) Option {
	return optionFunc(func(o *options) { o.query["broker_password"] = brokerPassword })
}

// WithBrokerClientID broker_client_id获取
func (obj *_SystemConfMgr) WithBrokerClientID(brokerClientID string) Option {
	return optionFunc(func(o *options) { o.query["broker_client_id"] = brokerClientID })
}

// WithBrokerCleanSession broker_clean_session获取
func (obj *_SystemConfMgr) WithBrokerCleanSession(brokerCleanSession bool) Option {
	return optionFunc(func(o *options) { o.query["broker_clean_session"] = brokerCleanSession })
}

// WithBrokerTimeout broker_timeout获取
func (obj *_SystemConfMgr) WithBrokerTimeout(brokerTimeout int) Option {
	return optionFunc(func(o *options) { o.query["broker_timeout"] = brokerTimeout })
}

// WithBrokerKeepalive broker_keepalive获取
func (obj *_SystemConfMgr) WithBrokerKeepalive(brokerKeepalive int) Option {
	return optionFunc(func(o *options) { o.query["broker_keepalive"] = brokerKeepalive })
}

// WithBrokerMaxReconnectInterval broker_max_reconnect_interval获取
func (obj *_SystemConfMgr) WithBrokerMaxReconnectInterval(brokerMaxReconnectInterval int) Option {
	return optionFunc(func(o *options) { o.query["broker_max_reconnect_interval"] = brokerMaxReconnectInterval })
}

// WithBrokerMaxCacheMessages broker_max_cache_messages获取
func (obj *_SystemConfMgr) WithBrokerMaxCacheMessages(brokerMaxCacheMessages int) Option {
	return optionFunc(func(o *options) { o.query["broker_max_cache_messages"] = brokerMaxCacheMessages })
}

// WithBrokerDisableAutoAck broker_disable_auto_ack获取
func (obj *_SystemConfMgr) WithBrokerDisableAutoAck(brokerDisableAutoAck bool) Option {
	return optionFunc(func(o *options) { o.query["broker_disable_auto_ack"] = brokerDisableAutoAck })
}

// WithBrokerSubscriptions broker_subscriptions获取
func (obj *_SystemConfMgr) WithBrokerSubscriptions(brokerSubscriptions string) Option {
	return optionFunc(func(o *options) { o.query["broker_subscriptions"] = brokerSubscriptions })
}

// WithBrokerCa broker_ca获取
func (obj *_SystemConfMgr) WithBrokerCa(brokerCa string) Option {
	return optionFunc(func(o *options) { o.query["broker_ca"] = brokerCa })
}

// WithBrokerKey broker_key获取
func (obj *_SystemConfMgr) WithBrokerKey(brokerKey string) Option {
	return optionFunc(func(o *options) { o.query["broker_key"] = brokerKey })
}

// WithBrokerCert broker_cert获取
func (obj *_SystemConfMgr) WithBrokerCert(brokerCert string) Option {
	return optionFunc(func(o *options) { o.query["broker_cert"] = brokerCert })
}

// WithBrokerName broker_name获取
func (obj *_SystemConfMgr) WithBrokerName(brokerName string) Option {
	return optionFunc(func(o *options) { o.query["broker_name"] = brokerName })
}

// WithBrokerInsecureSkipVerify broker_insecure_skip_verify获取
func (obj *_SystemConfMgr) WithBrokerInsecureSkipVerify(brokerInsecureSkipVerify bool) Option {
	return optionFunc(func(o *options) { o.query["broker_insecure_skip_verify"] = brokerInsecureSkipVerify })
}

// WithBrokerClientAuthType broker_client_auth_type获取
func (obj *_SystemConfMgr) WithBrokerClientAuthType(brokerClientAuthType int) Option {
	return optionFunc(func(o *options) { o.query["broker_client_auth_type"] = brokerClientAuthType })
}

// WithLoggerLevel logger_level获取
func (obj *_SystemConfMgr) WithLoggerLevel(loggerLevel string) Option {
	return optionFunc(func(o *options) { o.query["logger_level"] = loggerLevel })
}

// WithLoggerEncoding logger_encoding获取
func (obj *_SystemConfMgr) WithLoggerEncoding(loggerEncoding string) Option {
	return optionFunc(func(o *options) { o.query["logger_encoding"] = loggerEncoding })
}

// WithLoggerFilename logger_filename获取
func (obj *_SystemConfMgr) WithLoggerFilename(loggerFilename string) Option {
	return optionFunc(func(o *options) { o.query["logger_filename"] = loggerFilename })
}

// WithLoggerCompress logger_compress获取
func (obj *_SystemConfMgr) WithLoggerCompress(loggerCompress bool) Option {
	return optionFunc(func(o *options) { o.query["logger_compress"] = loggerCompress })
}

// WithLoggerMaxAge logger_max_age获取
func (obj *_SystemConfMgr) WithLoggerMaxAge(loggerMaxAge int) Option {
	return optionFunc(func(o *options) { o.query["logger_max_age"] = loggerMaxAge })
}

// WithLoggerMaxSize logger_max_size获取
func (obj *_SystemConfMgr) WithLoggerMaxSize(loggerMaxSize int) Option {
	return optionFunc(func(o *options) { o.query["logger_max_size"] = loggerMaxSize })
}

// WithLoggerMaxBackups logger_max_backups获取
func (obj *_SystemConfMgr) WithLoggerMaxBackups(loggerMaxBackups int) Option {
	return optionFunc(func(o *options) { o.query["logger_max_backups"] = loggerMaxBackups })
}

// WithLoggerEncodeTime logger_encode_time获取
func (obj *_SystemConfMgr) WithLoggerEncodeTime(loggerEncodeTime string) Option {
	return optionFunc(func(o *options) { o.query["logger_encode_time"] = loggerEncodeTime })
}

// WithLoggerEncodeLevel logger_encode_level获取
func (obj *_SystemConfMgr) WithLoggerEncodeLevel(loggerEncodeLevel string) Option {
	return optionFunc(func(o *options) { o.query["logger_encode_level"] = loggerEncodeLevel })
}

// GetByOption 功能选项模式获取
func (obj *_SystemConfMgr) GetByOption(opts ...Option) (result SystemConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemConfMgr) GetByOptions(opts ...Option) (results []*SystemConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SystemConfMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SystemConf, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 唯一自增主键id
func (obj *_SystemConfMgr) GetFromID(id int64) (result SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一自增主键id
func (obj *_SystemConfMgr) GetBatchFromID(ids []int64) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCertCa 通过cert_ca获取内容 ca证书
func (obj *_SystemConfMgr) GetFromCertCa(certCa string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_ca` = ?", certCa).Find(&results).Error

	return
}

// GetBatchFromCertCa 批量查找 ca证书
func (obj *_SystemConfMgr) GetBatchFromCertCa(certCas []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_ca` IN (?)", certCas).Find(&results).Error

	return
}

// GetFromCertKey 通过cert_key获取内容 证书的key
func (obj *_SystemConfMgr) GetFromCertKey(certKey string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_key` = ?", certKey).Find(&results).Error

	return
}

// GetBatchFromCertKey 批量查找 证书的key
func (obj *_SystemConfMgr) GetBatchFromCertKey(certKeys []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_key` IN (?)", certKeys).Find(&results).Error

	return
}

// GetFromCertCert 通过cert_cert获取内容 cert证书
func (obj *_SystemConfMgr) GetFromCertCert(certCert string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_cert` = ?", certCert).Find(&results).Error

	return
}

// GetBatchFromCertCert 批量查找 cert证书
func (obj *_SystemConfMgr) GetBatchFromCertCert(certCerts []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_cert` IN (?)", certCerts).Find(&results).Error

	return
}

// GetFromCertName 通过cert_name获取内容 证书名称
func (obj *_SystemConfMgr) GetFromCertName(certName string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_name` = ?", certName).Find(&results).Error

	return
}

// GetBatchFromCertName 批量查找 证书名称
func (obj *_SystemConfMgr) GetBatchFromCertName(certNames []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_name` IN (?)", certNames).Find(&results).Error

	return
}

// GetFromCertInsecureSkipVerify 通过cert_insecure_skip_verify获取内容 跳过证书验证
func (obj *_SystemConfMgr) GetFromCertInsecureSkipVerify(certInsecureSkipVerify bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_insecure_skip_verify` = ?", certInsecureSkipVerify).Find(&results).Error

	return
}

// GetBatchFromCertInsecureSkipVerify 批量查找 跳过证书验证
func (obj *_SystemConfMgr) GetBatchFromCertInsecureSkipVerify(certInsecureSkipVerifys []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_insecure_skip_verify` IN (?)", certInsecureSkipVerifys).Find(&results).Error

	return
}

// GetFromCertClientAuthType 通过cert_client_auth_type获取内容 授权类型
func (obj *_SystemConfMgr) GetFromCertClientAuthType(certClientAuthType int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_client_auth_type` = ?", certClientAuthType).Find(&results).Error

	return
}

// GetBatchFromCertClientAuthType 批量查找 授权类型
func (obj *_SystemConfMgr) GetBatchFromCertClientAuthType(certClientAuthTypes []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`cert_client_auth_type` IN (?)", certClientAuthTypes).Find(&results).Error

	return
}

// GetFromFunctionAddress 通过function_address获取内容
func (obj *_SystemConfMgr) GetFromFunctionAddress(functionAddress string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_address` = ?", functionAddress).Find(&results).Error

	return
}

// GetBatchFromFunctionAddress 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionAddress(functionAddresss []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_address` IN (?)", functionAddresss).Find(&results).Error

	return
}

// GetFromFunctionTimeout 通过function_timeout获取内容
func (obj *_SystemConfMgr) GetFromFunctionTimeout(functionTimeout int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_timeout` = ?", functionTimeout).Find(&results).Error

	return
}

// GetBatchFromFunctionTimeout 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionTimeout(functionTimeouts []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_timeout` IN (?)", functionTimeouts).Find(&results).Error

	return
}

// GetFromFunctionKeepalive 通过function_keepalive获取内容
func (obj *_SystemConfMgr) GetFromFunctionKeepalive(functionKeepalive int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_keepalive` = ?", functionKeepalive).Find(&results).Error

	return
}

// GetBatchFromFunctionKeepalive 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionKeepalive(functionKeepalives []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_keepalive` IN (?)", functionKeepalives).Find(&results).Error

	return
}

// GetFromFunctionMaxIDleConns 通过function_max_idle_conns获取内容
func (obj *_SystemConfMgr) GetFromFunctionMaxIDleConns(functionMaxIDleConns int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_max_idle_conns` = ?", functionMaxIDleConns).Find(&results).Error

	return
}

// GetBatchFromFunctionMaxIDleConns 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionMaxIDleConns(functionMaxIDleConnss []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_max_idle_conns` IN (?)", functionMaxIDleConnss).Find(&results).Error

	return
}

// GetFromFunctionIDleConnTimeout 通过function_idle_conn_timeout获取内容
func (obj *_SystemConfMgr) GetFromFunctionIDleConnTimeout(functionIDleConnTimeout int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_idle_conn_timeout` = ?", functionIDleConnTimeout).Find(&results).Error

	return
}

// GetBatchFromFunctionIDleConnTimeout 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionIDleConnTimeout(functionIDleConnTimeouts []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_idle_conn_timeout` IN (?)", functionIDleConnTimeouts).Find(&results).Error

	return
}

// GetFromFunctionTLSHandshakeTimeout 通过function_tls_handshake_timeout获取内容
func (obj *_SystemConfMgr) GetFromFunctionTLSHandshakeTimeout(functionTLSHandshakeTimeout int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_tls_handshake_timeout` = ?", functionTLSHandshakeTimeout).Find(&results).Error

	return
}

// GetBatchFromFunctionTLSHandshakeTimeout 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionTLSHandshakeTimeout(functionTLSHandshakeTimeouts []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_tls_handshake_timeout` IN (?)", functionTLSHandshakeTimeouts).Find(&results).Error

	return
}

// GetFromFunctionExpectContinueTimeout 通过function_expect_continue_timeout获取内容
func (obj *_SystemConfMgr) GetFromFunctionExpectContinueTimeout(functionExpectContinueTimeout int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_expect_continue_timeout` = ?", functionExpectContinueTimeout).Find(&results).Error

	return
}

// GetBatchFromFunctionExpectContinueTimeout 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionExpectContinueTimeout(functionExpectContinueTimeouts []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_expect_continue_timeout` IN (?)", functionExpectContinueTimeouts).Find(&results).Error

	return
}

// GetFromFunctionCa 通过function_ca获取内容
func (obj *_SystemConfMgr) GetFromFunctionCa(functionCa string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_ca` = ?", functionCa).Find(&results).Error

	return
}

// GetBatchFromFunctionCa 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionCa(functionCas []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_ca` IN (?)", functionCas).Find(&results).Error

	return
}

// GetFromFunctionKey 通过function_key获取内容
func (obj *_SystemConfMgr) GetFromFunctionKey(functionKey string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_key` = ?", functionKey).Find(&results).Error

	return
}

// GetBatchFromFunctionKey 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionKey(functionKeys []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_key` IN (?)", functionKeys).Find(&results).Error

	return
}

// GetFromFunctionCert 通过function_cert获取内容
func (obj *_SystemConfMgr) GetFromFunctionCert(functionCert string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_cert` = ?", functionCert).Find(&results).Error

	return
}

// GetBatchFromFunctionCert 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionCert(functionCerts []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_cert` IN (?)", functionCerts).Find(&results).Error

	return
}

// GetFromFunctionName 通过function_name获取内容
func (obj *_SystemConfMgr) GetFromFunctionName(functionName string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_name` = ?", functionName).Find(&results).Error

	return
}

// GetBatchFromFunctionName 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionName(functionNames []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_name` IN (?)", functionNames).Find(&results).Error

	return
}

// GetFromFunctionInsecureSkipVerify 通过function_insecure_skip_verify获取内容
func (obj *_SystemConfMgr) GetFromFunctionInsecureSkipVerify(functionInsecureSkipVerify bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_insecure_skip_verify` = ?", functionInsecureSkipVerify).Find(&results).Error

	return
}

// GetBatchFromFunctionInsecureSkipVerify 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionInsecureSkipVerify(functionInsecureSkipVerifys []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_insecure_skip_verify` IN (?)", functionInsecureSkipVerifys).Find(&results).Error

	return
}

// GetFromFunctionClientAuthType 通过function_client_auth_type获取内容
func (obj *_SystemConfMgr) GetFromFunctionClientAuthType(functionClientAuthType int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_client_auth_type` = ?", functionClientAuthType).Find(&results).Error

	return
}

// GetBatchFromFunctionClientAuthType 批量查找
func (obj *_SystemConfMgr) GetBatchFromFunctionClientAuthType(functionClientAuthTypes []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`function_client_auth_type` IN (?)", functionClientAuthTypes).Find(&results).Error

	return
}

// GetFromBrokerAddress 通过broker_address获取内容
func (obj *_SystemConfMgr) GetFromBrokerAddress(brokerAddress string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_address` = ?", brokerAddress).Find(&results).Error

	return
}

// GetBatchFromBrokerAddress 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerAddress(brokerAddresss []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_address` IN (?)", brokerAddresss).Find(&results).Error

	return
}

// GetFromBrokerUsername 通过broker_username获取内容
func (obj *_SystemConfMgr) GetFromBrokerUsername(brokerUsername string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_username` = ?", brokerUsername).Find(&results).Error

	return
}

// GetBatchFromBrokerUsername 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerUsername(brokerUsernames []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_username` IN (?)", brokerUsernames).Find(&results).Error

	return
}

// GetFromBrokerPassword 通过broker_password获取内容
func (obj *_SystemConfMgr) GetFromBrokerPassword(brokerPassword string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_password` = ?", brokerPassword).Find(&results).Error

	return
}

// GetBatchFromBrokerPassword 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerPassword(brokerPasswords []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_password` IN (?)", brokerPasswords).Find(&results).Error

	return
}

// GetFromBrokerClientID 通过broker_client_id获取内容
func (obj *_SystemConfMgr) GetFromBrokerClientID(brokerClientID string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_client_id` = ?", brokerClientID).Find(&results).Error

	return
}

// GetBatchFromBrokerClientID 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerClientID(brokerClientIDs []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_client_id` IN (?)", brokerClientIDs).Find(&results).Error

	return
}

// GetFromBrokerCleanSession 通过broker_clean_session获取内容
func (obj *_SystemConfMgr) GetFromBrokerCleanSession(brokerCleanSession bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_clean_session` = ?", brokerCleanSession).Find(&results).Error

	return
}

// GetBatchFromBrokerCleanSession 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerCleanSession(brokerCleanSessions []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_clean_session` IN (?)", brokerCleanSessions).Find(&results).Error

	return
}

// GetFromBrokerTimeout 通过broker_timeout获取内容
func (obj *_SystemConfMgr) GetFromBrokerTimeout(brokerTimeout int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_timeout` = ?", brokerTimeout).Find(&results).Error

	return
}

// GetBatchFromBrokerTimeout 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerTimeout(brokerTimeouts []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_timeout` IN (?)", brokerTimeouts).Find(&results).Error

	return
}

// GetFromBrokerKeepalive 通过broker_keepalive获取内容
func (obj *_SystemConfMgr) GetFromBrokerKeepalive(brokerKeepalive int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_keepalive` = ?", brokerKeepalive).Find(&results).Error

	return
}

// GetBatchFromBrokerKeepalive 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerKeepalive(brokerKeepalives []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_keepalive` IN (?)", brokerKeepalives).Find(&results).Error

	return
}

// GetFromBrokerMaxReconnectInterval 通过broker_max_reconnect_interval获取内容
func (obj *_SystemConfMgr) GetFromBrokerMaxReconnectInterval(brokerMaxReconnectInterval int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_max_reconnect_interval` = ?", brokerMaxReconnectInterval).Find(&results).Error

	return
}

// GetBatchFromBrokerMaxReconnectInterval 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerMaxReconnectInterval(brokerMaxReconnectIntervals []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_max_reconnect_interval` IN (?)", brokerMaxReconnectIntervals).Find(&results).Error

	return
}

// GetFromBrokerMaxCacheMessages 通过broker_max_cache_messages获取内容
func (obj *_SystemConfMgr) GetFromBrokerMaxCacheMessages(brokerMaxCacheMessages int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_max_cache_messages` = ?", brokerMaxCacheMessages).Find(&results).Error

	return
}

// GetBatchFromBrokerMaxCacheMessages 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerMaxCacheMessages(brokerMaxCacheMessagess []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_max_cache_messages` IN (?)", brokerMaxCacheMessagess).Find(&results).Error

	return
}

// GetFromBrokerDisableAutoAck 通过broker_disable_auto_ack获取内容
func (obj *_SystemConfMgr) GetFromBrokerDisableAutoAck(brokerDisableAutoAck bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_disable_auto_ack` = ?", brokerDisableAutoAck).Find(&results).Error

	return
}

// GetBatchFromBrokerDisableAutoAck 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerDisableAutoAck(brokerDisableAutoAcks []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_disable_auto_ack` IN (?)", brokerDisableAutoAcks).Find(&results).Error

	return
}

// GetFromBrokerSubscriptions 通过broker_subscriptions获取内容
func (obj *_SystemConfMgr) GetFromBrokerSubscriptions(brokerSubscriptions string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_subscriptions` = ?", brokerSubscriptions).Find(&results).Error

	return
}

// GetBatchFromBrokerSubscriptions 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerSubscriptions(brokerSubscriptionss []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_subscriptions` IN (?)", brokerSubscriptionss).Find(&results).Error

	return
}

// GetFromBrokerCa 通过broker_ca获取内容
func (obj *_SystemConfMgr) GetFromBrokerCa(brokerCa string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_ca` = ?", brokerCa).Find(&results).Error

	return
}

// GetBatchFromBrokerCa 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerCa(brokerCas []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_ca` IN (?)", brokerCas).Find(&results).Error

	return
}

// GetFromBrokerKey 通过broker_key获取内容
func (obj *_SystemConfMgr) GetFromBrokerKey(brokerKey string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_key` = ?", brokerKey).Find(&results).Error

	return
}

// GetBatchFromBrokerKey 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerKey(brokerKeys []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_key` IN (?)", brokerKeys).Find(&results).Error

	return
}

// GetFromBrokerCert 通过broker_cert获取内容
func (obj *_SystemConfMgr) GetFromBrokerCert(brokerCert string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_cert` = ?", brokerCert).Find(&results).Error

	return
}

// GetBatchFromBrokerCert 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerCert(brokerCerts []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_cert` IN (?)", brokerCerts).Find(&results).Error

	return
}

// GetFromBrokerName 通过broker_name获取内容
func (obj *_SystemConfMgr) GetFromBrokerName(brokerName string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_name` = ?", brokerName).Find(&results).Error

	return
}

// GetBatchFromBrokerName 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerName(brokerNames []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_name` IN (?)", brokerNames).Find(&results).Error

	return
}

// GetFromBrokerInsecureSkipVerify 通过broker_insecure_skip_verify获取内容
func (obj *_SystemConfMgr) GetFromBrokerInsecureSkipVerify(brokerInsecureSkipVerify bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_insecure_skip_verify` = ?", brokerInsecureSkipVerify).Find(&results).Error

	return
}

// GetBatchFromBrokerInsecureSkipVerify 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerInsecureSkipVerify(brokerInsecureSkipVerifys []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_insecure_skip_verify` IN (?)", brokerInsecureSkipVerifys).Find(&results).Error

	return
}

// GetFromBrokerClientAuthType 通过broker_client_auth_type获取内容
func (obj *_SystemConfMgr) GetFromBrokerClientAuthType(brokerClientAuthType int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_client_auth_type` = ?", brokerClientAuthType).Find(&results).Error

	return
}

// GetBatchFromBrokerClientAuthType 批量查找
func (obj *_SystemConfMgr) GetBatchFromBrokerClientAuthType(brokerClientAuthTypes []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`broker_client_auth_type` IN (?)", brokerClientAuthTypes).Find(&results).Error

	return
}

// GetFromLoggerLevel 通过logger_level获取内容
func (obj *_SystemConfMgr) GetFromLoggerLevel(loggerLevel string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_level` = ?", loggerLevel).Find(&results).Error

	return
}

// GetBatchFromLoggerLevel 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerLevel(loggerLevels []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_level` IN (?)", loggerLevels).Find(&results).Error

	return
}

// GetFromLoggerEncoding 通过logger_encoding获取内容
func (obj *_SystemConfMgr) GetFromLoggerEncoding(loggerEncoding string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encoding` = ?", loggerEncoding).Find(&results).Error

	return
}

// GetBatchFromLoggerEncoding 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerEncoding(loggerEncodings []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encoding` IN (?)", loggerEncodings).Find(&results).Error

	return
}

// GetFromLoggerFilename 通过logger_filename获取内容
func (obj *_SystemConfMgr) GetFromLoggerFilename(loggerFilename string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_filename` = ?", loggerFilename).Find(&results).Error

	return
}

// GetBatchFromLoggerFilename 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerFilename(loggerFilenames []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_filename` IN (?)", loggerFilenames).Find(&results).Error

	return
}

// GetFromLoggerCompress 通过logger_compress获取内容
func (obj *_SystemConfMgr) GetFromLoggerCompress(loggerCompress bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_compress` = ?", loggerCompress).Find(&results).Error

	return
}

// GetBatchFromLoggerCompress 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerCompress(loggerCompresss []bool) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_compress` IN (?)", loggerCompresss).Find(&results).Error

	return
}

// GetFromLoggerMaxAge 通过logger_max_age获取内容
func (obj *_SystemConfMgr) GetFromLoggerMaxAge(loggerMaxAge int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_age` = ?", loggerMaxAge).Find(&results).Error

	return
}

// GetBatchFromLoggerMaxAge 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerMaxAge(loggerMaxAges []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_age` IN (?)", loggerMaxAges).Find(&results).Error

	return
}

// GetFromLoggerMaxSize 通过logger_max_size获取内容
func (obj *_SystemConfMgr) GetFromLoggerMaxSize(loggerMaxSize int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_size` = ?", loggerMaxSize).Find(&results).Error

	return
}

// GetBatchFromLoggerMaxSize 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerMaxSize(loggerMaxSizes []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_size` IN (?)", loggerMaxSizes).Find(&results).Error

	return
}

// GetFromLoggerMaxBackups 通过logger_max_backups获取内容
func (obj *_SystemConfMgr) GetFromLoggerMaxBackups(loggerMaxBackups int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_backups` = ?", loggerMaxBackups).Find(&results).Error

	return
}

// GetBatchFromLoggerMaxBackups 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerMaxBackups(loggerMaxBackupss []int) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_max_backups` IN (?)", loggerMaxBackupss).Find(&results).Error

	return
}

// GetFromLoggerEncodeTime 通过logger_encode_time获取内容
func (obj *_SystemConfMgr) GetFromLoggerEncodeTime(loggerEncodeTime string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encode_time` = ?", loggerEncodeTime).Find(&results).Error

	return
}

// GetBatchFromLoggerEncodeTime 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerEncodeTime(loggerEncodeTimes []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encode_time` IN (?)", loggerEncodeTimes).Find(&results).Error

	return
}

// GetFromLoggerEncodeLevel 通过logger_encode_level获取内容
func (obj *_SystemConfMgr) GetFromLoggerEncodeLevel(loggerEncodeLevel string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encode_level` = ?", loggerEncodeLevel).Find(&results).Error

	return
}

// GetBatchFromLoggerEncodeLevel 批量查找
func (obj *_SystemConfMgr) GetBatchFromLoggerEncodeLevel(loggerEncodeLevels []string) (results []*SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`logger_encode_level` IN (?)", loggerEncodeLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemConfMgr) FetchByPrimaryKey(id int64) (result SystemConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemConf{}).Where("`id` = ?", id).First(&result).Error

	return
}
