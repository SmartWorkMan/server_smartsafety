package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BuildingInfoRouter struct {
}

// InitBuildingInfoRouter 初始化 BuildingInfo 路由信息
func (s *BuildingInfoRouter) InitBuildingInfoRouter(Router *gin.RouterGroup) {
	buildingInfoRouter := Router.Group("buildingInfo").Use(middleware.OperationRecord())
	buildingInfoRouterWithoutRecord := Router.Group("buildingInfo")
	var buildingInfoApi = v1.ApiGroupApp.SafetyApiGroup.BuildingInfoApi
	{
		buildingInfoRouter.POST("createBuildingInfo", buildingInfoApi.CreateBuildingInfo)   // 新建BuildingInfo
		buildingInfoRouter.DELETE("deleteBuildingInfo", buildingInfoApi.DeleteBuildingInfo) // 删除BuildingInfo
		//buildingInfoRouter.DELETE("deleteBuildingInfoByIds", buildingInfoApi.DeleteBuildingInfoByIds) // 批量删除BuildingInfo
		//buildingInfoRouter.PUT("updateBuildingInfo", buildingInfoApi.UpdateBuildingInfo)    // 更新BuildingInfo
	}
	{
		buildingInfoRouterWithoutRecord.POST("getBuildingInfo", buildingInfoApi.GetBuildingInfo)        // 根据ID获取BuildingInfo
		//buildingInfoRouterWithoutRecord.GET("getBuildingInfoList", buildingInfoApi.GetBuildingInfoList)  // 获取BuildingInfo列表
	}
}
