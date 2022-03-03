package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SLAveInstanceMgr struct {
	*_BaseMgr
}

// SLAveInstanceMgr open func
func SLAveInstanceMgr(db *gorm.DB) *_SLAveInstanceMgr {
	if db == nil {
		panic(fmt.Errorf("SLAveInstanceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SLAveInstanceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("slave_instance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SLAveInstanceMgr) GetTableName() string {
	return "slave_instance"
}

// Reset 重置gorm会话
func (obj *_SLAveInstanceMgr) Reset() *_SLAveInstanceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SLAveInstanceMgr) Get() (result SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SLAveInstanceMgr) Gets() (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SLAveInstanceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一id
func (obj *_SLAveInstanceMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDeviceID device_id获取 设备id
func (obj *_SLAveInstanceMgr) WithDeviceID(deviceID string) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = deviceID })
}

// WithSLAveConfID slave_conf_id获取 设备配置id，决定使用什么配置
func (obj *_SLAveInstanceMgr) WithSLAveConfID(slaveConfID int) Option {
	return optionFunc(func(o *options) { o.query["slave_conf_id"] = slaveConfID })
}

// WithAddress address获取 设备地址
func (obj *_SLAveInstanceMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithEnable enable获取 启用 true 禁用false
func (obj *_SLAveInstanceMgr) WithEnable(enable bool) Option {
	return optionFunc(func(o *options) { o.query["enable"] = enable })
}

// WithStatus status获取 1 正常在线 0不在线
func (obj *_SLAveInstanceMgr) WithStatus(status int16) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SLAveInstanceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SLAveInstanceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_SLAveInstanceMgr) GetByOption(opts ...Option) (result SLAveInstance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SLAveInstanceMgr) GetByOptions(opts ...Option) (results []*SLAveInstance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SLAveInstanceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SLAveInstance, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where(options.query)
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

// GetFromID 通过id获取内容 唯一id
func (obj *_SLAveInstanceMgr) GetFromID(id int64) (result SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一id
func (obj *_SLAveInstanceMgr) GetBatchFromID(ids []int64) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备id
func (obj *_SLAveInstanceMgr) GetFromDeviceID(deviceID string) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`device_id` = ?", deviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量查找 设备id
func (obj *_SLAveInstanceMgr) GetBatchFromDeviceID(deviceIDs []string) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`device_id` IN (?)", deviceIDs).Find(&results).Error

	return
}

// GetFromSLAveConfID 通过slave_conf_id获取内容 设备配置id，决定使用什么配置
func (obj *_SLAveInstanceMgr) GetFromSLAveConfID(slaveConfID int) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`slave_conf_id` = ?", slaveConfID).Find(&results).Error

	return
}

// GetBatchFromSLAveConfID 批量查找 设备配置id，决定使用什么配置
func (obj *_SLAveInstanceMgr) GetBatchFromSLAveConfID(slaveConfIDs []int) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`slave_conf_id` IN (?)", slaveConfIDs).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 设备地址
func (obj *_SLAveInstanceMgr) GetFromAddress(address string) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 设备地址
func (obj *_SLAveInstanceMgr) GetBatchFromAddress(addresss []string) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromEnable 通过enable获取内容 启用 true 禁用false
func (obj *_SLAveInstanceMgr) GetFromEnable(enable bool) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`enable` = ?", enable).Find(&results).Error

	return
}

// GetBatchFromEnable 批量查找 启用 true 禁用false
func (obj *_SLAveInstanceMgr) GetBatchFromEnable(enables []bool) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`enable` IN (?)", enables).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 1 正常在线 0不在线
func (obj *_SLAveInstanceMgr) GetFromStatus(status int16) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 1 正常在线 0不在线
func (obj *_SLAveInstanceMgr) GetBatchFromStatus(statuss []int16) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SLAveInstanceMgr) GetFromUpdateTime(updateTime time.Time) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SLAveInstanceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SLAveInstanceMgr) GetFromCreateTime(createTime time.Time) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SLAveInstanceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SLAveInstanceMgr) FetchByPrimaryKey(id int64) (result SLAveInstance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SLAveInstance{}).Where("`id` = ?", id).First(&result).Error

	return
}
