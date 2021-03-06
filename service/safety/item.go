package safety

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"time"
)

type ItemService struct {
}

// CreateItem 创建Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService) CreateItem(inputItem safety.Item) (err error) {
	var item safety.Item
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND area_name = ? AND item_name = ? AND item_sn = ?", inputItem.FactoryName, inputItem.AreaName, inputItem.ItemName, inputItem.ItemSn).First(&item).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前巡检区域该巡检事项已存在")
	}
	err = global.GVA_DB.Create(&inputItem).Error
	return err
}

// DeleteItem 删除Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)DeleteItem(item safety.Item) (err error) {
	err = global.GVA_DB.Delete(&item).Error
	return err
}

func (itemService *ItemService)DeleteItemByAreaId(item safety.Item) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Item{},"factory_name = ? AND area_id = ?", item.FactoryName, item.AreaId).Error
	return err
}

func (itemService *ItemService)DeleteItemByFactoryName(item safety.Item) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Item{},"factory_name = ?", item.FactoryName).Error
	return err
}

// DeleteItemByIds 批量删除Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)DeleteItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Item{},"id in ?",ids.Ids).Error
	return err
}

// UpdateItem 更新Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)UpdateItem(item safety.Item) (err error) {
	err = global.GVA_DB.Save(&item).Error
	return err
}

func (itemService *ItemService)EnableItem(item safety.Item) (err error) {
	db := global.GVA_DB.Model(&safety.Item{})
	err = db.Where("factory_name = ? And id = ?", item.FactoryName, item.ID).UpdateColumn("enable", 1).Error
	return err
}

func (itemService *ItemService)DisableItem(item safety.Item) (err error) {
	db := global.GVA_DB.Model(&safety.Item{})
	err = db.Where("factory_name = ? And id = ?", item.FactoryName, item.ID).UpdateColumn("enable", 0).Error
	return err
}

// GetItem 根据id获取Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)GetItem(id uint) (err error, item safety.Item) {
	err = global.GVA_DB.Where("id = ?", id).First(&item).Error
	return
}

// GetItemInfoList 分页获取Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)GetItemInfoList(info safetyReq.ItemSearch, enableExist bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Item{})
    var items []safety.Item

	if info.Period == "" && enableExist == false && info.ItemName == "" {
		err = db.Where("factory_name = ? ", info.FactoryName).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? ", info.FactoryName).Error
	} else if info.Period != "" && enableExist == false && info.ItemName == "" {
		err = db.Where("factory_name = ? AND period = ?", info.FactoryName, info.Period).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND period = ?", info.FactoryName, info.Period).Error
	} else if info.Period == "" && enableExist == true && info.ItemName == "" {
		err = db.Where("factory_name = ? AND enable = ?", info.FactoryName, info.Enable).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND enable = ?", info.FactoryName, info.Enable).Error
	} else if info.Period == "" && enableExist == false && info.ItemName != "" {
		err = db.Where("factory_name = ? AND item_name like ?", info.FactoryName, "%"+info.ItemName+"%").Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND item_name like ?", info.FactoryName, "%"+info.ItemName+"%").Error
	} else if info.Period != "" && enableExist == true && info.ItemName == "" {
		err = db.Where("factory_name = ? AND period = ? AND enable = ?", info.FactoryName, info.Period, info.Enable).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND period = ? AND enable = ?", info.FactoryName, info.Period, info.Enable).Error
	} else if info.Period == "" && enableExist == true && info.ItemName != "" {
		err = db.Where("factory_name = ? AND enable = ? AND item_name like ?", info.FactoryName, info.Enable, "%"+info.ItemName+"%").Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND enable = ? AND item_name like ?", info.FactoryName, info.Enable, "%"+info.ItemName+"%").Error
	} else if info.Period != "" && enableExist == false && info.ItemName != "" {
		err = db.Where("factory_name = ? AND period = ? AND item_name like ?", info.FactoryName, info.Period, "%"+info.ItemName+"%").Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND period = ? AND item_name like ?", info.FactoryName, info.Period, "%"+info.ItemName+"%").Error
	} else if info.Period != "" && enableExist == true && info.ItemName != "" {
		err = db.Where("factory_name = ? AND period = ? AND enable = ? AND item_name like ?", info.FactoryName, info.Period, info.Enable, "%"+info.ItemName+"%").Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? AND period = ? AND enable = ? AND item_name like ?", info.FactoryName, info.Period, info.Enable, "%"+info.ItemName+"%").Error
	}

	return err, items, total
}

func (itemService *ItemService)GetAllValidItemList(period string) (error, []safety.Item) {
	// 创建db
	db := global.GVA_DB.Model(&safety.Item{})
	var items []safety.Item
	curTime := time.Now().Format("2006-01-02 15:04:05")
	err := db.Find(&items, "enable = 1 AND start_time < ? AND end_time > ? AND period = ?", curTime, curTime, period).Error
	
	return err, items
}

// GetItemInfoListByArea 分页获取Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService)GetItemInfoListByLeafAreaId(info safetyReq.ItemSearch, inIdList []uint) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Item{})
	var items []safety.Item

	err = db.Where("factory_name = ? And area_id in (?)", info.FactoryName, inIdList).Count(&total).Error
	if err!=nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? And area_id in (?)", info.FactoryName, inIdList).Error
	return err, items, total
}

func (itemService *ItemService) CreateNextPeriodDate(inputDate safety.ItemNextPeriodDate) (err error) {
	var nextDate safety.ItemNextPeriodDate
	if !errors.Is(global.GVA_DB.Where("period = ?", inputDate.Period).First(&nextDate).Error, gorm.ErrRecordNotFound) {
		return errors.New("周期已存在")
	}
	err = global.GVA_DB.Create(&inputDate).Error
	return err
}

func (itemService *ItemService)UpdateNextPeriodDate(inputDate safety.ItemNextPeriodDate) (err error) {
	db := global.GVA_DB.Model(&safety.ItemNextPeriodDate{})
	err = db.Where("period = ?", inputDate.Period).UpdateColumn("next_date", inputDate.NextDate).Error
	if err != nil {
		return err
	}

	if inputDate.Interval != 0 {
		err = db.Where("period = ?", inputDate.Period).UpdateColumn("interval", inputDate.Interval).Error
	}
	return err
}

func (itemService *ItemService)IsPeriodExist(inputDate safety.ItemNextPeriodDate) bool {
	var nextDate safety.ItemNextPeriodDate
	if !errors.Is(global.GVA_DB.Where("period = ?", inputDate.Period).First(&nextDate).Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (itemService *ItemService)GetNextPeriodDate(inputDate safety.ItemNextPeriodDate) (error, safety.ItemNextPeriodDate) {
	var nextDate safety.ItemNextPeriodDate
	err := global.GVA_DB.Where("period = ?", inputDate.Period).First(&nextDate).Error
	return err, nextDate
}