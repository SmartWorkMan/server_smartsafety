package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LawsRouter struct {
}

// InitLawsRouter 初始化 Laws 路由信息
func (s *LawsRouter) InitLawsRouter(Router *gin.RouterGroup) {
	lawsRouter := Router.Group("laws").Use(middleware.OperationRecord())
	lawsRouterWithoutRecord := Router.Group("laws")
	var lawsApi = v1.ApiGroupApp.SafetyApiGroup.LawsApi
	{
		lawsRouter.POST("createLaws", lawsApi.CreateLaws)   // 新建Laws
		lawsRouter.DELETE("deleteLaws", lawsApi.DeleteLaws) // 删除Laws
		lawsRouter.DELETE("deleteLawsByIds", lawsApi.DeleteLawsByIds) // 批量删除Laws
		lawsRouter.PUT("updateLawsStatus", lawsApi.UpdateLawsStatus)    // 更新Laws
	}
	{
		lawsRouterWithoutRecord.GET("findLaws", lawsApi.FindLaws)        // 根据ID获取Laws
		lawsRouterWithoutRecord.POST("getLawsList", lawsApi.GetLawsList)  // 获取Laws列表
		lawsRouterWithoutRecord.POST("getLawsListForAdmin", lawsApi.GetLawsListForAdmin)  // 获取Laws列表
	}
}
