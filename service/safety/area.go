package safety

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"gorm.io/gorm"
)

type AreaService struct {
}

// CreateArea 创建Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) CreateArea(inputArea safety.Area) (err error, areaId uint) {
	var area safety.Area
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND area_name = ?", inputArea.FactoryName, inputArea.AreaName).First(&area).Error, gorm.ErrRecordNotFound) {
		return errors.New("区域已存在"), 0
	}

	err = global.GVA_DB.Create(&inputArea).Error
	if err != nil {
		return err, 0
	}

	var searchArea safety.Area
	err = global.GVA_DB.Where("factory_name = ? And area_name = ?", inputArea.FactoryName, inputArea.AreaName).First(&searchArea).Error
	if err != nil {
		return err, 0
	}

	return nil, searchArea.ID
}

// DeleteArea 删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)DeleteArea(area safety.Area) (err error) {
	if !areaService.IsLeafNode(area) {
		return errors.New("不允许删除非叶子节点")
	}
	err = global.GVA_DB.Delete(&area).Error
	return err
}

func (areaService *AreaService)DeleteAreaByFactoryName(area safety.Area) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Area{},"factory_name = ?", area.FactoryName).Error
	return err
}

func (areaService *AreaService)IsLeafNode(area safety.Area) bool {
	var total int64
	db := global.GVA_DB.Model(&safety.Area{})
	err := db.Where("factory_name = ? And parent_id = ?", area.FactoryName, area.ID).Count(&total).Error
	if err!=nil {
		return false
	}

	if total == 0 {
		return true
	}
	return false
}

// DeleteAreaByIds 批量删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)DeleteAreaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Area{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArea 更新Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)UpdateArea(area safety.Area) (err error) {
	err = global.GVA_DB.Save(&area).Error
	return err
}

// GetArea 根据id获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)GetArea(id uint) (err error, area safety.Area) {
	err = global.GVA_DB.Where("id = ?", id).First(&area).Error
	return
}

func (areaService *AreaService)GetAreaByParentId(factoryName string, parentId int) (error, []safety.Area) {
	var areaChildren []safety.Area
	err := global.GVA_DB.Where("factory_name = ? And parent_id = ?", factoryName, parentId).Find(&areaChildren).Error
	if err != nil {
		return err, nil
	}
	return nil, areaChildren
}

func (areaService *AreaService)GetRootAreaId(factoryName string) (error, uint) {
	var rootArea safety.Area
	err := global.GVA_DB.Where("factory_name = ? And area_name = ?", factoryName, factoryName).First(&rootArea).Error
	if err != nil {
		return err, 0
	}
	return nil, rootArea.ID
}

//// GetAreaInfoList 分页获取Area记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (areaService *AreaService)GetAreaInfoList(info safetyReq.AreaSearch) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//    // 创建db
//	db := global.GVA_DB.Model(&safety.Area{})
//    var areas []safety.Area
//    // 如果有条件搜索 下方会自动创建搜索语句
//	err = db.Where("factory_name = ? ", info.FactoryName).Count(&total).Error
//	if err!=nil {
//    	return
//    }
//	err = db.Limit(limit).Offset(offset).Find(&areas, "factory_name = ? ", info.FactoryName).Error
//	return err, areas, total
//}
