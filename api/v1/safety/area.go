package safety

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type AreaApi struct {
}

type AreaListRes struct {
	AreaId    int            `json:"areaId"`
	AreaName  string         `json:"areaName"`
	AreaStatus  string         `json:"areaStatus"`
	Children  []AreaListRes  `json:"children"`
}


// CreateArea 创建Area
// @Tags Area
// @Summary 创建Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Area true "创建Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/createArea [post]
func (areaApi *AreaApi) CreateArea(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("创建巡检区域失败!", zap.Error(err))
		response.FailWithMessage("创建巡检区域失败", c)
		return
	}

	var area safety.Area
	_ = c.ShouldBindJSON(&area)
	if area.ParentId == 0 {
		global.GVA_LOG.Error("创建巡检区域失败!父节点ID为空!")
		response.FailWithMessage("创建巡检区域失败!父节点ID为空!", c)
		return
	}
	area.FactoryName = curUser.FactoryName
	if err, areaId := areaService.CreateArea(area); err != nil {
        global.GVA_LOG.Error("创建巡检区域失败!", zap.Error(err))
		response.FailWithMessage("创建巡检区域失败", c)
	} else {
		//response.OkWithMessage("创建巡检区域成功", c)
		response.OkWithData(gin.H{"areaId": areaId}, c)
	}
}

// DeleteArea 删除Area
// @Tags Area
// @Summary 删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Area true "删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /area/deleteArea [delete]
func (areaApi *AreaApi) DeleteArea(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("删除巡检区域失败!", zap.Error(err))
		response.FailWithMessage("删除巡检区域失败", c)
		return
	}

	var area safety.Area
	_ = c.ShouldBindJSON(&area)
	area.FactoryName = curUser.FactoryName
	if err := areaService.DeleteArea(area); err != nil {
        global.GVA_LOG.Error("删除巡检区域失败!", zap.Error(err))
		response.FailWithMessage("删除巡检区域失败", c)
        return
	} else {
		var item safety.Item
		item.FactoryName = curUser.FactoryName
		item.AreaId = area.ID
		err = itemService.DeleteItemByAreaId(item)
		if err != nil {
			global.GVA_LOG.Error("删除巡检区域下的巡检事项失败!", zap.Error(err))
			response.FailWithMessage("删除巡检区域的巡检事项失败", c)
		} else {
			err = taskService.DeleteTaskByAreaId(area.ID)
			if err != nil {
				global.GVA_LOG.Error("删除巡检区域下的巡检任务失败!", zap.Error(err))
				response.FailWithMessage("删除巡检区域下的巡检任务失败", c)
			} else {
				response.OkWithMessage("删除巡检区域成功", c)
			}
		}
	}
}

// DeleteAreaByIds 批量删除Area
// @Tags Area
// @Summary 批量删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /area/deleteAreaByIds [delete]
func (areaApi *AreaApi) DeleteAreaByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := areaService.DeleteAreaByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArea 更新Area
// @Tags Area
// @Summary 更新Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Area true "更新Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /area/updateArea [put]
func (areaApi *AreaApi) UpdateArea(c *gin.Context) {
	selfUserInfo := utils.GetUserInfo(c)
	var curUser *system.SysUser
	var err error
	err, curUser = userService.FindUserById(int(selfUserInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("更新巡检区域失败!", zap.Error(err))
		response.FailWithMessage("更新巡检区域失败", c)
		return
	}

	if curUser.FactoryName == "" {
		global.GVA_LOG.Error("更新巡检区域失败!当前用户工厂名称为空")
		response.FailWithMessage("更新巡检区域失败!当前用户工厂名称为空", c)
		return
	}

	var area safety.Area
	_ = c.ShouldBindJSON(&area)
	area.FactoryName = curUser.FactoryName
	if err := areaService.UpdateArea(area); err != nil {
        global.GVA_LOG.Error("更新巡检区域失败!", zap.Error(err))
		response.FailWithMessage("更新巡检区域失败", c)
	} else {
		response.OkWithMessage("更新巡检区域成功", c)
	}
}

// FindArea 用id查询Area
// @Tags Area
// @Summary 用id查询Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Area true "用id查询Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /area/findArea [get]
func (areaApi *AreaApi) FindArea(c *gin.Context) {
	var area safety.Area
	_ = c.ShouldBindQuery(&area)
	if err, rearea := areaService.GetArea(area.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearea": rearea}, c)
	}
}

// GetAreaList 分页获取Area列表
// @Tags Area
// @Summary 分页获取Area列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.AreaSearch true "分页获取Area列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/getAreaList [get]
func (areaApi *AreaApi) GetAreaList(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检区域列表失败", c)
		return
	}

	err, id := areaService.GetRootAreaId(curUser.FactoryName)
	if err != nil {
		global.GVA_LOG.Error("获取巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检区域列表失败", c)
		return
	}

	var listRes AreaListRes
	listRes.AreaName = curUser.FactoryName
	listRes.AreaId = int(id)

	var reqArea safety.Area
	reqArea.FactoryName = curUser.FactoryName
	reqArea.ID = id
	err, listRes.Children = areaApi.GetAreaChildren(reqArea)
	if err != nil {
		global.GVA_LOG.Error("获取巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检区域列表失败", c)
	} else {
		response.OkWithDetailed(listRes, "获取巡检区域列表成功", c)
	}
}

