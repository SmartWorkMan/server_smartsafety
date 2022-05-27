package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
	"time"
)

type ItemService struct {
}

// CreateItem 创建Item记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemService *ItemService) CreateItem(inputItem safety.Item) (err error) {
	var item safety.Item
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND area_name = ? AND item_name = ?", inputItem.FactoryName, inputItem.AreaName, inputItem.ItemName).First(&item).Error, gorm.ErrRecordNotFound) {
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
func (itemService *ItemService)GetItemInfoList(info safetyReq.ItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Item{})
    var items []safety.Item
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? ", info.FactoryName).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? ", info.FactoryName).Error
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
func (itemService *ItemService)GetItemInfoListByLeafAreaId(info safetyReq.ItemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Item{})
	var items []safety.Item
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? And area_id = ?", info.FactoryName, info.AreaId).Count(&total).Error
	if err!=nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&items, "factory_name = ? And area_id = ?", info.FactoryName, info.AreaId).Error
	return err, items, total
}

