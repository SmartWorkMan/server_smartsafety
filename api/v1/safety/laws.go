package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type LawsApi struct {
}


// CreateLaws 创建Laws
// @Tags Laws
// @Summary 创建Laws
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Laws true "创建Laws"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /laws/createLaws [post]
func (lawsApi *LawsApi) CreateLaws(c *gin.Context) {
	var laws safety.Laws
	_ = c.ShouldBindJSON(&laws)
	if err := lawsService.CreateLaws(laws); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLaws 删除Laws
// @Tags Laws
// @Summary 删除Laws
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Laws true "删除Laws"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /laws/deleteLaws [delete]
func (lawsApi *LawsApi) DeleteLaws(c *gin.Context) {
	var laws safety.Laws
	_ = c.ShouldBindJSON(&laws)
	if err := lawsService.DeleteLaws(laws); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLawsByIds 批量删除Laws
// @Tags Laws
// @Summary 批量删除Laws
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Laws"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /laws/deleteLawsByIds [delete]
func (lawsApi *LawsApi) DeleteLawsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := lawsService.DeleteLawsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLaws 更新Laws
// @Tags Laws
// @Summary 更新Laws
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Laws true "更新Laws"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /laws/updateLawsStatus [put]
func (lawsApi *LawsApi) UpdateLawsStatus(c *gin.Context) {
	var laws safety.Laws
	_ = c.ShouldBindJSON(&laws)
	if laws.LawName == "" {
		global.GVA_LOG.Error("更新失败!请输入正确的法律法规名称!")
		response.FailWithMessage("更新失败!请输入正确的法律法规名称!", c)
	}
	if laws.LawStatus == "" || (laws.LawStatus != "废止" && laws.LawStatus != "现行") {
		global.GVA_LOG.Error("更新失败!请输入正确的状态!")
		response.FailWithMessage("更新失败!请输入正确的状态!", c)
	}
	if err := lawsService.UpdateLawsStatus(laws); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLaws 用id查询Laws
// @Tags Laws
// @Summary 用id查询Laws
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Laws true "用id查询Laws"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /laws/findLaws [get]
func (lawsApi *LawsApi) FindLaws(c *gin.Context) {
	var laws safety.Laws
	_ = c.ShouldBindQuery(&laws)
	if err, relaws := lawsService.GetLaws(laws.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relaws": relaws}, c)
	}
}

// GetLawsList 分页获取Laws列表
// @Tags Laws
// @Summary 分页获取Laws列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.LawsSearch true "分页获取Laws列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /laws/getLawsList [get]
func (lawsApi *LawsApi) GetLawsList(c *gin.Context) {
	var pageInfo safetyReq.LawsSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.Sort == "" {
		pageInfo.Sort = "desc"
	} else if strings.ToLower(pageInfo.Sort) != "desc" &&  strings.ToLower(pageInfo.Sort) != "asc" {
		global.GVA_LOG.Error("获取法律法规失败!请输入正确的排序方式!")
		response.FailWithMessage("获取法律法规失败!请输入正确的排序方式!", c)
		return
	}

	if pageInfo.LawType == "" {
		global.GVA_LOG.Error("获取法律法规失败!请输入正确的法律类型!")
		response.FailWithMessage("获取法律法规失败!请输入正确的法律类型!", c)
		return
	}

	if err, list, total := lawsService.GetLawsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取法律法规失败!", zap.Error(err))
        response.FailWithMessage("获取法律法规失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取法律法规成功", c)
    }
}

// @Router /laws/app/getLawsListForApp [get]
func (lawsApi *LawsApi) GetLawsListForApp(c *gin.Context) {
	var pageInfo safetyReq.LawsSearch
	_ = c.ShouldBindJSON(&pageInfo)

	if err, list, total := lawsService.GetLawsListForApp(pageInfo); err != nil {
		global.GVA_LOG.Error("获取法律法规失败!", zap.Error(err))
		response.FailWithMessage("获取法律法规失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取法律法规成功", c)
	}
}
