package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type SafetyFactoryApi struct {
}

// CreateSafetyFactory 创建SafetyFactory
// @Tags SafetyFactory
// @Summary 创建SafetyFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.SafetyFactory true "创建SafetyFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /safetyFactory/createSafetyFactory [post]
func (safetyFactoryApi *SafetyFactoryApi) CreateSafetyFactory(c *gin.Context) {
	var safetyFactory safetyReq.SafetyFactoryJson
	_ = c.ShouldBindJSON(&safetyFactory)

	if err := safetyFactoryService.CreateSafetyFactory(factoryJson2Factory(safetyFactory)); err != nil {
        global.GVA_LOG.Error("创建工厂失败!", zap.Error(err))
		response.FailWithMessage("创建工厂失败", c)
	} else {
		//创建巡检区域根节点
		var area safety.Area
		area.FactoryName = safetyFactory.FactoryName
		area.AreaName = safetyFactory.FactoryName
		area.ParentId = commval.AreaRootParentId
		err, _ = areaService.CreateArea(area)
		if err != nil {
			global.GVA_LOG.Error("创建工厂失败!创建巡检区域根节点失败!", zap.Error(err))
			response.FailWithMessage("创建工厂失败!创建巡检区域根节点失败!", c)
			return
		}

		response.OkWithMessage("创建工厂成功", c)
	}
}

func factoryJson2Factory(factoryJson safetyReq.SafetyFactoryJson) safety.SafetyFactory {
	var factory safety.SafetyFactory
	factory = factoryJson.SafetyFactory
	if len(factoryJson.ProvinceCity) == 0 {
		factory.City = ""
	} else {
		factory.City = ""
		for i := 0; i < len(factoryJson.ProvinceCity); i++ {
			if i != len(factoryJson.ProvinceCity) - 1 {
				factory.City += factoryJson.ProvinceCity[i] + ","
			} else {
				factory.City += factoryJson.ProvinceCity[i]
			}
		}
	}

	return factory
}

func factory2FactoryJson (factory safety.SafetyFactory) safetyReq.SafetyFactoryJson {
	var factoryJson safetyReq.SafetyFactoryJson
	factoryJson.SafetyFactory = factory
	factoryJson.City = ""
	factoryJson.ProvinceCity = strings.Split(factory.City, ",")

	return factoryJson
}

func factoryList2FactoryJsonList (factoryList []safety.SafetyFactory) []safetyReq.SafetyFactoryJson {
	var factoryJsonList []safetyReq.SafetyFactoryJson
	for _, factory := range factoryList {
		var factoryJson safetyReq.SafetyFactoryJson
		factoryJson.SafetyFactory = factory
		factoryJson.City = ""
		factoryJson.ProvinceCity = strings.Split(factory.City, ",")
		factoryJsonList = append(factoryJsonList, factoryJson)
	}
	return factoryJsonList
}

// DeleteSafetyFactory 删除SafetyFactory
// @Tags SafetyFactory
// @Summary 删除SafetyFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.SafetyFactory true "删除SafetyFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /safetyFactory/deleteSafetyFactory [delete]
func (safetyFactoryApi *SafetyFactoryApi) DeleteSafetyFactory(c *gin.Context) {
	var safetyFactory safety.SafetyFactory
	_ = c.ShouldBindJSON(&safetyFactory)

	var queryFac safety.SafetyFactory
	var err error
	if err, queryFac = safetyFactoryService.GetSafetyFactory(safetyFactory.ID); err != nil {
		global.GVA_LOG.Error("删除工厂失败!", zap.Error(err))
		response.FailWithMessage("删除工厂失败", c)
		return
	}

	if err := safetyFactoryService.DeleteSafetyFactory(safetyFactory); err != nil {
        global.GVA_LOG.Error("删除工厂失败!", zap.Error(err))
		response.FailWithMessage("删除工厂失败", c)
        return
	} else {
		//删除工厂管理员和维保管理员
		err = userService.DeleteUserByFactoryName(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除工厂管理员和维保管理员失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除工厂管理员和维保管理员失败!", c)
		}

		//删除巡检员
		err = inspectorService.DeleteInspectorByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除巡检员失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除巡检员失败!", c)
		}

		//删除巡检区域
		var area safety.Area
		area.FactoryName = queryFac.FactoryName
		err = areaService.DeleteAreaByFactoryName(area)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除工厂巡检区域失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除工厂巡检区域失败!", c)
		}

		//删除巡检事项
		var item safety.Item
		item.FactoryName = queryFac.FactoryName
		err = itemService.DeleteItemByFactoryName(item)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除工厂巡检事项失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除工厂巡检事项失败!", c)
		}

		//删除打卡记录
		err = clockService.DeleteClockByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除打卡记录失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除打卡记录失败!", c)
		}

		//删除通知
		err = noticeService.DeleteNoticeByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除通知失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除通知失败!", c)
		}

		//删除报告
		err = reportService.DeleteReportByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除报告失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除报告失败!", c)
		}

		//删除培训
		err = trainingService.DeleteTrainingByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除培训失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除培训失败!", c)
		}

		//删除应急预案
		err = planService.DeletePlanByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除应急预案失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除应急预案失败!", c)
		}

		//删除基础信息
		err = basicInfoService.DeleteBasicInfoByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除基础信息失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除基础信息失败!", c)
		}

		//删除建筑信息
		err = buildingInfoService.DeleteBuildingInfoByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除建筑信息失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除建筑信息失败!", c)
		}

		//删除重点部位
		err = keyLocationService.DeleteKeyLocationByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除重点部位失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除重点部位失败!", c)
		}

		//删除重点部位库
		err = locationLibraryService.DeleteLocationLibraryByFactory(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除重点部位库失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除重点部位库失败!", c)
		}

		//删除巡检任务记录
		err = taskService.DeleteTaskHistoryByFactoryName(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除工厂巡检任务记录失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除工厂巡检任务记录失败!", c)
		}

		//删除巡检任务
		err = taskService.DeleteTaskByFactoryName(queryFac.FactoryName)
		if err != nil {
			global.GVA_LOG.Error( "删除工厂失败!删除工厂巡检任务失败!", zap.Error(err))
			response.FailWithMessage("删除工厂失败!删除工厂巡检任务失败!", c)
		} else {
			response.OkWithMessage("删除工厂成功", c)
		}
	}
}

