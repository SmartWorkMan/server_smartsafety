package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
)

type LocationLibraryService struct {
}

// CreateLocationLibrary 创建LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService) CreateLocationLibrary(locationLibrary safety.LocationLibrary) (err error) {
	var info safety.LocationLibrary
	if errors.Is(global.GVA_DB.Where("factory_name = ? AND location_name = ?", locationLibrary.FactoryName, locationLibrary.LocationName).First(&info).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&locationLibrary).Error
	} else {
		return errors.New("重点部位已存在")
	}
	return err
}

// DeleteLocationLibrary 删除LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService)DeleteLocationLibrary(locationLibrary safety.LocationLibrary) (err error) {
	err = global.GVA_DB.Delete(&locationLibrary).Error
	return err
}

// DeleteLocationLibraryByIds 批量删除LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService)DeleteLocationLibraryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.LocationLibrary{},"id in ?",ids.Ids).Error
	return err
}

func (locationLibraryService *LocationLibraryService)DeleteLocationLibraryByFactory(factoryName string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.LocationLibrary{},"factory_name = ?",factoryName).Error
	return err
}

// UpdateLocationLibrary 更新LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService)UpdateLocationLibrary(locationLibrary safety.LocationLibrary) (err error) {
	err = global.GVA_DB.Model(&safety.LocationLibrary{}).Where("id = ?", locationLibrary.ID).Updates(safety.LocationLibrary{LocationName: locationLibrary.LocationName}).Error
	return err
}

// GetLocationLibrary 根据id获取LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService)GetLocationLibrary(id uint) (err error, locationLibrary safety.LocationLibrary) {
	err = global.GVA_DB.Where("id = ?", id).First(&locationLibrary).Error
	return
}

// GetLocationLibraryInfoList 分页获取LocationLibrary记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationLibraryService *LocationLibraryService)GetLocationLibraryInfoList(info safetyReq.LocationLibrarySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.LocationLibrary{})
    var locationLibrarys []safety.LocationLibrary
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ?", info.FactoryName).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&locationLibrarys, "factory_name = ?", info.FactoryName).Error
	return err, locationLibrarys, total
}
