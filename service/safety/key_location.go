package safety

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
)

type KeyLocationService struct {
}

// CreateKeyLocation 创建KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService) CreateKeyLocation(keyLocation safety.KeyLocation) (err error) {
	var kl safety.KeyLocation
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND location_name = ?", keyLocation.FactoryName, keyLocation.LocationName).First(&kl).Error, gorm.ErrRecordNotFound) {
		return errors.New("重点部位已创建")
	}
	err = global.GVA_DB.Create(&keyLocation).Error
	return err
}

// DeleteKeyLocation 删除KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService)DeleteKeyLocation(keyLocation safety.KeyLocation) (err error) {
	err = global.GVA_DB.Delete(&keyLocation).Error
	return err
}

// DeleteKeyLocationByIds 批量删除KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService)DeleteKeyLocationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.KeyLocation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateKeyLocation 更新KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService)UpdateKeyLocation(keyLocation safety.KeyLocation) (err error) {
	db := global.GVA_DB.Model(&safety.KeyLocation{})
	updateKL := safety.KeyLocation{
		LocationName: keyLocation.LocationName,
		LocationImage: keyLocation.LocationImage,
		DutyOfficer: keyLocation.DutyOfficer,
		Description: keyLocation.Description,
		Place: keyLocation.Place,
		Attachment: keyLocation.Attachment,
		AttachmentName: keyLocation.AttachmentName,
	}
	err = db.Where("id = ?", keyLocation.ID).Updates(updateKL).Error
	return err
}

// GetKeyLocation 根据id获取KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService)GetKeyLocation(id uint) (err error, keyLocation safety.KeyLocation) {
	err = global.GVA_DB.Where("id = ?", id).First(&keyLocation).Error
	return
}

// GetKeyLocationInfoList 分页获取KeyLocation记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyLocationService *KeyLocationService)GetKeyLocationInfoList(info safetyReq.KeyLocationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&safety.KeyLocation{})
    var keyLocations []safety.KeyLocation

	err = db.Where("factory_name = ?", info.FactoryName).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&keyLocations, "factory_name = ?", info.FactoryName).Error
	return err, keyLocations, total
}