// DeleteSafetyFactoryByIds 批量删除SafetyFactory
// @Tags SafetyFactory
// @Summary 批量删除SafetyFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SafetyFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /safetyFactory/deleteSafetyFactoryByIds [delete]
func (safetyFactoryApi *SafetyFactoryApi) DeleteSafetyFactoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := safetyFactoryService.DeleteSafetyFactoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSafetyFactory 更新SafetyFactory
// @Tags SafetyFactory
// @Summary 更新SafetyFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.SafetyFactory true "更新SafetyFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /safetyFactory/updateSafetyFactory [put]
func (safetyFactoryApi *SafetyFactoryApi) UpdateSafetyFactory(c *gin.Context) {
	var safetyFactory safety.SafetyFactory
	_ = c.ShouldBindJSON(&safetyFactory)
	if err := safetyFactoryService.UpdateSafetyFactory(safetyFactory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Router /safetyFactory/updateFactoryLatLng [put]
func (safetyFactoryApi *SafetyFactoryApi) UpdateFactoryLatLng(c *gin.Context) {
	var safetyFactory safetyReq.SafetyFactoryJson
	_ = c.ShouldBindJSON(&safetyFactory)
	if err := safetyFactoryService.UpdateFactoryLatLng(factoryJson2Factory(safetyFactory)); err != nil {
		global.GVA_LOG.Error("更新经纬度失败!", zap.Error(err))
		response.FailWithMessage("更新经纬度失败", c)
	} else {
		response.OkWithMessage("更新经纬度成功", c)
	}
}

// FindSafetyFactory 用id查询SafetyFactory
// @Tags SafetyFactory
// @Summary 用id查询SafetyFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.SafetyFactory true "用id查询SafetyFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /safetyFactory/findSafetyFactory [get]
func (safetyFactoryApi *SafetyFactoryApi) FindSafetyFactory(c *gin.Context) {
	var safetyFactory safety.SafetyFactory
	_ = c.ShouldBindQuery(&safetyFactory)
	if err, resafetyFactory := safetyFactoryService.GetSafetyFactory(safetyFactory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resafetyFactory": resafetyFactory}, c)
	}
}

// @Router /safetyFactory/querySafetyFactory [post]
func (safetyFactoryApi *SafetyFactoryApi) QuerySafetyFactory(c *gin.Context) {
	var safetyFactory safety.SafetyFactory
	_ = c.ShouldBindJSON(&safetyFactory)
	if err, resafetyFactory := safetyFactoryService.QuerySafetyFactory(safetyFactory); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(factory2FactoryJson(resafetyFactory), "获取成功", c)
	}
}


// GetSafetyFactoryList 分页获取SafetyFactory列表
// @Tags SafetyFactory
// @Summary 分页获取SafetyFactory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.SafetyFactorySearch true "分页获取SafetyFactory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /safetyFactory/getSafetyFactoryList [post]
func (safetyFactoryApi *SafetyFactoryApi) GetSafetyFactoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := safetyFactoryService.GetSafetyFactoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     factoryList2FactoryJsonList(list.([]safety.SafetyFactory)),
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
