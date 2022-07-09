package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlanApi struct {
}


// CreatePlan 创建Plan
// @Tags Plan
// @Summary 创建Plan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Plan true "创建Plan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /plan/createPlan [post]
func (planApi *PlanApi) CreatePlan(c *gin.Context) {
	var plan safety.Plan
	_ = c.ShouldBindJSON(&plan)
	if plan.FactoryName == "" || plan.FileName == "" || plan.FileType == "" || plan.FileAddr == "" || plan.UploadTime == "" {
		global.GVA_LOG.Error("创建失败!请检查输入!")
		response.FailWithMessage("创建失败!请检查输入!", c)
		return
	}
	if plan.FileType != "2" && plan.FileType != "3" {
		global.GVA_LOG.Error("创建失败!文件类型错误!")
		response.FailWithMessage("创建失败!文件类型错误!", c)
		return
	}

	if err := planService.CreatePlan(plan); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlan 删除Plan
// @Tags Plan
// @Summary 删除Plan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Plan true "删除Plan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /plan/deletePlan [delete]
func (planApi *PlanApi) DeletePlan(c *gin.Context) {
	var plan safety.Plan
	_ = c.ShouldBindJSON(&plan)
	if err := planService.DeletePlan(plan); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlanByIds 批量删除Plan
// @Tags Plan
// @Summary 批量删除Plan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Plan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /plan/deletePlanByIds [delete]
func (planApi *PlanApi) DeletePlanByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := planService.DeletePlanByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlan 更新Plan
// @Tags Plan
// @Summary 更新Plan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Plan true "更新Plan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /plan/updatePlan [put]
func (planApi *PlanApi) UpdatePlan(c *gin.Context) {
	var plan safety.Plan
	_ = c.ShouldBindJSON(&plan)
	if err := planService.UpdatePlan(plan); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlan 用id查询Plan
// @Tags Plan
// @Summary 用id查询Plan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Plan true "用id查询Plan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /plan/findPlan [get]
func (planApi *PlanApi) FindPlan(c *gin.Context) {
	var plan safety.Plan
	_ = c.ShouldBindQuery(&plan)
	if err, replan := planService.GetPlan(plan.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replan": replan}, c)
	}
}

// GetPlanList 分页获取Plan列表
// @Tags Plan
// @Summary 分页获取Plan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.PlanSearch true "分页获取Plan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /plan/getPlanList [post]
func (planApi *PlanApi) GetPlanList(c *gin.Context) {
	var pageInfo safetyReq.PlanSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称为空!")
		response.FailWithMessage("获取失败!工厂名称为空!", c)
		return
	}
	if err, list, total := planService.GetPlanInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// @Router /plan/downloadPlanTemplate [get]
func (planApi *PlanApi) DownloadPlanTemplate(c *gin.Context) {
	var plan safety.Plan
	_ = c.ShouldBindQuery(&plan)

	filepath := "./resource/template.doc"
	filename := "template.doc"

	c.FileAttachment(filepath, filename)
}