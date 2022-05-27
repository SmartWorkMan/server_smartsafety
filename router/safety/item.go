package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ItemRouter struct {
}

// InitItemRouter 初始化 Item 路由信息
func (s *ItemRouter) InitItemRouter(Router *gin.RouterGroup) {
	itemRouter := Router.Group("item").Use(middleware.OperationRecord())
	itemRouterWithoutRecord := Router.Group("item")
	var itemApi = v1.ApiGroupApp.SafetyApiGroup.ItemApi
	{
		itemRouter.POST("createItem", itemApi.CreateItem)   // 新建Item
		itemRouter.DELETE("deleteItem", itemApi.DeleteItem) // 删除Item
		itemRouter.DELETE("deleteItemByIds", itemApi.DeleteItemByIds) // 批量删除Item
		itemRouter.PUT("updateItem", itemApi.UpdateItem)    // 更新Item
		itemRouter.PUT("enableItem", itemApi.EnableItem)    // 启用Item
		itemRouter.PUT("disableItem", itemApi.DisableItem)    // 禁用Item
	}
	{
		itemRouterWithoutRecord.GET("findItem", itemApi.FindItem)        // 根据ID获取Item
		itemRouterWithoutRecord.POST("getItemList", itemApi.GetItemList)  // 获取Item列表
		itemRouterWithoutRecord.POST("getItemListByAreaId", itemApi.GetItemListByAreaId)  // 按区域获取Item列表
	}
}
