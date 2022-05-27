package safety

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ItemApi struct {
}


// CreateItem 创建Item
// @Tags Item
// @Summary 创建Item
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Item true "创建Item"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /item/createItem [post]
func (itemApi *ItemApi) CreateItem(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("创建巡检事项失败!", zap.Error(err))
		response.FailWithMessage("创建巡检事项失败", c)
		return
	}

	var item safety.Item
	_ = c.ShouldBindJSON(&item)
	item.FactoryName = curUser.FactoryName
	global.GVA_LOG.Info(fmt.Sprintf("item:%+v", item))

	if item.Period != commval.ItemPeriodDay &&
		item.Period != commval.ItemPeriodWeek &&
		item.Period != commval.ItemPeriodMonth {
		global.GVA_LOG.Error(fmt.Sprintf("创建巡检事项失败!不支持周期%s", item.Period))
		response.FailWithMessage(fmt.Sprintf("创建巡检事项失败!不支持周期%s", item.Period), c)
		return
	}

	var area safety.Area
	area.FactoryName = curUser.FactoryName
	area.ID = item.AreaId
	if !areaService.IsLeafNode(area) {
		global.GVA_LOG.Error("创建巡检事项失败!只能在最底层区域创建巡检事项!", zap.Error(err))
		response.FailWithMessage("创建巡检事项失败!只能在最底层区域创建巡检事项!", c)
		return
	}

	err, areaPath := itemApi.getAreaPath(item.AreaId)
	if err != nil {
		global.GVA_LOG.Error("创建巡检事项失败!", zap.Error(err))
		response.FailWithMessage("创建巡检事项失败", c)
	}
	item.AreaName = areaPath

	if err = itemService.CreateItem(item); err != nil {
        global.GVA_LOG.Error("创建巡检事项失败!", zap.Error(err))
		response.FailWithMessage("创建巡检事项失败", c)
	} else {
		response.OkWithMessage("创建巡检事项成功", c)
	}
}

func (itemApi *ItemApi) getAreaPath(areaId uint)(error, string) {
	areaPath := ""

	for {
		err, area := areaService.GetArea(areaId)
		if err != nil {
			return err, ""
		}

		if areaPath == "" {
			areaPath = area.AreaName
		} else {
			areaPath = area.AreaName + ">" + areaPath
		}

		if area.ParentId == commval.AreaRootParentId {
			break
		}
		areaId = uint(area.ParentId)
	}

	return nil, areaPath
}

// DeleteItem 删除Item
// @Tags Item
// @Summary 删除Item
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Item true "删除Item"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /item/deleteItem [delete]
func (itemApi *ItemApi) DeleteItem(c *gin.Context) {
	var item safety.Item
	_ = c.ShouldBindJSON(&item)
	if err := itemService.DeleteItem(item); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteItemByIds 批量删除Item
// @Tags Item
// @Summary 批量删除Item
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Item"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /item/deleteItemByIds [delete]
func (itemApi *ItemApi) DeleteItemByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := itemService.DeleteItemByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateItem 更新Item
// @Tags Item
// @Summary 更新Item
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Item true "更新Item"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /item/updateItem [put]
func (itemApi *ItemApi) UpdateItem(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("更新巡检事项失败!", zap.Error(err))
		response.FailWithMessage("更新巡检事项失败", c)
		return
	}

	var item safety.Item
	_ = c.ShouldBindJSON(&item)
	item.FactoryName = curUser.FactoryName
	if err = itemService.UpdateItem(item); err != nil {
        global.GVA_LOG.Error("更新巡检事项失败!", zap.Error(err))
		response.FailWithMessage("更新巡检事项失败", c)
	} else {
		response.OkWithMessage("更新巡检事项成功", c)
	}
}

