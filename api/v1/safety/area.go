package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AreaApi struct {
}

type AreaListRes struct {
	AreaId    int            `json:"areaId"`
	AreaName  string         `json:"areaName"`
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
		global.GVA_LOG.Error("创建巡检区域失败!", zap.Error(err))
		response.FailWithMessage("创建巡检区域失败", c)
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
			global.GVA_LOG.Error("删除巡检区域的巡检事项失败!", zap.Error(err))
			response.FailWithMessage("删除巡检区域的巡检事项失败", c)
		} else {
			response.OkWithMessage("删除巡检区域成功", c)
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
// @Router /area/getAreaList [post]
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

