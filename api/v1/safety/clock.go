package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type ClockApi struct {
}


// CreateClock 创建Clock
// @Tags Clock
// @Summary 创建Clock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Clock true "创建Clock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock/app/createClock [post]
func (clockApi *ClockApi) CreateClock(c *gin.Context) {
	var submitClock safetyReq.SubmitClock
	_ = c.ShouldBindJSON(&submitClock)
	if submitClock.InspectorUsername == "" ||
		(submitClock.ClockType != "上班" && submitClock.ClockType != "下班") ||
		submitClock.ClockTime == ""{
		global.GVA_LOG.Error("打卡失败!请检查输入!")
		response.FailWithMessage("打卡失败!请检查输入!", c)
		return
	}
	err, inspector := inspectorService.GetInspectorByUserName(submitClock.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("打卡失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("打卡失败!请输入正确的巡检员用户名!", c)
		return
	}

	var clock safety.Clock
	curDate := time.Now().Format("2006-01-02")
	clock.ClockDate = curDate
	clock.FactoryName = submitClock.FactoryName
	clock.InspectorUsername = submitClock.InspectorUsername
	clock.InspectorName = inspector.Name
	clock.Job = inspector.Job
	clock.Depart = inspector.Depart
	if submitClock.ClockType == "上班" {
		clock.ClockInLocation = submitClock.Location
		clock.ClockInPic = submitClock.ClockPic
		clock.ClockInTime = submitClock.ClockTime
	} else {
		clock.ClockOutLocation = submitClock.Location
		clock.ClockOutPic = submitClock.ClockPic
		clock.ClockOutTime = submitClock.ClockTime
	}

	if err := clockService.CreateClock(clock, submitClock.ClockType); err != nil {
        global.GVA_LOG.Error("打卡失败!", zap.Error(err))
		response.FailWithMessage("打卡失败", c)
	} else {
		response.OkWithMessage("打卡成功", c)
	}
}

// @Router /clock/app/queryClock [post]
func (clockApi *ClockApi) QueryClock(c *gin.Context) {
	var clock safety.Clock
	_ = c.ShouldBindJSON(&clock)
	if clock.InspectorUsername == "" || clock.ClockDate == "" {
		global.GVA_LOG.Error("查询打卡失败!请检查输入!")
		response.FailWithMessage("查询打卡失败!请检查输入!", c)
		return
	}
	if err, reclock := clockService.QueryClock(clock); err != nil {
		global.GVA_LOG.Error("查询打卡失败!", zap.Error(err))
		response.FailWithMessage("查询打卡失败", c)
	} else {
		response.OkWithDetailed(reclock, "查询打卡成功", c)
	}
}

// DeleteClock 删除Clock
// @Tags Clock
// @Summary 删除Clock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Clock true "删除Clock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clock/deleteClock [delete]
func (clockApi *ClockApi) DeleteClock(c *gin.Context) {
	var clock safety.Clock
	_ = c.ShouldBindJSON(&clock)
	if err := clockService.DeleteClock(clock); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClockByIds 批量删除Clock
// @Tags Clock
// @Summary 批量删除Clock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Clock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clock/deleteClockByIds [delete]
func (clockApi *ClockApi) DeleteClockByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := clockService.DeleteClockByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClock 更新Clock
// @Tags Clock
// @Summary 更新Clock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Clock true "更新Clock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clock/updateClock [put]
/*func (clockApi *ClockApi) UpdateClock(c *gin.Context) {
	var clock safety.Clock
	_ = c.ShouldBindJSON(&clock)
	if err := clockService.UpdateClock(clock); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}*/

// FindClock 用id查询Clock
// @Tags Clock
// @Summary 用id查询Clock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Clock true "用id查询Clock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clock/findClock [get]
func (clockApi *ClockApi) FindClock(c *gin.Context) {
	var clock safety.Clock
	_ = c.ShouldBindQuery(&clock)
	if err, reclock := clockService.GetClock(clock.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclock": reclock}, c)
	}
}

// GetClockList 分页获取Clock列表
// @Tags Clock
// @Summary 分页获取Clock列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.ClockSearch true "分页获取Clock列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock/getTodayClockList [post]
func (clockApi *ClockApi) GetTodayClockList(c *gin.Context) {
	var pageInfo safetyReq.ClockSearch
	_ = c.ShouldBindJSON(&pageInfo)
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取今日打卡记录失败!", zap.Error(err))
		response.FailWithMessage("获取今日打卡记录失败", c)
		return
	}
	pageInfo.FactoryName = curUser.FactoryName
	curDate := time.Now().Format("2006-01-02")
	pageInfo.ClockDate = curDate

	if err, list, total := clockService.GetTodayClockInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取今日打卡记录失败!", zap.Error(err))
        response.FailWithMessage("获取今日打卡记录失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取今日打卡记录成功", c)
    }
}

// @Router /clock/getOnDutyNum [get]
func (clockApi *ClockApi) GetOnDutyNum(c *gin.Context) {
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取今日在岗人数失败!", zap.Error(err))
		response.FailWithMessage("获取今日在岗人数失败", c)
		return
	}
	var clock safety.Clock
	clock.FactoryName = curUser.FactoryName
	curDate := time.Now().Format("2006-01-02")
	clock.ClockDate = curDate

	if err, total := clockService.GetOnDutyNum(clock); err != nil {
		global.GVA_LOG.Error("获取今日在岗人数失败!", zap.Error(err))
		response.FailWithMessage("获取今日在岗人数失败", c)
	} else {
		response.OkWithDetailed(total, "获取今日在岗人数成功", c)
	}
}

// @Router /clock/getHistoryClockList [post]
func (clockApi *ClockApi) GetHistoryClockList(c *gin.Context) {
	var pageInfo safetyReq.ClockSearch
	_ = c.ShouldBindJSON(&pageInfo)
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取历史打卡记录失败!", zap.Error(err))
		response.FailWithMessage("获取历史打卡记录失败", c)
		return
	}
	pageInfo.FactoryName = curUser.FactoryName
	if pageInfo.InspectorUsername == "" {
		global.GVA_LOG.Error("获取历史打卡记录失败!巡检员用户名为空!")
		response.FailWithMessage("获取历史打卡记录失败!巡检员用户名为空!", c)
		return
	}

	if err, list, total := clockService.GetHistoryClockList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取历史打卡记录失败!", zap.Error(err))
		response.FailWithMessage("获取历史打卡记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取历史打卡记录成功", c)
	}
}
