package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm/clause"
)

type PlanService struct {
}

// CreatePlan 创建Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService) CreatePlan(plan safety.Plan) (err error) {
	err = global.GVA_DB.Create(&plan).Error
	return err
}

// DeletePlan 删除Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService)DeletePlan(plan safety.Plan) (err error) {
	err = global.GVA_DB.Delete(&plan).Error
	return err
}

// DeletePlanByIds 批量删除Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService)DeletePlanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Plan{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePlan 更新Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService)UpdatePlan(plan safety.Plan) (err error) {
	err = global.GVA_DB.Save(&plan).Error
	return err
}

// GetPlan 根据id获取Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService)GetPlan(id uint) (err error, plan safety.Plan) {
	err = global.GVA_DB.Where("id = ?", id).First(&plan).Error
	return
}

// GetPlanInfoList 分页获取Plan记录
// Author [piexlmax](https://github.com/piexlmax)
func (planService *PlanService)GetPlanInfoList(info safetyReq.PlanSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Plan{})
    var plans []safety.Plan
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ?", info.FactoryName).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "upload_time"},
		Desc:   true,
	}).Offset(offset).Find(&plans, "factory_name = ?", info.FactoryName).Error
	return err, plans, total
}
