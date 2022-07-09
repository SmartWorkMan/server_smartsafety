package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlanRouter struct {
}

// InitPlanRouter 初始化 Plan 路由信息
func (s *PlanRouter) InitPlanRouter(Router *gin.RouterGroup) {
	planRouter := Router.Group("plan").Use(middleware.OperationRecord())
	planRouterWithoutRecord := Router.Group("plan")
	var planApi = v1.ApiGroupApp.SafetyApiGroup.PlanApi
	{
		planRouter.POST("createPlan", planApi.CreatePlan)   // 新建Plan
		planRouter.DELETE("deletePlan", planApi.DeletePlan) // 删除Plan
		//planRouter.DELETE("deletePlanByIds", planApi.DeletePlanByIds) // 批量删除Plan
		//planRouter.PUT("updatePlan", planApi.UpdatePlan)    // 更新Plan
	}
	{
		//planRouterWithoutRecord.GET("findPlan", planApi.FindPlan)        // 根据ID获取Plan
		planRouterWithoutRecord.POST("getPlanList", planApi.GetPlanList)  // 获取Plan列表
		planRouterWithoutRecord.GET("downloadPlanTemplate", planApi.DownloadPlanTemplate)
	}
}