func (areaApi *AreaApi) GetAreaChildren(area safety.Area)(error, []AreaListRes) {
	if areaService.IsLeafNode(area) {
		return nil, nil
	} else {
		err, areaChildren := areaService.GetAreaByParentId(area.FactoryName, int(area.ID))
		if err != nil {
			return err, nil
		} else {
			var childrenRes []AreaListRes
			for _, child := range areaChildren {
				var childRes AreaListRes
				childRes.AreaId = int(child.ID)
				childRes.AreaName = child.AreaName
				err, childRes.Children = areaApi.GetAreaChildren(child)
				if err != nil {
					return err, nil
				}
				childrenRes = append(childrenRes, childRes)
			}

			return nil, childrenRes
		}
	}
}

func (areaApi *AreaApi) GetLeafAreaIdList(area safety.Area)(error, []uint) {
	var areaIdList []uint
	if areaService.IsLeafNode(area) {
		areaIdList = append(areaIdList, area.ID)
		return nil, areaIdList
	} else {
		err, areaChildren := areaService.GetAreaByParentId(area.FactoryName, int(area.ID))
		if err != nil {
			return err, nil
		} else {
			for _, child := range areaChildren {
				if areaService.IsLeafNode(child) {
					areaIdList = append(areaIdList, child.ID)
				} else {
					err, list := areaApi.GetLeafAreaIdList(child)
					if err != nil {
						return err, nil
					}
					areaIdList = append(areaIdList, list...)
				}
			}
		}
	}

	return nil, areaIdList
}

// @Router /area/app/getAreaListByInspector [post]
func (areaApi *AreaApi) GetAreaListByInspector(c *gin.Context) {
	//检查输入巡检员用户名
	var pageInfo safetyReq.TaskSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.InspectorUsername == "" {
		global.GVA_LOG.Error("获取巡检员巡检区域列表失败!巡检员用户名为空!")
		response.FailWithMessage("获取巡检员巡检区域列表失败!巡检员用户名为空!", c)
		return
	}

	//获取工厂名
	err, inspector := inspectorService.GetInspectorByUserName(pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检区域列表失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检员巡检区域列表失败!请输入正确的巡检员用户名!", c)
		return
	}

	//检查巡检员是否有巡检任务
	pageInfo.FactoryName = inspector.FactoryName
	curDate := time.Now().Format("2006-01-02")
	pageInfo.PlanInspectionDate = curDate
	err, _, total := taskService.GetTaskListByInspectorForAreaTree(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员巡检区域列表失败", c)
		return
	}

	if total == 0 {
		global.GVA_LOG.Error(fmt.Sprintf("获取巡检员巡检区域列表失败!巡检员: %s 今天没有巡检任务", pageInfo.InspectorUsername))
		response.FailWithMessage(fmt.Sprintf("获取巡检员巡检区域列表失败!巡检员: %s 今天没有巡检任务", pageInfo.InspectorUsername), c)
		return
	}

	//获取工厂全部区域列表
	err, id := areaService.GetRootAreaId(inspector.FactoryName)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员巡检区域列表失败", c)
		return
	}

	var listRes AreaListRes
	listRes.AreaName = inspector.FactoryName
	listRes.AreaId = int(id)

	var reqArea safety.Area
	reqArea.FactoryName = inspector.FactoryName
	reqArea.ID = id
	err, listRes.Children, listRes.AreaStatus = areaApi.GetAreaChildrenByInspector(reqArea, pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检区域列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员巡检区域列表失败", c)
		return
	} else {
		response.OkWithDetailed(listRes, "获取巡检员巡检区域列表成功", c)
	}
}

