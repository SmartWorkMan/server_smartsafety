package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
)

type TecentService struct {
}

// CreateTecent 创建Tecent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecentService *TecentService) CreateTecent(tecent safety.Tecent) (err error) {
	err = global.GVA_DB.Create(&tecent).Error
	return err
}

// DeleteTecent 删除Tecent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecentService *TecentService)DeleteTecent(tecent safety.Tecent) (err error) {
	err = global.GVA_DB.Delete(&tecent).Error
	return err
}

// DeleteTecentByIds 批量删除Tecent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecentService *TecentService)DeleteTecentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Tecent{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTecent 更新Tecent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecentService *TecentService)UpdateTecent(tecent safety.Tecent) (err error) {
	err = global.GVA_DB.Save(&tecent).Error
	return err
}

// GetTecent 根据id获取Tecent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecentService *TecentService)GetTecent(tencentType string) (err error, tecent safety.Tecent) {
	err = global.GVA_DB.Where("tencent_type = ?", tencentType).First(&tecent).Error
	return
}

