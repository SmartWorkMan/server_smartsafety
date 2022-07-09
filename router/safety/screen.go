package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ScreenRouter struct {
}

// InitTaskRouter 初始化 Task 路由信息
func (s *ScreenRouter) InitScreenRouter(Router *gin.RouterGroup) {
	screenRouterWithoutRecord := Router.Group("screen")
	var taskApi = v1.ApiGroupApp.SafetyApiGroup.TaskApi
	{
		screenRouterWithoutRecord.POST("getNormalTaskCount", taskApi.GetNormalTaskCount)
		screenRouterWithoutRecord.POST("getPendingTaskCount", taskApi.GetPendingTaskCount)
		screenRouterWithoutRecord.POST("getFixedTaskCount", taskApi.GetFixedTaskCount)
		screenRouterWithoutRecord.POST("getNotFixedTaskCount", taskApi.GetNotFixedTaskCount)
		screenRouterWithoutRecord.POST("getTopFailureItems", taskApi.GetTopFailureItems)
		screenRouterWithoutRecord.POST("getFixedStatistics", taskApi.GetFixedStatistics)
		screenRouterWithoutRecord.POST("getWeeklyHealthIndex", taskApi.GetWeeklyHealthIndex)
		screenRouterWithoutRecord.POST("getWeeklyFixedCurve", taskApi.GetWeeklyFixedCurve)
		screenRouterWithoutRecord.POST("getStartInspectInfo", taskApi.GetStartInspectInfo)
	}
}
