package model

import (
	"time"
)

// JobConf 定时任务配置
type JobConf struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"`              // 唯一自增主键id
	SLAveID       int       `gorm:"column:slave_id" json:"slaveId"`             // slave的id
	Interval      int       `gorm:"column:interval" json:"interval"`            // 内部循环时间，默认为5s
	Encoding      string    `gorm:"column:encoding" json:"encoding"`            // 二进制或者json默认为json binary|json
	TimeName      string    `gorm:"column:time_name" json:"timeName"`           // 时间格式的名称
	TimeType      string    `gorm:"column:time_type" json:"timeType"`           // 时间类型
	TimeFormat    string    `gorm:"column:time_format" json:"timeFormat"`       // 时间格式
	TimePrecision string    `gorm:"column:time_precision" json:"timePrecision"` // s或者ns
	PublishQos    int       `gorm:"column:publish_qos" json:"publishQos"`       // 消息的qos
	PublicTopic   string    `gorm:"column:public_topic" json:"publicTopic"`     // 发布消息的topic
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *JobConf) TableName() string {
	return "job_conf"
}

// MapConfig [...]
type MapConfig struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"-"` // 唯一自增主键
	JobID      int       `gorm:"column:job_id" json:"jobId"`    // 定时任务id
	Function   int       `gorm:"column:function" json:"function"`
	Address    int       `gorm:"column:address" json:"address"`
	Quantity   int       `gorm:"column:quantity" json:"quantity"`
	FieldName  string    `gorm:"column:field_name" json:"fieldName"`
	FieldType  string    `gorm:"column:field_type" json:"fieldType"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *MapConfig) TableName() string {
	return "map_config"
}

// Rs485Conf rs485配置
type Rs485Conf struct {
	ID                 int64     `gorm:"primaryKey;column:id" json:"-"`                          // 唯一自增主键
	SLAveConfID        int       `gorm:"column:slave_conf_id" json:"slaveConfId"`                // slave配置id
	Rs485ConfID        string    `gorm:"column:rs485_conf_id" json:"rs485ConfId"`                // rs485配置id
	Enabled            bool      `gorm:"column:enabled" json:"enabled"`                          // 配置是否开启，默认是开启的
	DelayRtsBeforeSend int       `gorm:"column:delay_rts_before_send" json:"delayRtsBeforeSend"` // 发送前rts，rts(Residual Time Stamp)
	DelayRtsAfterSend  int       `gorm:"column:delay_rts_after_send" json:"delayRtsAfterSend"`   // 发送后rts,rts(Residual Time Stamp)
	RtsHighDuringSend  bool      `gorm:"column:rts_high_during_send" json:"rtsHighDuringSend"`   // 发送过程中是否设置rts为high
	RtsHighAfterSend   bool      `gorm:"column:rts_high_after_send" json:"rtsHighAfterSend"`     // 发送后是否设置rts为high
	RxDuringTx         bool      `gorm:"column:rx_during_tx" json:"rxDuringTx"`                  // 在tx过程中的rx
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                   // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Rs485Conf) TableName() string {
	return "rs485_conf"
}

// SLAveConf 节点配置表
type SLAveConf struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"`
	Mode        string    `gorm:"column:mode" json:"mode"`               // tcp和rtu二选一
	Address     string    `gorm:"column:address" json:"address"`         // 设备地址，默认为/dev/ttyS0
	Timeout     int       `gorm:"column:timeout" json:"timeout"`         // 读写超时时间，读写一致的，默认10s
	IDletimeout int       `gorm:"column:idletimeout" json:"idletimeout"` // 空闲多长时间关闭连接，默认一分钟
	Baudrate    int       `gorm:"column:baudrate" json:"baudrate"`       // 波特率，默认为19200，仅仅在rtc模式下才有用
	Databits    int       `gorm:"column:databits" json:"databits"`       // 数据位，5, 6, 7 or 8 (default 8)
	Stopbits    int16     `gorm:"column:stopbits" json:"stopbits"`       // 停止位置，1或者2，默认为1
	Parity      string    `gorm:"column:parity" json:"parity"`           // 奇偶校验，E|N|O 三个其中一个，默认为E
	Rs485ID     string    `gorm:"column:rs485_id" json:"rs485Id"`        // rs485配置id
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`  // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`  // 更新时间
	IsDefault   bool      `gorm:"column:is_default" json:"isDefault"`    // 是否是默认的配置
}

// TableName get sql table name.获取数据库表名
func (m *SLAveConf) TableName() string {
	return "slave_conf"
}

// SLAveInstance 设备实例
type SLAveInstance struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"`           // 唯一id
	DeviceID    string    `gorm:"column:device_id" json:"deviceId"`        // 设备id
	SLAveConfID int       `gorm:"column:slave_conf_id" json:"slaveConfId"` // 设备配置id，决定使用什么配置
	Address     string    `gorm:"column:address" json:"address"`           // 设备地址
	Enable      bool      `gorm:"column:enable" json:"enable"`             // 启用 true 禁用false
	Status      int16     `gorm:"column:status" json:"status"`             // 1 正常在线 0不在线
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`    // 更新时间
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`    // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *SLAveInstance) TableName() string {
	return "slave_instance"
}

