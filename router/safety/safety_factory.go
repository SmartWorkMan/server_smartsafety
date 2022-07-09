package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SafetyFactoryRouter struct {
}

// InitSafetyFactoryRouter 初始化 SafetyFactory 路由信息
func (s *SafetyFactoryRouter) InitSafetyFactoryRouter(Router *gin.RouterGroup) {
	safetyFactoryRouter := Router.Group("safetyFactory").Use(middleware.OperationRecord())
	safetyFactoryRouterWithoutRecord := Router.Group("safetyFactory")
	var safetyFactoryApi = v1.ApiGroupApp.SafetyApiGroup.SafetyFactoryApi
	{
		safetyFactoryRouter.POST("createSafetyFactory", safetyFactoryApi.CreateSafetyFactory)   // 新建SafetyFactory
		safetyFactoryRouter.DELETE("deleteSafetyFactory", safetyFactoryApi.DeleteSafetyFactory) // 删除SafetyFactory
		safetyFactoryRouter.DELETE("deleteSafetyFactoryByIds", safetyFactoryApi.DeleteSafetyFactoryByIds) // 批量删除SafetyFactory
		//safetyFactoryRouter.PUT("updateSafetyFactory", safetyFactoryApi.UpdateSafetyFactory)    // 更新SafetyFactory
		safetyFactoryRouter.PUT("updateFactoryLatLng", safetyFactoryApi.UpdateFactoryLatLng)    // 更新SafetyFactory
	}
	{
		safetyFactoryRouterWithoutRecord.GET("findSafetyFactory", safetyFactoryApi.FindSafetyFactory)        // 根据ID获取SafetyFactory
		safetyFactoryRouterWithoutRecord.POST("getSafetyFactoryList", safetyFactoryApi.GetSafetyFactoryList)  // 获取SafetyFactory列表
		safetyFactoryRouterWithoutRecord.POST("querySafetyFactory", safetyFactoryApi.QuerySafetyFactory)  // 获取SafetyFactory列表
	}
}