func (areaApi *AreaApi) GetAreaChildrenByInspector(area safety.Area, inspectorUserName string)(error, []AreaListRes, string) {
	if areaService.IsLeafNode(area) {
		err, status := areaApi.getAreaStatus(area, inspectorUserName)
		if err != nil {
			return err, nil, ""
		}
		return nil, nil, status
	} else {
		err, hasTask := areaApi.isAreaHasInspectorTask(area, inspectorUserName)
		if err != nil {
			return err, nil, ""
		}
		if !hasTask {
			return nil, nil, ""
		}

		err, areaChildren := areaService.GetAreaByParentId(area.FactoryName, int(area.ID))
		if err != nil {
			return err, nil, ""
		} else {
			var childrenRes []AreaListRes
			childrenStatus := ""
			for _, child := range areaChildren {

				var tempArea safety.Area
				tempArea.FactoryName = area.FactoryName
				tempArea.ID = child.ID
				err, hasTask := areaApi.isAreaHasInspectorTask(tempArea, inspectorUserName)
				if err != nil {
					return err, nil, ""
				}
				if !hasTask {
					continue
				}

				var childRes AreaListRes
				childRes.AreaId = int(child.ID)
				childRes.AreaName = child.AreaName
				err, childRes.Children, childRes.AreaStatus = areaApi.GetAreaChildrenByInspector(child, inspectorUserName)
				if err != nil {
					return err, nil, ""
				}
				childrenRes = append(childrenRes, childRes)
				if childRes.AreaStatus == "异常" {
					childrenStatus = "异常"
				} else if childRes.AreaStatus == "未开始" && childrenStatus != "异常"{
					childrenStatus = "未开始"
				} else if childRes.AreaStatus == "正常" && childrenStatus == "" {
					childrenStatus = "正常"
				}
			}

			return nil, childrenRes, childrenStatus
		}
	}
}

func (areaApi *AreaApi) getAreaStatus(area safety.Area, inspectorUserName string)(error, string) {
	err, hasTask := areaApi.isAreaHasInspectorTask(area, inspectorUserName)
	if err != nil {
		return err, ""
	}
	if !hasTask {
		return nil, ""
	}

	var pageInfo safetyReq.TaskSearch
	pageInfo.FactoryName = area.FactoryName
	curDate := time.Now().Format("2006-01-02")
	pageInfo.PlanInspectionDate = curDate
	pageInfo.InspectorUsername = inspectorUserName
	pageInfo.AreaId = area.ID
	var taskList interface{}
	if err, taskList, _ = taskService.GetTaskListByAreaForInspector(pageInfo); err != nil {
		return err, ""
	}

	notStartCount := 0
	endCount := 0
	for _, task := range taskList.([]safety.Task) {
		if task.TaskStatus != commval.TaskStatusNotStart && task.TaskStatus != commval.TaskStatusEnd {
			return nil, "异常"
		} else if task.TaskStatus == commval.TaskStatusNotStart {
			notStartCount ++
		} else {
			endCount ++
		}
	}

	if notStartCount > 0 {
		return nil, "未开始"
	} else {
		return nil, "正常"
	}
}

func (areaApi *AreaApi) isAreaHasInspectorTask(area safety.Area, inspectorUserName string)(error, bool) {
	err, areaIdList := areaApi.GetLeafAreaIdList(area)
	if err != nil {
		return err, false
	}

	for _, areaId := range areaIdList {
		var pageInfo safetyReq.TaskSearch
		pageInfo.FactoryName = area.FactoryName
		curDate := time.Now().Format("2006-01-02")
		pageInfo.PlanInspectionDate = curDate
		pageInfo.InspectorUsername = inspectorUserName
		pageInfo.AreaId = areaId
		err, _, total := taskService.GetTaskListByAreaForInspector(pageInfo)
		if err != nil {
			return err, false
		}
		if total > 0 {
			return nil, true
		}
	}

	return nil, false
}