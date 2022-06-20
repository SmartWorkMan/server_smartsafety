package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KeyLocationRouter struct {
}

// InitKeyLocationRouter 初始化 KeyLocation 路由信息
func (s *KeyLocationRouter) InitKeyLocationRouter(Router *gin.RouterGroup) {
	keyLocationRouter := Router.Group("keyLocation").Use(middleware.OperationRecord())
	keyLocationRouterWithoutRecord := Router.Group("keyLocation")
	var keyLocationApi = v1.ApiGroupApp.SafetyApiGroup.KeyLocationApi
	{
		keyLocationRouter.POST("createKeyLocation", keyLocationApi.CreateKeyLocation)   // 新建KeyLocation
		keyLocationRouter.DELETE("deleteKeyLocation", keyLocationApi.DeleteKeyLocation) // 删除KeyLocation
		keyLocationRouter.DELETE("deleteKeyLocationByIds", keyLocationApi.DeleteKeyLocationByIds) // 批量删除KeyLocation
		keyLocationRouter.PUT("updateKeyLocation", keyLocationApi.UpdateKeyLocation)    // 更新KeyLocation
	}
	{
		keyLocationRouterWithoutRecord.GET("findKeyLocation", keyLocationApi.FindKeyLocation)        // 根据ID获取KeyLocation
		keyLocationRouterWithoutRecord.POST("getKeyLocationList", keyLocationApi.GetKeyLocationList)  // 获取KeyLocation列表
	}
}
