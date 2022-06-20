package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
)

type BuildingInfoService struct {
}

// CreateBuildingInfo 创建BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService) CreateBuildingInfo(buildingInfo safety.BuildingInfo) (err error) {
	var info safety.BuildingInfo
	if errors.Is(global.GVA_DB.Where("factory_name = ?", buildingInfo.FactoryName).First(&info).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&buildingInfo).Error
	} else {
		err = global.GVA_DB.Model(&safety.BuildingInfo{}).Where("factory_name = ?", buildingInfo.FactoryName).Updates(safety.BuildingInfo{InfoJson: buildingInfo.InfoJson}).Error
	}
	return err
}

// DeleteBuildingInfo 删除BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService)DeleteBuildingInfo(buildingInfo safety.BuildingInfo) (err error) {
	err = global.GVA_DB.Delete(&buildingInfo).Error
	return err
}

// DeleteBuildingInfoByIds 批量删除BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService)DeleteBuildingInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.BuildingInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBuildingInfo 更新BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService)UpdateBuildingInfo(buildingInfo safety.BuildingInfo) (err error) {
	err = global.GVA_DB.Save(&buildingInfo).Error
	return err
}

// GetBuildingInfo 根据id获取BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService)GetBuildingInfo(factoryName string) (err error, buildingInfo safety.BuildingInfo) {
	err = global.GVA_DB.Where("factory_name = ?", factoryName).First(&buildingInfo).Error
	return
}

// GetBuildingInfoInfoList 分页获取BuildingInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildingInfoService *BuildingInfoService)GetBuildingInfoInfoList(info safetyReq.BuildingInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.BuildingInfo{})
    var buildingInfos []safety.BuildingInfo
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&buildingInfos).Error
	return err, buildingInfos, total
}
