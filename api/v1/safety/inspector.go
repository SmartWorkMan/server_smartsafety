package safety

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InspectorApi struct {
}


// CreateInspector 创建Inspector
// @Tags Inspector
// @Summary 创建Inspector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Inspector true "创建Inspector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inspector/createInspector [post]
func (inspectorApi *InspectorApi) CreateInspector(c *gin.Context) {
	selfUserInfo := utils.GetUserInfo(c)
	global.GVA_LOG.Info(fmt.Sprintf("user info:%+v", selfUserInfo))

	var curUser *system.SysUser
	var err error
	err, curUser = userService.FindUserById(int(selfUserInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("创建巡检员失败!", zap.Error(err))
		response.FailWithMessage("创建巡检员失败", c)
		return
	}

	if curUser.FactoryName == "" {
		global.GVA_LOG.Error("创建巡检员失败!当前用户工厂名称为空")
		response.FailWithMessage("创建巡检员失败!当前用户工厂名称为空", c)
		return
	}

	var inspector safety.Inspector
	_ = c.ShouldBindJSON(&inspector)
	inspector.FactoryName = curUser.FactoryName
	if err := inspectorService.CreateInspector(inspector); err != nil {
        global.GVA_LOG.Error("创建巡检员失败!", zap.Error(err))
		response.FailWithMessage("创建巡检员失败", c)
	} else {
		response.OkWithMessage("创建巡检员成功", c)
	}
}

// DeleteInspector 删除Inspector
// @Tags Inspector
// @Summary 删除Inspector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Inspector true "删除Inspector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inspector/deleteInspector [delete]
func (inspectorApi *InspectorApi) DeleteInspector(c *gin.Context) {
	var inspector safety.Inspector
	_ = c.ShouldBindJSON(&inspector)
	if err := inspectorService.DeleteInspector(inspector); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInspectorByIds 批量删除Inspector
// @Tags Inspector
// @Summary 批量删除Inspector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Inspector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inspector/deleteInspectorByIds [delete]
func (inspectorApi *InspectorApi) DeleteInspectorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := inspectorService.DeleteInspectorByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInspector 更新Inspector
// @Tags Inspector
// @Summary 更新Inspector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Inspector true "更新Inspector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inspector/updateInspector [put]
func (inspectorApi *InspectorApi) UpdateInspector(c *gin.Context) {
	selfUserInfo := utils.GetUserInfo(c)
	var curUser *system.SysUser
	var err error
	err, curUser = userService.FindUserById(int(selfUserInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("创建巡检员失败!", zap.Error(err))
		response.FailWithMessage("创建巡检员失败", c)
		return
	}

	if curUser.FactoryName == "" {
		global.GVA_LOG.Error("创建巡检员失败!当前用户工厂名称为空")
		response.FailWithMessage("创建巡检员失败!当前用户工厂名称为空", c)
		return
	}

	var inspector safety.Inspector
	_ = c.ShouldBindJSON(&inspector)
	inspector.FactoryName = curUser.FactoryName

	if err := inspectorService.UpdateInspector(inspector); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInspector 用id查询Inspector
// @Tags Inspector
// @Summary 用id查询Inspector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Inspector true "用id查询Inspector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inspector/findInspector [get]
func (inspectorApi *InspectorApi) FindInspector(c *gin.Context) {
	var inspector safety.Inspector
	_ = c.ShouldBindQuery(&inspector)
	if err, reinspector := inspectorService.GetInspector(inspector.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinspector": reinspector}, c)
	}
}

// GetInspectorList 分页获取Inspector列表
// @Tags Inspector
// @Summary 分页获取Inspector列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.InspectorSearch true "分页获取Inspector列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inspector/getInspectorList [get]
func (inspectorApi *InspectorApi) GetInspectorList(c *gin.Context) {
	selfUserInfo := utils.GetUserInfo(c)

	var curUser *system.SysUser
	var err error
	err, curUser = userService.FindUserById(int(selfUserInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("获取巡检员列表失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员列表失败", c)
		return
	}

	if curUser.FactoryName == "" {
		global.GVA_LOG.Error("获取巡检员列表失败!当前用户工厂名称为空")
		response.FailWithMessage("获取巡检员列表失败!当前用户工厂名称为空", c)
		return
	}

	var pageInfo safetyReq.InspectorSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.FactoryName = curUser.FactoryName
	if err, list, total := inspectorService.GetInspectorInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取巡检员列表失败!", zap.Error(err))
        response.FailWithMessage("获取巡检员列表失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取巡检员列表成功", c)
    }
}

// @Router /inspector/login [post]
func (inspectorApi *InspectorApi) Login(c *gin.Context) {
	var inspector safety.Inspector
	_ = c.ShouldBindJSON(&inspector)

	if err := inspectorService.Login(inspector); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		response.OkWithMessage("登录成功", c)
	}
}