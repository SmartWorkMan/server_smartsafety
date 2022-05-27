package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InspectorRouter struct {
}

// InitInspectorRouter 初始化 Inspector 路由信息
func (s *InspectorRouter) InitInspectorRouter(Router *gin.RouterGroup) {
	inspectorRouter := Router.Group("inspector").Use(middleware.OperationRecord())
	inspectorRouterWithoutRecord := Router.Group("inspector")
	var inspectorApi = v1.ApiGroupApp.SafetyApiGroup.InspectorApi
	{
		inspectorRouter.POST("createInspector", inspectorApi.CreateInspector)   // 新建Inspector
		inspectorRouter.DELETE("deleteInspector", inspectorApi.DeleteInspector) // 删除Inspector
		inspectorRouter.DELETE("deleteInspectorByIds", inspectorApi.DeleteInspectorByIds) // 批量删除Inspector
		inspectorRouter.PUT("updateInspector", inspectorApi.UpdateInspector)    // 更新Inspector
		inspectorRouter.POST("login", inspectorApi.Login)                       // 登录小程序
	}
	{
		inspectorRouterWithoutRecord.GET("findInspector", inspectorApi.FindInspector)        // 根据ID获取Inspector
		inspectorRouterWithoutRecord.POST("getInspectorList", inspectorApi.GetInspectorList)  // 获取Inspector列表
	}
}
