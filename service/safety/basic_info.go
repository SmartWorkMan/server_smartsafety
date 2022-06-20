package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
)

type BasicInfoService struct {
}

// CreateBasicInfo 创建BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService) CreateBasicInfo(basicInfo safety.BasicInfo) (err error) {
	var info safety.BasicInfo
	if errors.Is(global.GVA_DB.Where("factory_name = ?", basicInfo.FactoryName).First(&info).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&basicInfo).Error
	} else {
		err = global.GVA_DB.Model(&safety.BasicInfo{}).Where("factory_name = ?", basicInfo.FactoryName).Updates(safety.BasicInfo{InfoJson: basicInfo.InfoJson}).Error
	}
	return err
}

// DeleteBasicInfo 删除BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService)DeleteBasicInfo(basicInfo safety.BasicInfo) (err error) {
	err = global.GVA_DB.Delete(&basicInfo).Error
	return err
}

// DeleteBasicInfoByIds 批量删除BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService)DeleteBasicInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.BasicInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBasicInfo 更新BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService)UpdateBasicInfo(basicInfo safety.BasicInfo) (err error) {
	err = global.GVA_DB.Save(&basicInfo).Error
	return err
}

// GetBasicInfo 根据id获取BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService)GetBasicInfo(factoryName string) (err error, basicInfo safety.BasicInfo) {
	err = global.GVA_DB.Where("factory_name = ?", factoryName).First(&basicInfo).Error
	return
}

// GetBasicInfoInfoList 分页获取BasicInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicInfoService *BasicInfoService)GetBasicInfoInfoList(info safetyReq.BasicInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.BasicInfo{})
    var basicInfos []safety.BasicInfo
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&basicInfos).Error
	return err, basicInfos, total
}