// FindItem 用id查询Item
// @Tags Item
// @Summary 用id查询Item
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Item true "用id查询Item"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /item/findItem [get]
func (itemApi *ItemApi) FindItem(c *gin.Context) {
	var item safety.Item
	_ = c.ShouldBindQuery(&item)
	if err, reitem := itemService.GetItem(item.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reitem": reitem}, c)
	}
}

// GetItemList 分页获取Item列表
// @Tags Item
// @Summary 分页获取Item列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.ItemSearch true "分页获取Item列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /item/getItemList [post]
func (itemApi *ItemApi) GetItemList(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取巡检事项列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检事项列表失败", c)
		return
	}

	var pageInfo safetyReq.ItemSearch
	_ = c.ShouldBindJSON(&pageInfo)
	pageInfo.FactoryName = curUser.FactoryName
	if err, list, total := itemService.GetItemInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取巡检事项列表失败!", zap.Error(err))
        response.FailWithMessage("获取巡检事项列表失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取巡检事项列表成功", c)
    }
}

// @Router /item/getItemListByAreaId [post]
func (itemApi *ItemApi) GetItemListByAreaId(c *gin.Context) {
	//获取当前用户
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取巡检事项列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检事项列表失败", c)
		return
	}

	//序列化reqeust body
	var pageInfo safetyReq.ItemSearch
	_ = c.ShouldBindJSON(&pageInfo)
	pageInfo.FactoryName = curUser.FactoryName

	//获取当前area的所有叶子节点Id
	var leafAreaIdList []uint
	var reqArea safety.Area
	reqArea.FactoryName = curUser.FactoryName
	reqArea.ID = pageInfo.AreaId
	areaApi := new(AreaApi)
	err, leafAreaIdList = areaApi.GetLeafAreaIdList(reqArea)
	if err != nil {
		global.GVA_LOG.Error("获取巡检事项列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检事项列表失败", c)
	}

	global.GVA_LOG.Info(fmt.Sprintf("areaId:%d的所有叶子节点ID:%v", reqArea.ID, leafAreaIdList))
	//获取item list
	var itemList []safety.Item
	var itemTotal int64
	for _, leafAreaId := range leafAreaIdList {
		pageInfo.AreaId = leafAreaId
		if err, list, total := itemService.GetItemInfoListByLeafAreaId(pageInfo); err != nil {
			global.GVA_LOG.Error("获取巡检事项列表失败!", zap.Error(err))
			response.FailWithMessage("获取巡检事项列表失败", c)
		} else {
			itemList = append(itemList, list.([]safety.Item)...)
			itemTotal += total
		}
	}

	response.OkWithDetailed(response.PageResult{
		List:     itemList,
		Total:    itemTotal,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取巡检事项列表成功", c)
}

// @Router /item/enableItem [put]
func (itemApi *ItemApi) EnableItem(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("启用巡检事项失败!", zap.Error(err))
		response.FailWithMessage("启用巡检事项失败", c)
		return
	}

	var item safety.Item
	_ = c.ShouldBindJSON(&item)
	item.FactoryName = curUser.FactoryName
	if err = itemService.EnableItem(item); err != nil {
		global.GVA_LOG.Error("启用巡检事项失败!", zap.Error(err))
		response.FailWithMessage("启用巡检事项失败", c)
	} else {
		response.OkWithMessage("启用巡检事项成功", c)
	}
}

// @Router /item/disableItem [put]
func (itemApi *ItemApi) DisableItem(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("禁用巡检事项失败!", zap.Error(err))
		response.FailWithMessage("禁用巡检事项失败", c)
		return
	}

	var item safety.Item
	_ = c.ShouldBindJSON(&item)
	item.FactoryName = curUser.FactoryName
	if err = itemService.DisableItem(item); err != nil {
		global.GVA_LOG.Error("禁用巡检事项失败!", zap.Error(err))
		response.FailWithMessage("禁用巡检事项失败", c)
	} else {
		response.OkWithMessage("禁用巡检事项成功", c)
	}
}