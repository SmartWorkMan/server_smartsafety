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

type KeyLocationApi struct {
}


// CreateKeyLocation 创建KeyLocation
// @Tags KeyLocation
// @Summary 创建KeyLocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.KeyLocation true "创建KeyLocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /keyLocation/createKeyLocation [post]
func (keyLocationApi *KeyLocationApi) CreateKeyLocation(c *gin.Context) {
	var keyLocation safety.KeyLocation
	_ = c.ShouldBindJSON(&keyLocation)
	if err := keyLocationService.CreateKeyLocation(keyLocation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKeyLocation 删除KeyLocation
// @Tags KeyLocation
// @Summary 删除KeyLocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.KeyLocation true "删除KeyLocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keyLocation/deleteKeyLocation [delete]
func (keyLocationApi *KeyLocationApi) DeleteKeyLocation(c *gin.Context) {
	var keyLocation safety.KeyLocation
	_ = c.ShouldBindJSON(&keyLocation)
	if err := keyLocationService.DeleteKeyLocation(keyLocation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKeyLocationByIds 批量删除KeyLocation
// @Tags KeyLocation
// @Summary 批量删除KeyLocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KeyLocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /keyLocation/deleteKeyLocationByIds [delete]
func (keyLocationApi *KeyLocationApi) DeleteKeyLocationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := keyLocationService.DeleteKeyLocationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKeyLocation 更新KeyLocation
// @Tags KeyLocation
// @Summary 更新KeyLocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.KeyLocation true "更新KeyLocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /keyLocation/updateKeyLocation [put]
func (keyLocationApi *KeyLocationApi) UpdateKeyLocation(c *gin.Context) {
	var keyLocation safety.KeyLocation
	_ = c.ShouldBindJSON(&keyLocation)
	if err := keyLocationService.UpdateKeyLocation(keyLocation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKeyLocation 用id查询KeyLocation
// @Tags KeyLocation
// @Summary 用id查询KeyLocation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.KeyLocation true "用id查询KeyLocation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /keyLocation/findKeyLocation [get]
func (keyLocationApi *KeyLocationApi) FindKeyLocation(c *gin.Context) {
	var keyLocation safety.KeyLocation
	_ = c.ShouldBindQuery(&keyLocation)
	if err, rekeyLocation := keyLocationService.GetKeyLocation(keyLocation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekeyLocation": rekeyLocation}, c)
	}
}

// GetKeyLocationList 分页获取KeyLocation列表
// @Tags KeyLocation
// @Summary 分页获取KeyLocation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.KeyLocationSearch true "分页获取KeyLocation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /keyLocation/getKeyLocationList [post]
func (keyLocationApi *KeyLocationApi) GetKeyLocationList(c *gin.Context) {
	var pageInfo safetyReq.KeyLocationSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := keyLocationService.GetKeyLocationInfoList(pageInfo); err != nil {
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
