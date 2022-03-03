package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MapConfigMgr struct {
	*_BaseMgr
}

// MapConfigMgr open func
func MapConfigMgr(db *gorm.DB) *_MapConfigMgr {
	if db == nil {
		panic(fmt.Errorf("MapConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MapConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("map_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MapConfigMgr) GetTableName() string {
	return "map_config"
}

// Reset 重置gorm会话
func (obj *_MapConfigMgr) Reset() *_MapConfigMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MapConfigMgr) Get() (result MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MapConfigMgr) Gets() (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MapConfigMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一自增主键
func (obj *_MapConfigMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithJobID job_id获取 定时任务id
func (obj *_MapConfigMgr) WithJobID(jobID int) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithFunction function获取
func (obj *_MapConfigMgr) WithFunction(function int) Option {
	return optionFunc(func(o *options) { o.query["function"] = function })
}

// WithAddress address获取
func (obj *_MapConfigMgr) WithAddress(address int) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithQuantity quantity获取
func (obj *_MapConfigMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithFieldName field_name获取
func (obj *_MapConfigMgr) WithFieldName(fieldName string) Option {
	return optionFunc(func(o *options) { o.query["field_name"] = fieldName })
}

// WithFieldType field_type获取
func (obj *_MapConfigMgr) WithFieldType(fieldType string) Option {
	return optionFunc(func(o *options) { o.query["field_type"] = fieldType })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MapConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_MapConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_MapConfigMgr) GetByOption(opts ...Option) (result MapConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MapConfigMgr) GetByOptions(opts ...Option) (results []*MapConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MapConfigMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MapConfig, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where(options.query)
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

// GetFromID 通过id获取内容 唯一自增主键
func (obj *_MapConfigMgr) GetFromID(id int64) (result MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一自增主键
func (obj *_MapConfigMgr) GetBatchFromID(ids []int64) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromJobID 通过job_id获取内容 定时任务id
func (obj *_MapConfigMgr) GetFromJobID(jobID int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`job_id` = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量查找 定时任务id
func (obj *_MapConfigMgr) GetBatchFromJobID(jobIDs []int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`job_id` IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromFunction 通过function获取内容
func (obj *_MapConfigMgr) GetFromFunction(function int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`function` = ?", function).Find(&results).Error

	return
}

// GetBatchFromFunction 批量查找
func (obj *_MapConfigMgr) GetBatchFromFunction(functions []int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`function` IN (?)", functions).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容
func (obj *_MapConfigMgr) GetFromAddress(address int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找
func (obj *_MapConfigMgr) GetBatchFromAddress(addresss []int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_MapConfigMgr) GetFromQuantity(quantity int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找
func (obj *_MapConfigMgr) GetBatchFromQuantity(quantitys []int) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromFieldName 通过field_name获取内容
func (obj *_MapConfigMgr) GetFromFieldName(fieldName string) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`field_name` = ?", fieldName).Find(&results).Error

	return
}

// GetBatchFromFieldName 批量查找
func (obj *_MapConfigMgr) GetBatchFromFieldName(fieldNames []string) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`field_name` IN (?)", fieldNames).Find(&results).Error

	return
}

// GetFromFieldType 通过field_type获取内容
func (obj *_MapConfigMgr) GetFromFieldType(fieldType string) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`field_type` = ?", fieldType).Find(&results).Error

	return
}

// GetBatchFromFieldType 批量查找
func (obj *_MapConfigMgr) GetBatchFromFieldType(fieldTypes []string) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`field_type` IN (?)", fieldTypes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MapConfigMgr) GetFromCreateTime(createTime time.Time) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MapConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_MapConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_MapConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MapConfigMgr) FetchByPrimaryKey(id int64) (result MapConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MapConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}
