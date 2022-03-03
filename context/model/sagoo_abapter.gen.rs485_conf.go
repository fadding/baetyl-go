package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _Rs485ConfMgr struct {
	*_BaseMgr
}

// Rs485ConfMgr open func
func Rs485ConfMgr(db *gorm.DB) *_Rs485ConfMgr {
	if db == nil {
		panic(fmt.Errorf("Rs485ConfMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Rs485ConfMgr{_BaseMgr: &_BaseMgr{DB: db.Table("rs485_conf"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Rs485ConfMgr) GetTableName() string {
	return "rs485_conf"
}

// Reset 重置gorm会话
func (obj *_Rs485ConfMgr) Reset() *_Rs485ConfMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_Rs485ConfMgr) Get() (result Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Rs485ConfMgr) Gets() (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Rs485ConfMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一自增主键
func (obj *_Rs485ConfMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSLAveConfID slave_conf_id获取 slave配置id
func (obj *_Rs485ConfMgr) WithSLAveConfID(slaveConfID int) Option {
	return optionFunc(func(o *options) { o.query["slave_conf_id"] = slaveConfID })
}

// WithRs485ConfID rs485_conf_id获取 rs485配置id
func (obj *_Rs485ConfMgr) WithRs485ConfID(rs485ConfID string) Option {
	return optionFunc(func(o *options) { o.query["rs485_conf_id"] = rs485ConfID })
}

// WithEnabled enabled获取 配置是否开启，默认是开启的
func (obj *_Rs485ConfMgr) WithEnabled(enabled bool) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithDelayRtsBeforeSend delay_rts_before_send获取 发送前rts，rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) WithDelayRtsBeforeSend(delayRtsBeforeSend int) Option {
	return optionFunc(func(o *options) { o.query["delay_rts_before_send"] = delayRtsBeforeSend })
}

// WithDelayRtsAfterSend delay_rts_after_send获取 发送后rts,rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) WithDelayRtsAfterSend(delayRtsAfterSend int) Option {
	return optionFunc(func(o *options) { o.query["delay_rts_after_send"] = delayRtsAfterSend })
}

// WithRtsHighDuringSend rts_high_during_send获取 发送过程中是否设置rts为high
func (obj *_Rs485ConfMgr) WithRtsHighDuringSend(rtsHighDuringSend bool) Option {
	return optionFunc(func(o *options) { o.query["rts_high_during_send"] = rtsHighDuringSend })
}

// WithRtsHighAfterSend rts_high_after_send获取 发送后是否设置rts为high
func (obj *_Rs485ConfMgr) WithRtsHighAfterSend(rtsHighAfterSend bool) Option {
	return optionFunc(func(o *options) { o.query["rts_high_after_send"] = rtsHighAfterSend })
}

// WithRxDuringTx rx_during_tx获取 在tx过程中的rx
func (obj *_Rs485ConfMgr) WithRxDuringTx(rxDuringTx bool) Option {
	return optionFunc(func(o *options) { o.query["rx_during_tx"] = rxDuringTx })
}

// WithCreateTime create_time获取 创建时间
func (obj *_Rs485ConfMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_Rs485ConfMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_Rs485ConfMgr) GetByOption(opts ...Option) (result Rs485Conf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Rs485ConfMgr) GetByOptions(opts ...Option) (results []*Rs485Conf, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_Rs485ConfMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Rs485Conf, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where(options.query)
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
func (obj *_Rs485ConfMgr) GetFromID(id int64) (result Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一自增主键
func (obj *_Rs485ConfMgr) GetBatchFromID(ids []int64) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSLAveConfID 通过slave_conf_id获取内容 slave配置id
func (obj *_Rs485ConfMgr) GetFromSLAveConfID(slaveConfID int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`slave_conf_id` = ?", slaveConfID).Find(&results).Error

	return
}

// GetBatchFromSLAveConfID 批量查找 slave配置id
func (obj *_Rs485ConfMgr) GetBatchFromSLAveConfID(slaveConfIDs []int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`slave_conf_id` IN (?)", slaveConfIDs).Find(&results).Error

	return
}

// GetFromRs485ConfID 通过rs485_conf_id获取内容 rs485配置id
func (obj *_Rs485ConfMgr) GetFromRs485ConfID(rs485ConfID string) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rs485_conf_id` = ?", rs485ConfID).Find(&results).Error

	return
}

// GetBatchFromRs485ConfID 批量查找 rs485配置id
func (obj *_Rs485ConfMgr) GetBatchFromRs485ConfID(rs485ConfIDs []string) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rs485_conf_id` IN (?)", rs485ConfIDs).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容 配置是否开启，默认是开启的
func (obj *_Rs485ConfMgr) GetFromEnabled(enabled bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`enabled` = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量查找 配置是否开启，默认是开启的
func (obj *_Rs485ConfMgr) GetBatchFromEnabled(enableds []bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`enabled` IN (?)", enableds).Find(&results).Error

	return
}

// GetFromDelayRtsBeforeSend 通过delay_rts_before_send获取内容 发送前rts，rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) GetFromDelayRtsBeforeSend(delayRtsBeforeSend int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`delay_rts_before_send` = ?", delayRtsBeforeSend).Find(&results).Error

	return
}

// GetBatchFromDelayRtsBeforeSend 批量查找 发送前rts，rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) GetBatchFromDelayRtsBeforeSend(delayRtsBeforeSends []int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`delay_rts_before_send` IN (?)", delayRtsBeforeSends).Find(&results).Error

	return
}

// GetFromDelayRtsAfterSend 通过delay_rts_after_send获取内容 发送后rts,rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) GetFromDelayRtsAfterSend(delayRtsAfterSend int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`delay_rts_after_send` = ?", delayRtsAfterSend).Find(&results).Error

	return
}

// GetBatchFromDelayRtsAfterSend 批量查找 发送后rts,rts(Residual Time Stamp)
func (obj *_Rs485ConfMgr) GetBatchFromDelayRtsAfterSend(delayRtsAfterSends []int) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`delay_rts_after_send` IN (?)", delayRtsAfterSends).Find(&results).Error

	return
}

// GetFromRtsHighDuringSend 通过rts_high_during_send获取内容 发送过程中是否设置rts为high
func (obj *_Rs485ConfMgr) GetFromRtsHighDuringSend(rtsHighDuringSend bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rts_high_during_send` = ?", rtsHighDuringSend).Find(&results).Error

	return
}

// GetBatchFromRtsHighDuringSend 批量查找 发送过程中是否设置rts为high
func (obj *_Rs485ConfMgr) GetBatchFromRtsHighDuringSend(rtsHighDuringSends []bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rts_high_during_send` IN (?)", rtsHighDuringSends).Find(&results).Error

	return
}

// GetFromRtsHighAfterSend 通过rts_high_after_send获取内容 发送后是否设置rts为high
func (obj *_Rs485ConfMgr) GetFromRtsHighAfterSend(rtsHighAfterSend bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rts_high_after_send` = ?", rtsHighAfterSend).Find(&results).Error

	return
}

// GetBatchFromRtsHighAfterSend 批量查找 发送后是否设置rts为high
func (obj *_Rs485ConfMgr) GetBatchFromRtsHighAfterSend(rtsHighAfterSends []bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rts_high_after_send` IN (?)", rtsHighAfterSends).Find(&results).Error

	return
}

// GetFromRxDuringTx 通过rx_during_tx获取内容 在tx过程中的rx
func (obj *_Rs485ConfMgr) GetFromRxDuringTx(rxDuringTx bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rx_during_tx` = ?", rxDuringTx).Find(&results).Error

	return
}

// GetBatchFromRxDuringTx 批量查找 在tx过程中的rx
func (obj *_Rs485ConfMgr) GetBatchFromRxDuringTx(rxDuringTxs []bool) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`rx_during_tx` IN (?)", rxDuringTxs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_Rs485ConfMgr) GetFromCreateTime(createTime time.Time) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_Rs485ConfMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_Rs485ConfMgr) GetFromUpdateTime(updateTime time.Time) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_Rs485ConfMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_Rs485ConfMgr) FetchByPrimaryKey(id int64) (result Rs485Conf, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Rs485Conf{}).Where("`id` = ?", id).First(&result).Error

	return
}
