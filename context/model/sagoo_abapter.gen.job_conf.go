package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _JobConfMgr struct {
	*_BaseMgr
}

// JobConfMgr open func
func JobConfMgr(db *gorm.DB) *_JobConfMgr {
	if db == nil {
		panic(fmt.Errorf("JobConfMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_JobConfMgr{_BaseMgr: &_BaseMgr{DB: db.Table("job_conf"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_JobConfMgr) GetTableName() string {
	return "job_conf"
}

// Reset 重置gorm会话
func (obj *_JobConfMgr) Reset() *_JobConfMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_JobConfMgr) Get() (result JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_JobConfMgr) Gets() (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_JobConfMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(JobConf{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一自增主键id
func (obj *_JobConfMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSLAveID slave_id获取 slave的id
func (obj *_JobConfMgr) WithSLAveID(slaveID int) Option {
	return optionFunc(func(o *options) { o.query["slave_id"] = slaveID })
}

// WithInterval interval获取 内部循环时间，默认为5s
func (obj *_JobConfMgr) WithInterval(interval int) Option {
	return optionFunc(func(o *options) { o.query["interval"] = interval })
}

// WithEncoding encoding获取 二进制或者json默认为json binary|json
func (obj *_JobConfMgr) WithEncoding(encoding string) Option {
	return optionFunc(func(o *options) { o.query["encoding"] = encoding })
}

// WithTimeName time_name获取 时间格式的名称
func (obj *_JobConfMgr) WithTimeName(timeName string) Option {
	return optionFunc(func(o *options) { o.query["time_name"] = timeName })
}

// WithTimeType time_type获取 时间类型
func (obj *_JobConfMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithTimeFormat time_format获取 时间格式
func (obj *_JobConfMgr) WithTimeFormat(timeFormat string) Option {
	return optionFunc(func(o *options) { o.query["time_format"] = timeFormat })
}

// WithTimePrecision time_precision获取 s或者ns
func (obj *_JobConfMgr) WithTimePrecision(timePrecision string) Option {
	return optionFunc(func(o *options) { o.query["time_precision"] = timePrecision })
}

// WithPublishQos publish_qos获取 消息的qos
func (obj *_JobConfMgr) WithPublishQos(publishQos int) Option {
	return optionFunc(func(o *options) { o.query["publish_qos"] = publishQos })
}

// WithPublicTopic public_topic获取 发布消息的topic
func (obj *_JobConfMgr) WithPublicTopic(publicTopic string) Option {
	return optionFunc(func(o *options) { o.query["public_topic"] = publicTopic })
}

// WithCreateTime create_time获取 创建时间
func (obj *_JobConfMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_JobConfMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_JobConfMgr) GetByOption(opts ...Option) (result JobConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_JobConfMgr) GetByOptions(opts ...Option) (results []*JobConf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_JobConfMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]JobConf, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where(options.query)
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
func (obj *_JobConfMgr) GetFromID(id int64) (result JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一自增主键id
func (obj *_JobConfMgr) GetBatchFromID(ids []int64) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSLAveID 通过slave_id获取内容 slave的id
func (obj *_JobConfMgr) GetFromSLAveID(slaveID int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`slave_id` = ?", slaveID).Find(&results).Error

	return
}

// GetBatchFromSLAveID 批量查找 slave的id
func (obj *_JobConfMgr) GetBatchFromSLAveID(slaveIDs []int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`slave_id` IN (?)", slaveIDs).Find(&results).Error

	return
}

// GetFromInterval 通过interval获取内容 内部循环时间，默认为5s
func (obj *_JobConfMgr) GetFromInterval(interval int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`interval` = ?", interval).Find(&results).Error

	return
}

// GetBatchFromInterval 批量查找 内部循环时间，默认为5s
func (obj *_JobConfMgr) GetBatchFromInterval(intervals []int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`interval` IN (?)", intervals).Find(&results).Error

	return
}

// GetFromEncoding 通过encoding获取内容 二进制或者json默认为json binary|json
func (obj *_JobConfMgr) GetFromEncoding(encoding string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`encoding` = ?", encoding).Find(&results).Error

	return
}

// GetBatchFromEncoding 批量查找 二进制或者json默认为json binary|json
func (obj *_JobConfMgr) GetBatchFromEncoding(encodings []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`encoding` IN (?)", encodings).Find(&results).Error

	return
}

// GetFromTimeName 通过time_name获取内容 时间格式的名称
func (obj *_JobConfMgr) GetFromTimeName(timeName string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_name` = ?", timeName).Find(&results).Error

	return
}

// GetBatchFromTimeName 批量查找 时间格式的名称
func (obj *_JobConfMgr) GetBatchFromTimeName(timeNames []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_name` IN (?)", timeNames).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型
func (obj *_JobConfMgr) GetFromTimeType(timeType string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_type` = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量查找 时间类型
func (obj *_JobConfMgr) GetBatchFromTimeType(timeTypes []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_type` IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromTimeFormat 通过time_format获取内容 时间格式
func (obj *_JobConfMgr) GetFromTimeFormat(timeFormat string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_format` = ?", timeFormat).Find(&results).Error

	return
}

// GetBatchFromTimeFormat 批量查找 时间格式
func (obj *_JobConfMgr) GetBatchFromTimeFormat(timeFormats []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_format` IN (?)", timeFormats).Find(&results).Error

	return
}

// GetFromTimePrecision 通过time_precision获取内容 s或者ns
func (obj *_JobConfMgr) GetFromTimePrecision(timePrecision string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_precision` = ?", timePrecision).Find(&results).Error

	return
}

// GetBatchFromTimePrecision 批量查找 s或者ns
func (obj *_JobConfMgr) GetBatchFromTimePrecision(timePrecisions []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`time_precision` IN (?)", timePrecisions).Find(&results).Error

	return
}

// GetFromPublishQos 通过publish_qos获取内容 消息的qos
func (obj *_JobConfMgr) GetFromPublishQos(publishQos int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`publish_qos` = ?", publishQos).Find(&results).Error

	return
}

// GetBatchFromPublishQos 批量查找 消息的qos
func (obj *_JobConfMgr) GetBatchFromPublishQos(publishQoss []int) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`publish_qos` IN (?)", publishQoss).Find(&results).Error

	return
}

// GetFromPublicTopic 通过public_topic获取内容 发布消息的topic
func (obj *_JobConfMgr) GetFromPublicTopic(publicTopic string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`public_topic` = ?", publicTopic).Find(&results).Error

	return
}

// GetBatchFromPublicTopic 批量查找 发布消息的topic
func (obj *_JobConfMgr) GetBatchFromPublicTopic(publicTopics []string) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`public_topic` IN (?)", publicTopics).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_JobConfMgr) GetFromCreateTime(createTime time.Time) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_JobConfMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_JobConfMgr) GetFromUpdateTime(updateTime time.Time) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_JobConfMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_JobConfMgr) FetchByPrimaryKey(id int64) (result JobConf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(JobConf{}).Where("`id` = ?", id).First(&result).Error

	return
}
