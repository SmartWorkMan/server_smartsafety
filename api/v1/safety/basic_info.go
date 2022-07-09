package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"errors"
)

type BasicInfoApi struct {
}


// CreateBasicInfo 创建BasicInfo
// @Tags BasicInfo
// @Summary 创建BasicInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BasicInfo true "创建BasicInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicInfo/createBasicInfo [post]
func (basicInfoApi *BasicInfoApi) CreateBasicInfo(c *gin.Context) {
	var basicInfo safety.BasicInfo
	_ = c.ShouldBindJSON(&basicInfo)
	if err := basicInfoService.CreateBasicInfo(basicInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBasicInfo 删除BasicInfo
// @Tags BasicInfo
// @Summary 删除BasicInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BasicInfo true "删除BasicInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicInfo/deleteBasicInfo [delete]
func (basicInfoApi *BasicInfoApi) DeleteBasicInfo(c *gin.Context) {
	var basicInfo safety.BasicInfo
	_ = c.ShouldBindJSON(&basicInfo)
	if err := basicInfoService.DeleteBasicInfo(basicInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicInfoByIds 批量删除BasicInfo
// @Tags BasicInfo
// @Summary 批量删除BasicInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /basicInfo/deleteBasicInfoByIds [delete]
func (basicInfoApi *BasicInfoApi) DeleteBasicInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := basicInfoService.DeleteBasicInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBasicInfo 更新BasicInfo
// @Tags BasicInfo
// @Summary 更新BasicInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.BasicInfo true "更新BasicInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicInfo/updateBasicInfo [put]
func (basicInfoApi *BasicInfoApi) UpdateBasicInfo(c *gin.Context) {
	var basicInfo safety.BasicInfo
	_ = c.ShouldBindJSON(&basicInfo)
	if err := basicInfoService.UpdateBasicInfo(basicInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBasicInfo 用id查询BasicInfo
// @Tags BasicInfo
// @Summary 用id查询BasicInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.BasicInfo true "用id查询BasicInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicInfo/getBasicInfo [post]
func (basicInfoApi *BasicInfoApi) GetBasicInfo(c *gin.Context) {
	var basicInfo safety.BasicInfo
	_ = c.ShouldBindJSON(&basicInfo)
	if err, rebasicInfo := basicInfoService.GetBasicInfo(basicInfo.FactoryName); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OkWithDetailed(nil, "获取成功", c)
		} else {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
	} else {
		response.OkWithDetailed(rebasicInfo, "获取成功", c)
	}
}

// GetBasicInfoList 分页获取BasicInfo列表
// @Tags BasicInfo
// @Summary 分页获取BasicInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.BasicInfoSearch true "分页获取BasicInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicInfo/getBasicInfoList [get]
func (basicInfoApi *BasicInfoApi) GetBasicInfoList(c *gin.Context) {
	var pageInfo safetyReq.BasicInfoSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := basicInfoService.GetBasicInfoInfoList(pageInfo); err != nil {
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
