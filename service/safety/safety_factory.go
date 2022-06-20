package safety

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"gorm.io/gorm"
)

type SafetyFactoryService struct {
}

// CreateSafetyFactory 创建SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService) CreateSafetyFactory(safetyFactory safety.SafetyFactory) (err error) {
	var factory safety.SafetyFactory
	if !errors.Is(global.GVA_DB.Where("factory_name = ?", safetyFactory.FactoryName).First(&factory).Error, gorm.ErrRecordNotFound) {
		return errors.New("工厂名称已经存在")
	}

	err = global.GVA_DB.Create(&safetyFactory).Error
	return err
}

// DeleteSafetyFactory 删除SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService)DeleteSafetyFactory(safetyFactory safety.SafetyFactory) (err error) {
	err = global.GVA_DB.Delete(&safetyFactory).Error
	return err
}

// DeleteSafetyFactoryByIds 批量删除SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService)DeleteSafetyFactoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.SafetyFactory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSafetyFactory 更新SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService)UpdateSafetyFactory(safetyFactory safety.SafetyFactory) (err error) {
	err = global.GVA_DB.Save(&safetyFactory).Error
	return err
}

func (safetyFactoryService *SafetyFactoryService)UpdateFactoryLatLng(safetyFactory safety.SafetyFactory) (err error) {
	db := global.GVA_DB.Model(&safety.SafetyFactory{})
	updateFactory := safety.SafetyFactory{ Lat: safetyFactory.Lat, Lng: safetyFactory.Lng}
	err = db.Where("id = ?", safetyFactory.ID).Updates(updateFactory).Error
	return err
}

// GetSafetyFactory 根据id获取SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService)QuerySafetyFactory(inputFactory safety.SafetyFactory) (err error, safetyFactory safety.SafetyFactory) {
	err = global.GVA_DB.Where("factory_name = ?", inputFactory.FactoryName).First(&safetyFactory).Error
	return
}

func (safetyFactoryService *SafetyFactoryService)GetSafetyFactory(id uint) (err error, safetyFactory safety.SafetyFactory) {
	err = global.GVA_DB.Where("id = ?", id).First(&safetyFactory).Error
	return
}

// GetSafetyFactoryInfoList 分页获取SafetyFactory记录
// Author [piexlmax](https://github.com/piexlmax)
func (safetyFactoryService *SafetyFactoryService)GetSafetyFactoryInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.SafetyFactory{})
    var safetyFactorys []safety.SafetyFactory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&safetyFactorys).Error
	return err, safetyFactorys, total
}

func (safetyFactoryService *SafetyFactoryService)GetSafetyFactoryByFactoryName(factoryName string) (err error, safetyFactory safety.SafetyFactory) {
	err = global.GVA_DB.Where("factory_name = ?", factoryName).First(&safetyFactory).Error
	return
}

func (safetyFactoryService *SafetyFactoryService)GetSafetyFactoryByFactoryId(factoryId string) (err error, safetyFactory safety.SafetyFactory) {
	err = global.GVA_DB.Where("factory_id = ?", factoryId).First(&safetyFactory).Error
	return
}