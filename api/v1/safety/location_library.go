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

type LocationLibraryApi struct {
}


// CreateLocationLibrary 创建LocationLibrary
// @Tags LocationLibrary
// @Summary 创建LocationLibrary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.LocationLibrary true "创建LocationLibrary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationLibrary/createLocationLibrary [post]
func (locationLibraryApi *LocationLibraryApi) CreateLocationLibrary(c *gin.Context) {
	var locationLibrary safety.LocationLibrary
	_ = c.ShouldBindJSON(&locationLibrary)
	if locationLibrary.FactoryName == "" || locationLibrary.LocationName == "" {
		global.GVA_LOG.Error("创建失败!输入不能为空!")
		response.FailWithMessage("创建失败!输入不能为空!", c)
		return
	}
	if err := locationLibraryService.CreateLocationLibrary(locationLibrary); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLocationLibrary 删除LocationLibrary
// @Tags LocationLibrary
// @Summary 删除LocationLibrary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.LocationLibrary true "删除LocationLibrary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationLibrary/deleteLocationLibrary [delete]
func (locationLibraryApi *LocationLibraryApi) DeleteLocationLibrary(c *gin.Context) {
	var locationLibrary safety.LocationLibrary
	_ = c.ShouldBindJSON(&locationLibrary)
	if err := locationLibraryService.DeleteLocationLibrary(locationLibrary); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLocationLibraryByIds 批量删除LocationLibrary
// @Tags LocationLibrary
// @Summary 批量删除LocationLibrary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocationLibrary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /locationLibrary/deleteLocationLibraryByIds [delete]
func (locationLibraryApi *LocationLibraryApi) DeleteLocationLibraryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := locationLibraryService.DeleteLocationLibraryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLocationLibrary 更新LocationLibrary
// @Tags LocationLibrary
// @Summary 更新LocationLibrary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.LocationLibrary true "更新LocationLibrary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /locationLibrary/updateLocationLibrary [put]
func (locationLibraryApi *LocationLibraryApi) UpdateLocationLibrary(c *gin.Context) {
	var locationLibrary safety.LocationLibrary
	_ = c.ShouldBindJSON(&locationLibrary)
	if locationLibrary.ID == 0 || locationLibrary.LocationName == "" {
		global.GVA_LOG.Error("创建失败!输入不能为空!")
		response.FailWithMessage("创建失败!输入不能为空!", c)
		return
	}
	if err := locationLibraryService.UpdateLocationLibrary(locationLibrary); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLocationLibrary 用id查询LocationLibrary
// @Tags LocationLibrary
// @Summary 用id查询LocationLibrary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.LocationLibrary true "用id查询LocationLibrary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /locationLibrary/findLocationLibrary [get]
func (locationLibraryApi *LocationLibraryApi) FindLocationLibrary(c *gin.Context) {
	var locationLibrary safety.LocationLibrary
	_ = c.ShouldBindQuery(&locationLibrary)
	if err, relocationLibrary := locationLibraryService.GetLocationLibrary(locationLibrary.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relocationLibrary": relocationLibrary}, c)
	}
}

// GetLocationLibraryList 分页获取LocationLibrary列表
// @Tags LocationLibrary
// @Summary 分页获取LocationLibrary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.LocationLibrarySearch true "分页获取LocationLibrary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationLibrary/getLocationLibraryList [post]
func (locationLibraryApi *LocationLibraryApi) GetLocationLibraryList(c *gin.Context) {
	var pageInfo safetyReq.LocationLibrarySearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不为空!")
		response.FailWithMessage("获取失败!工厂名称不为空!", c)
		return
	}
	if err, list, total := locationLibraryService.GetLocationLibraryInfoList(pageInfo); err != nil {
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
