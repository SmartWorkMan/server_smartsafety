package safety

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
)

type InspectorService struct {
}

// CreateInspector 创建Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService) CreateInspector(inspector safety.Inspector) (err error) {
	var insp safety.Inspector
	if !errors.Is(global.GVA_DB.Where("username = ?", inspector.Username).First(&insp).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已存在")
	}

	err = global.GVA_DB.Create(&inspector).Error
	return err
}

// DeleteInspector 删除Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService)DeleteInspector(inspector safety.Inspector) (err error) {
	err = global.GVA_DB.Delete(&inspector).Error
	return err
}

// DeleteInspectorByIds 批量删除Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService)DeleteInspectorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Inspector{},"id in ?",ids.Ids).Error
	return err
}

func (inspectorService *InspectorService)DeleteInspectorByFactory(factoryName string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Inspector{},"factory_name = ?", factoryName).Error
	return err
}

// UpdateInspector 更新Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService)UpdateInspector(inspector safety.Inspector) (err error) {
	err = global.GVA_DB.Save(&inspector).Error
	return err
}

// GetInspector 根据id获取Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService)GetInspector(id uint) (err error, inspector safety.Inspector) {
	err = global.GVA_DB.Where("id = ?", id).First(&inspector).Error
	return
}

func (inspectorService *InspectorService)GetInspectorByUserName(userName string) (err error, inspector safety.Inspector) {
	err = global.GVA_DB.Where("username = ?", userName).First(&inspector).Error
	return
}

// GetInspectorInfoList 分页获取Inspector记录
// Author [piexlmax](https://github.com/piexlmax)
func (inspectorService *InspectorService)GetInspectorInfoList(info safetyReq.InspectorSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Inspector{})
    var inspectors []safety.Inspector
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? ", info.FactoryName).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&inspectors, "factory_name = ? ", info.FactoryName).Error
	return err, inspectors, total
}

func (inspectorService *InspectorService) Login(inputInspector *safety.Inspector) (err error) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init")
	}

	err = global.GVA_DB.Where("username = ? AND password = ?", inputInspector.Username, inputInspector.Password).First(inputInspector).Error
	return err
}

func (inspectorService *InspectorService) IsUserNameExist(username string) bool {
	var insp safety.Inspector
	if !errors.Is(global.GVA_DB.Where("username = ?", username).First(&insp).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return true
	}
	return false
}