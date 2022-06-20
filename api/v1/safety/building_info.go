package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BuildingInfoApi struct {
}


// CreateBuildingInfo 创建BuildingInfo
// @Tags BuildingInfo
// @Summary 创建BuildingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BuildingInfo true "创建BuildingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInfo/createBuildingInfo [post]
func (buildingInfoApi *BuildingInfoApi) CreateBuildingInfo(c *gin.Context) {
	var buildingInfo safety.BuildingInfo
	_ = c.ShouldBindJSON(&buildingInfo)
	if err := buildingInfoService.CreateBuildingInfo(buildingInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBuildingInfo 删除BuildingInfo
// @Tags BuildingInfo
// @Summary 删除BuildingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BuildingInfo true "删除BuildingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildingInfo/deleteBuildingInfo [delete]
func (buildingInfoApi *BuildingInfoApi) DeleteBuildingInfo(c *gin.Context) {
	var buildingInfo safety.BuildingInfo
	_ = c.ShouldBindJSON(&buildingInfo)
	if err := buildingInfoService.DeleteBuildingInfo(buildingInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBuildingInfoByIds 批量删除BuildingInfo
// @Tags BuildingInfo
// @Summary 批量删除BuildingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /buildingInfo/deleteBuildingInfoByIds [delete]
func (buildingInfoApi *BuildingInfoApi) DeleteBuildingInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := buildingInfoService.DeleteBuildingInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBuildingInfo 更新BuildingInfo
// @Tags BuildingInfo
// @Summary 更新BuildingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BuildingInfo true "更新BuildingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildingInfo/updateBuildingInfo [put]
func (buildingInfoApi *BuildingInfoApi) UpdateBuildingInfo(c *gin.Context) {
	var buildingInfo safety.BuildingInfo
	_ = c.ShouldBindJSON(&buildingInfo)
	if err := buildingInfoService.UpdateBuildingInfo(buildingInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBuildingInfo 用id查询BuildingInfo
// @Tags BuildingInfo
// @Summary 用id查询BuildingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.BuildingInfo true "用id查询BuildingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildingInfo/getBuildingInfo [post]
func (buildingInfoApi *BuildingInfoApi) GetBuildingInfo(c *gin.Context) {
	var buildingInfo safety.BuildingInfo
	_ = c.ShouldBindJSON(&buildingInfo)
	if err, rebuildingInfo := buildingInfoService.GetBuildingInfo(buildingInfo.FactoryName); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rebuildingInfo, "获取成功", c)
	}
}

// GetBuildingInfoList 分页获取BuildingInfo列表
// @Tags BuildingInfo
// @Summary 分页获取BuildingInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.BuildingInfoSearch true "分页获取BuildingInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildingInfo/getBuildingInfoList [post]
func (buildingInfoApi *BuildingInfoApi) GetBuildingInfoList(c *gin.Context) {
	var pageInfo safetyReq.BuildingInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := buildingInfoService.GetBuildingInfoInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
