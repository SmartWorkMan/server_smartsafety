package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClockRouter struct {
}

// InitClockRouter 初始化 Clock 路由信息
func (s *ClockRouter) InitClockRouter(Router *gin.RouterGroup) {
	clockRouter := Router.Group("clock").Use(middleware.OperationRecord())
	clockRouterWithoutRecord := Router.Group("clock")
	var clockApi = v1.ApiGroupApp.SafetyApiGroup.ClockApi
	{
		clockRouter.POST("app/createClock", clockApi.CreateClock)   // 新建Clock
		clockRouter.POST("app/queryClock", clockApi.QueryClock)   // 新建Clock
		clockRouter.DELETE("deleteClock", clockApi.DeleteClock) // 删除Clock
		clockRouter.DELETE("deleteClockByIds", clockApi.DeleteClockByIds) // 批量删除Clock
		//clockRouter.PUT("updateClock", clockApi.UpdateClock)    // 更新Clock
	}
	{
		clockRouterWithoutRecord.GET("findClock", clockApi.FindClock)        // 根据ID获取Clock
		clockRouterWithoutRecord.POST("getTodayClockList", clockApi.GetTodayClockList)
		clockRouterWithoutRecord.GET("getOnDutyNum", clockApi.GetOnDutyNum)
		clockRouterWithoutRecord.POST("getHistoryClockList", clockApi.GetHistoryClockList)
	}
}