// SystemConf 系统参数配置
type SystemConf struct {
	ID                            int64  `gorm:"primaryKey;column:id" json:"-"`                                  // 唯一自增主键id
	CertCa                        string `gorm:"column:cert_ca" json:"certCa"`                                   // ca证书
	CertKey                       string `gorm:"column:cert_key" json:"certKey"`                                 // 证书的key
	CertCert                      string `gorm:"column:cert_cert" json:"certCert"`                               // cert证书
	CertName                      string `gorm:"column:cert_name" json:"certName"`                               // 证书名称
	CertInsecureSkipVerify        bool   `gorm:"column:cert_insecure_skip_verify" json:"certInsecureSkipVerify"` // 跳过证书验证
	CertClientAuthType            int    `gorm:"column:cert_client_auth_type" json:"certClientAuthType"`         // 授权类型
	FunctionAddress               string `gorm:"column:function_address" json:"functionAddress"`
	FunctionTimeout               int    `gorm:"column:function_timeout" json:"functionTimeout"`
	FunctionKeepalive             int    `gorm:"column:function_keepalive" json:"functionKeepalive"`
	FunctionMaxIDleConns          int    `gorm:"column:function_max_idle_conns" json:"functionMaxIdleConns"`
	FunctionIDleConnTimeout       int    `gorm:"column:function_idle_conn_timeout" json:"functionIdleConnTimeout"`
	FunctionTLSHandshakeTimeout   int    `gorm:"column:function_tls_handshake_timeout" json:"functionTlsHandshakeTimeout"`
	FunctionExpectContinueTimeout int    `gorm:"column:function_expect_continue_timeout" json:"functionExpectContinueTimeout"`
	FunctionCa                    string `gorm:"column:function_ca" json:"functionCa"`
	FunctionKey                   string `gorm:"column:function_key" json:"functionKey"`
	FunctionCert                  string `gorm:"column:function_cert" json:"functionCert"`
	FunctionName                  string `gorm:"column:function_name" json:"functionName"`
	FunctionInsecureSkipVerify    bool   `gorm:"column:function_insecure_skip_verify" json:"functionInsecureSkipVerify"`
	FunctionClientAuthType        int    `gorm:"column:function_client_auth_type" json:"functionClientAuthType"`
	BrokerAddress                 string `gorm:"column:broker_address" json:"brokerAddress"`
	BrokerUsername                string `gorm:"column:broker_username" json:"brokerUsername"`
	BrokerPassword                string `gorm:"column:broker_password" json:"brokerPassword"`
	BrokerClientID                string `gorm:"column:broker_client_id" json:"brokerClientId"`
	BrokerCleanSession            bool   `gorm:"column:broker_clean_session" json:"brokerCleanSession"`
	BrokerTimeout                 int    `gorm:"column:broker_timeout" json:"brokerTimeout"`
	BrokerKeepalive               int    `gorm:"column:broker_keepalive" json:"brokerKeepalive"`
	BrokerMaxReconnectInterval    int    `gorm:"column:broker_max_reconnect_interval" json:"brokerMaxReconnectInterval"`
	BrokerMaxCacheMessages        int    `gorm:"column:broker_max_cache_messages" json:"brokerMaxCacheMessages"`
	BrokerDisableAutoAck          bool   `gorm:"column:broker_disable_auto_ack" json:"brokerDisableAutoAck"`
	BrokerSubscriptions           string `gorm:"column:broker_subscriptions" json:"brokerSubscriptions"`
	BrokerCa                      string `gorm:"column:broker_ca" json:"brokerCa"`
	BrokerKey                     string `gorm:"column:broker_key" json:"brokerKey"`
	BrokerCert                    string `gorm:"column:broker_cert" json:"brokerCert"`
	BrokerName                    string `gorm:"column:broker_name" json:"brokerName"`
	BrokerInsecureSkipVerify      bool   `gorm:"column:broker_insecure_skip_verify" json:"brokerInsecureSkipVerify"`
	BrokerClientAuthType          int    `gorm:"column:broker_client_auth_type" json:"brokerClientAuthType"`
	LoggerLevel                   string `gorm:"column:logger_level" json:"loggerLevel"`
	LoggerEncoding                string `gorm:"column:logger_encoding" json:"loggerEncoding"`
	LoggerFilename                string `gorm:"column:logger_filename" json:"loggerFilename"`
	LoggerCompress                bool   `gorm:"column:logger_compress" json:"loggerCompress"`
	LoggerMaxAge                  int    `gorm:"column:logger_max_age" json:"loggerMaxAge"`
	LoggerMaxSize                 int    `gorm:"column:logger_max_size" json:"loggerMaxSize"`
	LoggerMaxBackups              int    `gorm:"column:logger_max_backups" json:"loggerMaxBackups"`
	LoggerEncodeTime              string `gorm:"column:logger_encode_time" json:"loggerEncodeTime"`
	LoggerEncodeLevel             string `gorm:"column:logger_encode_level" json:"loggerEncodeLevel"`
}

// TableName get sql table name.获取数据库表名
func (m *SystemConf) TableName() string {
	return "system_conf"
}
