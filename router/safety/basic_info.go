package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicInfoRouter struct {
}

// InitBasicInfoRouter 初始化 BasicInfo 路由信息
func (s *BasicInfoRouter) InitBasicInfoRouter(Router *gin.RouterGroup) {
	basicInfoRouter := Router.Group("basicInfo").Use(middleware.OperationRecord())
	basicInfoRouterWithoutRecord := Router.Group("basicInfo")
	var basicInfoApi = v1.ApiGroupApp.SafetyApiGroup.BasicInfoApi
	{
		basicInfoRouter.POST("createBasicInfo", basicInfoApi.CreateBasicInfo)   // 新建BasicInfo
		basicInfoRouter.DELETE("deleteBasicInfo", basicInfoApi.DeleteBasicInfo) // 删除BasicInfo
	}
	{
		basicInfoRouterWithoutRecord.POST("getBasicInfo", basicInfoApi.GetBasicInfo)
	}
}
