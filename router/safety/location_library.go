package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LocationLibraryRouter struct {
}

// InitLocationLibraryRouter 初始化 LocationLibrary 路由信息
func (s *LocationLibraryRouter) InitLocationLibraryRouter(Router *gin.RouterGroup) {
	locationLibraryRouter := Router.Group("locationLibrary").Use(middleware.OperationRecord())
	locationLibraryRouterWithoutRecord := Router.Group("locationLibrary")
	var locationLibraryApi = v1.ApiGroupApp.SafetyApiGroup.LocationLibraryApi
	{
		locationLibraryRouter.POST("createLocationLibrary", locationLibraryApi.CreateLocationLibrary)   // 新建LocationLibrary
		locationLibraryRouter.DELETE("deleteLocationLibrary", locationLibraryApi.DeleteLocationLibrary) // 删除LocationLibrary
		//locationLibraryRouter.DELETE("deleteLocationLibraryByIds", locationLibraryApi.DeleteLocationLibraryByIds) // 批量删除LocationLibrary
		locationLibraryRouter.PUT("updateLocationLibrary", locationLibraryApi.UpdateLocationLibrary)    // 更新LocationLibrary
	}
	{
		//locationLibraryRouterWithoutRecord.GET("findLocationLibrary", locationLibraryApi.FindLocationLibrary)        // 根据ID获取LocationLibrary
		locationLibraryRouterWithoutRecord.POST("getLocationLibraryList", locationLibraryApi.GetLocationLibraryList)  // 获取LocationLibrary列表
	}
}
