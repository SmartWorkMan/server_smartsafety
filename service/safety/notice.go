package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
	"gorm.io/gorm/clause"
)

type NoticeService struct {
}

// CreateNotice 创建Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService) CreateNotice(inputNotice safety.Notice) (err error) {
	var notice safety.Notice
	if !errors.Is(global.GVA_DB.Where("org_name = ? AND topic = ? AND type = ?", inputNotice.OrgName, inputNotice.Topic, inputNotice.Type).First(&notice).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在同名通知")
	}
	err = global.GVA_DB.Create(&inputNotice).Error
	return err
}

func (noticeService *NoticeService) ReadNotice(noticeRead safety.NoticeRead) (err error) {
	var read safety.NoticeRead
	if errors.Is(global.GVA_DB.Where("username = ? AND notice_id = ?", noticeRead.Username, noticeRead.NoticeId).First(&read).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&noticeRead).Error
		return err
	} else {
		return nil
	}
}

// DeleteNotice 删除Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)DeleteNotice(notice safety.Notice) (err error) {
	err = global.GVA_DB.Delete(&notice).Error
	return err
}

// DeleteNoticeByIds 批量删除Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)DeleteNoticeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Notice{},"id in ?",ids.Ids).Error
	return err
}

// UpdateNotice 更新Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)UpdateNotice(notice safety.Notice) (err error) {
	update := safety.Notice{
		Type: notice.Type,
		OrgName: notice.OrgName,
		Topic: notice.Topic,
		Content: notice.Content,
		NoticeTime: notice.NoticeTime,
	}
	err = global.GVA_DB.Model(&safety.Notice{}).Where("id = ?", notice.ID).Updates(update).Error
	return err
}

// GetNotice 根据id获取Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)GetNotice(id uint) (err error, notice safety.Notice) {
	err = global.GVA_DB.Where("id = ?", id).First(&notice).Error
	return
}

// GetNoticeInfoList 分页获取Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)GetNoticeInfoList(info safetyReq.NoticeSearch) (err error, notices []safety.Notice, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Notice{})

    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("org_name = ? AND type = ?", info.OrgName, info.Type).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "notice_time"},
		Desc:   true,
	}).Offset(offset).Find(&notices, "org_name = ? AND type = ?", info.OrgName, info.Type).Error
	return err, notices, total
}

func (noticeService *NoticeService)GetNoticeReadList(info safetyReq.NoticeSearch) (err error, reads []safety.NoticeRead) {
	// 创建db
	db := global.GVA_DB.Model(&safety.NoticeRead{})

	err = db.Find(&reads, "username = ?", info.Username).Error
	return err, reads
}