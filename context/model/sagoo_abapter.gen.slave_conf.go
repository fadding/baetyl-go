package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SLAveConfMgr struct {
	*_BaseMgr
}

// SLAveConfMgr open func
func SLAveConfMgr(db *gorm.DB) *_SLAveConfMgr {
	if db == nil {
		panic(fmt.Errorf("SLAveConfMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SLAveConfMgr{_BaseMgr: &_BaseMgr{DB: db.Table("slave_conf"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SLAveConfMgr) GetTableName() string {
	return "slave_conf"
}

// Reset 重置gorm会话
func (obj *_SLAveConfMgr) Reset() *_SLAveConfMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SLAveConfMgr) Get() (result SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SLAveConfMgr) Gets() (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SLAveConfMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SLAveConfMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMode mode获取 tcp和rtu二选一
func (obj *_SLAveConfMgr) WithMode(mode string) Option {
	return optionFunc(func(o *options) { o.query["mode"] = mode })
}

// WithAddress address获取 设备地址，默认为/dev/ttyS0
func (obj *_SLAveConfMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithTimeout timeout获取 读写超时时间，读写一致的，默认10s
func (obj *_SLAveConfMgr) WithTimeout(timeout int) Option {
	return optionFunc(func(o *options) { o.query["timeout"] = timeout })
}

// WithIDletimeout idletimeout获取 空闲多长时间关闭连接，默认一分钟
func (obj *_SLAveConfMgr) WithIDletimeout(idletimeout int) Option {
	return optionFunc(func(o *options) { o.query["idletimeout"] = idletimeout })
}

// WithBaudrate baudrate获取 波特率，默认为19200，仅仅在rtc模式下才有用
func (obj *_SLAveConfMgr) WithBaudrate(baudrate int) Option {
	return optionFunc(func(o *options) { o.query["baudrate"] = baudrate })
}

// WithDatabits databits获取 数据位，5, 6, 7 or 8 (default 8)
func (obj *_SLAveConfMgr) WithDatabits(databits int) Option {
	return optionFunc(func(o *options) { o.query["databits"] = databits })
}

// WithStopbits stopbits获取 停止位置，1或者2，默认为1
func (obj *_SLAveConfMgr) WithStopbits(stopbits int16) Option {
	return optionFunc(func(o *options) { o.query["stopbits"] = stopbits })
}

// WithParity parity获取 奇偶校验，E|N|O 三个其中一个，默认为E
func (obj *_SLAveConfMgr) WithParity(parity string) Option {
	return optionFunc(func(o *options) { o.query["parity"] = parity })
}

// WithRs485ID rs485_id获取 rs485配置id
func (obj *_SLAveConfMgr) WithRs485ID(rs485ID string) Option {
	return optionFunc(func(o *options) { o.query["rs485_id"] = rs485ID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SLAveConfMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SLAveConfMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithIsDefault is_default获取 是否是默认的配置
func (obj *_SLAveConfMgr) WithIsDefault(isDefault bool) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// GetByOption 功能选项模式获取
func (obj *_SLAveConfMgr) GetByOption(opts ...Option) (result SLAveConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SLAveConfMgr) GetByOptions(opts ...Option) (results []*SLAveConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SLAveConfMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SLAveConf, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_SLAveConfMgr) GetFromID(id int64) (result SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SLAveConfMgr) GetBatchFromID(ids []int64) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMode 通过mode获取内容 tcp和rtu二选一
func (obj *_SLAveConfMgr) GetFromMode(mode string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`mode` = ?", mode).Find(&results).Error

	return
}

// GetBatchFromMode 批量查找 tcp和rtu二选一
func (obj *_SLAveConfMgr) GetBatchFromMode(modes []string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`mode` IN (?)", modes).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 设备地址，默认为/dev/ttyS0
func (obj *_SLAveConfMgr) GetFromAddress(address string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 设备地址，默认为/dev/ttyS0
func (obj *_SLAveConfMgr) GetBatchFromAddress(addresss []string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromTimeout 通过timeout获取内容 读写超时时间，读写一致的，默认10s
func (obj *_SLAveConfMgr) GetFromTimeout(timeout int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`timeout` = ?", timeout).Find(&results).Error

	return
}

// GetBatchFromTimeout 批量查找 读写超时时间，读写一致的，默认10s
func (obj *_SLAveConfMgr) GetBatchFromTimeout(timeouts []int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`timeout` IN (?)", timeouts).Find(&results).Error

	return
}

// GetFromIDletimeout 通过idletimeout获取内容 空闲多长时间关闭连接，默认一分钟
func (obj *_SLAveConfMgr) GetFromIDletimeout(idletimeout int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`idletimeout` = ?", idletimeout).Find(&results).Error

	return
}

// GetBatchFromIDletimeout 批量查找 空闲多长时间关闭连接，默认一分钟
func (obj *_SLAveConfMgr) GetBatchFromIDletimeout(idletimeouts []int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`idletimeout` IN (?)", idletimeouts).Find(&results).Error

	return
}

// GetFromBaudrate 通过baudrate获取内容 波特率，默认为19200，仅仅在rtc模式下才有用
func (obj *_SLAveConfMgr) GetFromBaudrate(baudrate int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`baudrate` = ?", baudrate).Find(&results).Error

	return
}

// GetBatchFromBaudrate 批量查找 波特率，默认为19200，仅仅在rtc模式下才有用
func (obj *_SLAveConfMgr) GetBatchFromBaudrate(baudrates []int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`baudrate` IN (?)", baudrates).Find(&results).Error

	return
}

// GetFromDatabits 通过databits获取内容 数据位，5, 6, 7 or 8 (default 8)
func (obj *_SLAveConfMgr) GetFromDatabits(databits int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`databits` = ?", databits).Find(&results).Error

	return
}

// GetBatchFromDatabits 批量查找 数据位，5, 6, 7 or 8 (default 8)
func (obj *_SLAveConfMgr) GetBatchFromDatabits(databitss []int) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`databits` IN (?)", databitss).Find(&results).Error

	return
}

// GetFromStopbits 通过stopbits获取内容 停止位置，1或者2，默认为1
func (obj *_SLAveConfMgr) GetFromStopbits(stopbits int16) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`stopbits` = ?", stopbits).Find(&results).Error

	return
}

// GetBatchFromStopbits 批量查找 停止位置，1或者2，默认为1
func (obj *_SLAveConfMgr) GetBatchFromStopbits(stopbitss []int16) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`stopbits` IN (?)", stopbitss).Find(&results).Error

	return
}

// GetFromParity 通过parity获取内容 奇偶校验，E|N|O 三个其中一个，默认为E
func (obj *_SLAveConfMgr) GetFromParity(parity string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`parity` = ?", parity).Find(&results).Error

	return
}

// GetBatchFromParity 批量查找 奇偶校验，E|N|O 三个其中一个，默认为E
func (obj *_SLAveConfMgr) GetBatchFromParity(paritys []string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`parity` IN (?)", paritys).Find(&results).Error

	return
}

// GetFromRs485ID 通过rs485_id获取内容 rs485配置id
func (obj *_SLAveConfMgr) GetFromRs485ID(rs485ID string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`rs485_id` = ?", rs485ID).Find(&results).Error

	return
}

// GetBatchFromRs485ID 批量查找 rs485配置id
func (obj *_SLAveConfMgr) GetBatchFromRs485ID(rs485IDs []string) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`rs485_id` IN (?)", rs485IDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SLAveConfMgr) GetFromCreateTime(createTime time.Time) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SLAveConfMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SLAveConfMgr) GetFromUpdateTime(updateTime time.Time) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SLAveConfMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否是默认的配置
func (obj *_SLAveConfMgr) GetFromIsDefault(isDefault bool) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否是默认的配置
func (obj *_SLAveConfMgr) GetBatchFromIsDefault(isDefaults []bool) (results []*SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SLAveConfMgr) FetchByPrimaryKey(id int64) (result SLAveConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveConf{}).Where("`id` = ?", id).First(&result).Error

	return
}
