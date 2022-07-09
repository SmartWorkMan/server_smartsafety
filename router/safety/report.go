package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReportRouter struct {
}

// InitReportRouter 初始化 Report 路由信息
func (s *ReportRouter) InitReportRouter(Router *gin.RouterGroup) {
	reportRouter := Router.Group("report").Use(middleware.OperationRecord())
	reportRouterWithoutRecord := Router.Group("report")
	var reportApi = v1.ApiGroupApp.SafetyApiGroup.ReportApi
	{
		reportRouter.POST("app/createReport", reportApi.CreateReport)   // 新建Report
		reportRouter.POST("app/applyReport", reportApi.ApplyReport)
		reportRouter.DELETE("app/deleteReport", reportApi.DeleteReport) // 删除Report
		//reportRouter.DELETE("deleteReportByIds", reportApi.DeleteReportByIds) // 批量删除Report
		reportRouter.PUT("app/updateReport", reportApi.UpdateReport)    // 更新Report
	}
	{
		//reportRouterWithoutRecord.GET("findReport", reportApi.FindReport)        // 根据ID获取Report
		reportRouterWithoutRecord.POST("app/getReportListByUser", reportApi.GetReportListByUser)  // 获取Report列表
		reportRouterWithoutRecord.POST("app/getFormalReportListByUser", reportApi.GetFormalReportListByUser)  // 获取Report列表
		reportRouterWithoutRecord.POST("getFormalReportList", reportApi.GetFormalReportList)  // 获取Report列表
	}
}
