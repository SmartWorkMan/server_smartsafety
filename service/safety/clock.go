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

type ClockService struct {
}

// CreateClock 创建Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService) CreateClock(clock safety.Clock, clockType string) (err error) {
	var dbClock safety.Clock
	if errors.Is(global.GVA_DB.Where("inspector_username = ? AND clock_date = ?", clock.InspectorUsername, clock.ClockDate).First(&dbClock).Error, gorm.ErrRecordNotFound) {
		if clockType == "下班" {
			return errors.New("请先进行上班打卡才能下班打卡!")
		} else {
			err = global.GVA_DB.Create(&clock).Error
		}
	} else {
		db := global.GVA_DB.Model(&safety.Clock{})
		var updateClock safety.Clock
		if clockType == "上班" {
			updateClock = safety.Clock{
				ClockInPic: clock.ClockInPic,
				ClockInLocation: clock.ClockInLocation,
			}
		} else {
			updateClock = safety.Clock{
				ClockOutPic: clock.ClockOutPic,
				ClockOutLocation: clock.ClockOutLocation,
				ClockOutTime: clock.ClockOutTime,
			}
		}
		err = db.Where("id = ?", dbClock.ID).Updates(updateClock).Error
	}

	return err
}

func (clockService *ClockService)QueryClock(inputClock safety.Clock) (err error, clock safety.Clock) {
	err = global.GVA_DB.Where("inspector_username = ? AND clock_date = ?", inputClock.InspectorUsername, inputClock.ClockDate).First(&clock).Error
	return
}


// DeleteClock 删除Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService)DeleteClock(clock safety.Clock) (err error) {
	err = global.GVA_DB.Delete(&clock).Error
	return err
}

// DeleteClockByIds 批量删除Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService)DeleteClockByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Clock{},"id in ?",ids.Ids).Error
	return err
}

// UpdateClock 更新Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService)UpdateClock(clock safety.Clock) (err error) {
	err = global.GVA_DB.Save(&clock).Error
	return err
}

// GetClock 根据id获取Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService)GetClock(id uint) (err error, clock safety.Clock) {
	err = global.GVA_DB.Where("id = ?", id).First(&clock).Error
	return
}

// GetClockInfoList 分页获取Clock记录
// Author [piexlmax](https://github.com/piexlmax)
func (clockService *ClockService)GetTodayClockInfoList(info safetyReq.ClockSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Clock{})
    var clocks []safety.Clock

	err = db.Where("factory_name = ? AND clock_date = ?", info.FactoryName, info.ClockDate).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "clock_in_time"},
		Desc:   false,
	}).Offset(offset).Where("factory_name = ? AND clock_date = ?", info.FactoryName, info.ClockDate).Find(&clocks).Error
	return err, clocks, total
}

func (clockService *ClockService)GetOnDutyNum(clock safety.Clock) (err error, total int64) {
	db := global.GVA_DB.Model(&safety.Clock{})
	err = db.Where("factory_name = ? AND clock_date = ? AND LENGTH(trim(clock_out_time)) < 1 ", clock.FactoryName, clock.ClockDate).Count(&total).Error
	if err!=nil {
		return err, 0
	}

	return err, total
}

func (clockService *ClockService)GetHistoryClockList(info safetyReq.ClockSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Clock{})
	var clocks []safety.Clock

	err = db.Where("factory_name = ? AND inspector_username = ?", info.FactoryName, info.InspectorUsername).Count(&total).Error
	if err!=nil {
		return
	}
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "clock_in_time"},
		Desc:   true,
	}).Offset(offset).Where("factory_name = ? AND inspector_username = ?", info.FactoryName, info.InspectorUsername).Find(&clocks).Error
	return err, clocks, total
}