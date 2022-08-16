package safety

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type LawsService struct {
}

// CreateLaws 创建Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService) CreateLaws(inputLaws safety.Laws) (err error) {
	var laws safety.Laws
	if !errors.Is(global.GVA_DB.Where("law_name = ?", inputLaws.LawName).First(&laws).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前法律法规已存在")
	}
	err = global.GVA_DB.Create(&inputLaws).Error
	return err
}

// DeleteLaws 删除Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService)DeleteLaws(laws safety.Laws) (err error) {
	err = global.GVA_DB.Delete(&laws).Error
	return err
}

// DeleteLawsByIds 批量删除Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService)DeleteLawsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Laws{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLaws 更新Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService)UpdateLawsStatus(laws safety.Laws) (err error) {
	db := global.GVA_DB.Model(&safety.Laws{})
	err = db.Where("law_name = ?", laws.LawName).UpdateColumn("law_status", laws.LawStatus).Error
	return err
}

// GetLaws 根据id获取Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService)GetLaws(id uint) (err error, laws safety.Laws) {
	err = global.GVA_DB.Where("id = ?", id).First(&laws).Error
	return
}

// GetLawsInfoList 分页获取Laws记录
// Author [piexlmax](https://github.com/piexlmax)
func (lawsService *LawsService)GetLawsInfoList(info safetyReq.LawsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Laws{})
    var lawss []safety.Laws
    // 如果有条件搜索 下方会自动创建搜索语句
    descSort := true
    if strings.ToLower(info.Sort) == "asc" {
    	descSort = false
	}
    if info.LawStatus != "" && info.LawName != "" {
		err = db.Where("law_type = ? AND law_status = ? AND law_name like ?", info.LawType, info.LawStatus, "%"+info.LawName+"%").Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "release_time"},
			Desc:   descSort,
		}).Offset(offset).Find(&lawss,"law_type = ? AND law_status = ? AND law_name like ?", info.LawType, info.LawStatus, "%"+info.LawName+"%").Error
	} else if info.LawStatus == "" && info.LawName != "" {
		err = db.Where("law_type = ? AND law_name like ?", info.LawType, "%"+info.LawName+"%").Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "release_time"},
			Desc:   descSort,
		}).Offset(offset).Find(&lawss,"law_type = ? AND law_name like ?", info.LawType, "%"+info.LawName+"%").Error
	} else if info.LawStatus != "" && info.LawName == "" {
		err = db.Where("law_type = ? AND law_status = ?", info.LawType, info.LawStatus).Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "release_time"},
			Desc:   descSort,
		}).Offset(offset).Find(&lawss,"law_type = ? AND law_status = ?", info.LawType, info.LawStatus).Error
	}else if info.LawStatus == "" && info.LawName == "" {
		err = db.Where("law_type = ?", info.LawType).Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "release_time"},
			Desc:   descSort,
		}).Offset(offset).Find(&lawss, "law_type = ?", info.LawType).Error
	}

	return err, lawss, total
}

func (lawsService *LawsService)GetLawsListForApp(info safetyReq.LawsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Laws{})
	var lawss []safety.Laws

	if info.LawType != "" && info.LawName != "" {
		err = db.Where("law_type = ? AND law_name like ?", info.LawType, "%"+info.LawName+"%").Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Offset(offset).Find(&lawss,"law_type = ? AND law_name like ?", info.LawType, "%"+info.LawName+"%").Error
	} else if info.LawType == "" && info.LawName != "" {
		err = db.Where("law_name like ?", "%"+info.LawName+"%").Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Offset(offset).Find(&lawss,"law_name like ?", "%"+info.LawName+"%").Error
	} else if info.LawType != "" && info.LawName == "" {
		err = db.Where("law_type = ?", info.LawType).Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Offset(offset).Find(&lawss,"law_type = ?", info.LawType).Error
	}else if info.LawType == "" && info.LawName == "" {
		err = db.Count(&total).Error
		if err!=nil {
			return
		}

		err = db.Limit(limit).Offset(offset).Find(&lawss).Error
	}

	return err, lawss, total
}
