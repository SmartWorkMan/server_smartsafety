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

type TecentApi struct {
}


// CreateTecent 创建Tecent
// @Tags Tecent
// @Summary 创建Tecent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Tecent true "创建Tecent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecent/createTecent [post]
func (tecentApi *TecentApi) CreateTecent(c *gin.Context) {
	var tecent safety.Tecent
	_ = c.ShouldBindJSON(&tecent)
	if err := tecentService.CreateTecent(tecent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTecent 删除Tecent
// @Tags Tecent
// @Summary 删除Tecent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Tecent true "删除Tecent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecent/deleteTecent [delete]
func (tecentApi *TecentApi) DeleteTecent(c *gin.Context) {
	var tecent safety.Tecent
	_ = c.ShouldBindJSON(&tecent)
	if err := tecentService.DeleteTecent(tecent); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTecentByIds 批量删除Tecent
// @Tags Tecent
// @Summary 批量删除Tecent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tecent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tecent/deleteTecentByIds [delete]
func (tecentApi *TecentApi) DeleteTecentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := tecentService.DeleteTecentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTecent 更新Tecent
// @Tags Tecent
// @Summary 更新Tecent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Tecent true "更新Tecent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecent/updateTecent [put]
func (tecentApi *TecentApi) UpdateTecent(c *gin.Context) {
	var tecent safety.Tecent
	_ = c.ShouldBindJSON(&tecent)
	if err := tecentService.UpdateTecent(tecent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetTecentList 分页获取Tecent列表
// @Tags Tecent
// @Summary 分页获取Tecent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.TecentSearch true "分页获取Tecent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tencent/getTencent [get]
func (tecentApi *TecentApi) GetTecent(c *gin.Context) {
	var tencent safetyReq.TecentSearch
	_ = c.ShouldBindJSON(&tencent)
	if tencent.TencentType == "" || (tencent.TencentType != "cos" && tencent.TencentType != "map" && tencent.TencentType != "ocr"){
		global.GVA_LOG.Error("获取失败!类型不正确!")
		response.FailWithMessage("获取失败!类型不正确!", c)
	}

	cosMap := make(map[string]map[string]string)
	cosMap["tencent_cos"] = map[string]string{"secretId":"AKIDxjHaeldFckcFcjbEbjbEbiaDaiaDaiaD","secretKey":"PZU9ZjeJjuUpukLkzzjEUYOzLSJvJSM0"}
	mapMap := make(map[string]map[string]string)
	mapMap["tencent_map"] = map[string]string{"key":"6Z7BZ-6LP64-I46U2-DD7QB-ILYY7-YOBVF"}
	orcMap := make(map[string]map[string]string)
	orcMap["ocr_map"] = map[string]string{"apiKey":"Ose3OSge4K0TmYv5TLkWffu4","secretKey":"waZh6L8BGubAC2PYknpYvFnDpdNEmdOy"}

	if tencent.TencentType == "cos" {
        response.OkWithDetailed(cosMap, "获取成功", c)
    } else if tencent.TencentType == "map" {
		response.OkWithDetailed(mapMap, "获取成功", c)
	} else if tencent.TencentType == "ocr" {
		response.OkWithDetailed(orcMap, "获取成功", c)
	}
}
