package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NoticeRouter struct {
}

// InitNoticeRouter 初始化 Notice 路由信息
func (s *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noticeRouter := Router.Group("notice").Use(middleware.OperationRecord())
	noticeRouterWithoutRecord := Router.Group("notice")
	var noticeApi = v1.ApiGroupApp.SafetyApiGroup.NoticeApi
	{
		noticeRouter.POST("createNotice", noticeApi.CreateNotice)   // 新建Notice
		noticeRouter.DELETE("deleteNotice", noticeApi.DeleteNotice) // 删除Notice
		noticeRouter.DELETE("deleteNoticeByIds", noticeApi.DeleteNoticeByIds) // 批量删除Notice
		noticeRouter.PUT("updateNotice", noticeApi.UpdateNotice)    // 更新Notice
		noticeRouter.POST("readNotice", noticeApi.ReadNotice)
	}
	{
		//noticeRouterWithoutRecord.GET("findNotice", noticeApi.FindNotice)        // 根据ID获取Notice
		noticeRouterWithoutRecord.POST("getNoticeList", noticeApi.GetNoticeList)  // 获取Notice列表
		noticeRouterWithoutRecord.POST("getNoticeListForSuperUser", noticeApi.GetNoticeListForSuperUser)  // 获取Notice列表
	}
}
