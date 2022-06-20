package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TecentRouter struct {
}

// InitTecentRouter 初始化 Tecent 路由信息
func (s *TecentRouter) InitTecentRouter(Router *gin.RouterGroup) {
	tecentRouter := Router.Group("tencent").Use(middleware.OperationRecord())
	tecentRouterWithoutRecord := Router.Group("tencent")
	var tecentApi = v1.ApiGroupApp.SafetyApiGroup.TecentApi
	{
		tecentRouter.POST("createTecent", tecentApi.CreateTecent)   // 新建Tecent
		//tecentRouter.DELETE("deleteTecent", tecentApi.DeleteTecent) // 删除Tecent
		//tecentRouter.DELETE("deleteTecentByIds", tecentApi.DeleteTecentByIds) // 批量删除Tecent
		//tecentRouter.PUT("updateTecent", tecentApi.UpdateTecent)    // 更新Tecent
	}
	{
		//tecentRouterWithoutRecord.GET("findTecent", tecentApi.FindTecent)        // 根据ID获取Tecent
		tecentRouterWithoutRecord.POST("getTencent", tecentApi.GetTecent)  // 获取Tecent列表
	}
}
