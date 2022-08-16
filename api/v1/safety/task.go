package safety

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"bytes"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type TaskApi struct {
}


// CreateTask 创建Task
// @Tags Task
// @Summary 创建Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Task true "创建Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTask [post]
func (taskApi *TaskApi) CreateTask(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.CreateTask(task); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Router /task/temp/createTask [post]
func (taskApi *TaskApi) TempCreateTask(c *gin.Context) {
	GenerateTask(commval.ItemPeriodDay)
	GenerateTask(commval.ItemPeriodWeek)
	GenerateTask(commval.ItemPeriodMonth)
	GenerateTask(commval.ItemPeriodQuarter)
	GenerateTask(commval.ItemPeriodSemester)
	response.OkWithMessage("创建巡检任务成功(此API只为测试使用,上线后删除)", c)
}

// DeleteTask 删除Task
// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Task true "删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
func (taskApi *TaskApi) DeleteTask(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if err := taskService.DeleteTask(task); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskByIds 批量删除Task
// @Tags Task
// @Summary 批量删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTaskByIds [delete]
func (taskApi *TaskApi) DeleteTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := taskService.DeleteTaskByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func taskReport2Task(taskReport safetyReq.TaskReport) safety.Task {
	var task safety.Task
	task = taskReport.Task
	if len(taskReport.ItemPicList) == 0 {
		task.ItemPic = ""
	} else {
		task.ItemPic = ""
		for i := 0; i < len(taskReport.ItemPicList); i++ {
			if i != len(taskReport.ItemPicList) - 1 {
				task.ItemPic += taskReport.ItemPicList[i] + ","
			} else {
				task.ItemPic += taskReport.ItemPicList[i]
			}
		}
	}

	if len(taskReport.FixPicList) == 0 {
		task.FixPic = ""
	} else {
		task.FixPic = ""
		for i := 0; i < len(taskReport.FixPicList); i++ {
			if i != len(taskReport.FixPicList) - 1 {
				task.FixPic += taskReport.FixPicList[i] + ","
			} else {
				task.FixPic += taskReport.FixPicList[i]
			}
		}
	}
	return task
}

func taskList2TaskReportList (taskList []safety.Task) []safetyReq.TaskReport {
	var taskReportList []safetyReq.TaskReport
	for _, task := range taskList {
		var taskReport safetyReq.TaskReport
		taskReport.Task = task
		taskReport.ItemPic = ""
		taskReport.FixPic = ""
		itemPicList := strings.Split(task.ItemPic, ",")
		fixPicList := strings.Split(task.FixPic, ",")
		taskReport.ItemPicList = itemPicList
		taskReport.FixPicList = fixPicList
		taskReportList = append(taskReportList, taskReport)
	}
	return taskReportList
}

func taskHistoryList2TaskReportList (taskHistoryList []safety.TaskHistory) []safetyReq.TaskReport {
	var taskReportList []safetyReq.TaskReport
	for _, task := range taskHistoryList {
		var taskReport safetyReq.TaskReport
		taskReport.Task = task.Task
		taskReport.ItemPic = ""
		taskReport.FixPic = ""
		itemPicList := strings.Split(task.ItemPic, ",")
		fixPicList := strings.Split(task.FixPic, ",")
		taskReport.ItemPicList = itemPicList
		taskReport.FixPicList = fixPicList
		taskReportList = append(taskReportList, taskReport)
	}
	return taskReportList
}

// UpdateTask 更新Task
// @Tags Task
// @Summary 更新Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Task true "更新Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/app/reportTaskResult [put]
func (taskApi *TaskApi) ReportTaskResult(c *gin.Context) {
	var taskReport safetyReq.TaskReport
	_ = c.ShouldBindJSON(&taskReport)
	if taskReport.ID == 0 {
		global.GVA_LOG.Error("巡检员提交巡检任务失败!请输入正确任务ID")
		response.FailWithMessage("巡检员提交巡检任务失败!请输入正确任务ID", c)
		return
	}
	if taskReport.TaskStatusStr == "" || taskReport.TaskStatus == commval.TaskStatusNotStart {
		global.GVA_LOG.Error("巡检员提交巡检任务失败!请输入正确巡检任务状态!")
		response.FailWithMessage("巡检员提交巡检任务失败!请输入正确巡检任务状态!", c)
		return
	}

	_, taskInfo := taskService.GetTask(taskReport.ID)
	if !strings.Contains(taskInfo.InspectorUsername, taskReport.InspectorUsername) {
		global.GVA_LOG.Error("巡检员提交巡检任务失败!此任务已被其他巡检员处理!")
		response.FailWithMessage("巡检员提交巡检任务失败!此任务已被其他巡检员处理!", c)
		return
	}

	err, inspector := inspectorService.GetInspectorByUserName(taskReport.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("巡检员提交巡检任务失败!请输入正确巡检任务状态!")
		response.FailWithMessage("巡检员提交巡检任务失败!请输入正确巡检任务状态!", c)
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("inspector:%+v", inspector))
	taskReport.InspectorName = inspector.Name

	global.GVA_LOG.Info(fmt.Sprintf("taskReport:%+v", taskReport))
	global.GVA_LOG.Info(fmt.Sprintf("task:%+v", taskReport2Task(taskReport)))

	if err := taskService.ReportTaskResult(taskReport2Task(taskReport)); err != nil {
        global.GVA_LOG.Error("巡检员提交巡检任务失败!", zap.Error(err))
		response.FailWithMessage("巡检员提交巡检任务失败", c)
        return
	} else {
		if taskReport.TaskStatus == commval.TaskStatusReportIssue ||
			taskReport.TaskStatus == commval.TaskStatusApproval ||
			taskReport.TaskStatus == commval.TaskStatusFireAlarm {
			_, taskInfo := taskService.GetTask(taskReport.ID)
			//发短信给维保管理员
			sms := make(map[string]interface{})
			if taskReport.TaskStatus == commval.TaskStatusReportIssue {
				sms["reason"] = "上报维修"
			} else if taskReport.TaskStatus == commval.TaskStatusApproval {
				sms["reason"] = "工单审批"
			} else if taskReport.TaskStatus == commval.TaskStatusFireAlarm {
				sms["reason"] = "火警处置"
			}
			sms["areaInfo"] = taskInfo.AreaName
			sms["deviceInfo"] = taskInfo.ItemName + " " + taskInfo.ItemSn
			sms["people"] = taskInfo.InspectorName
			err, users := userService.GetMaintainUsers(taskInfo.FactoryName, commval.MaintainUserAuthorityId)
			if err != nil {
				global.GVA_LOG.Error("发送短信给维保管理员失败!", zap.Error(err))
			} else {
				var phoneList []string
				for _, user := range users {
					phoneList = append(phoneList, user.Phone)
				}
				sms["phonelist"] = phoneList
				err = sendSMS(sms)
				if err != nil {
					global.GVA_LOG.Error("发送短信给维保管理员失败!", zap.Error(err))
				}
			}

			if taskReport.TaskStatus == commval.TaskStatusFireAlarm {
				//发短信给所有当班巡检员
				var phoneList []string
				var clock safety.Clock
				clock.FactoryName = taskInfo.FactoryName
				err, clockList := clockService.GetOnDutyInspectors(clock)
				if err != nil {
					global.GVA_LOG.Error("发送短信给巡检员失败!", zap.Error(err))
					response.FailWithMessage("发送短信给巡检员失败!", c)
					return
				}
				for _, clockInfo := range clockList {
					err, inspector := inspectorService.GetInspectorByUserName(clockInfo.InspectorUsername)
					if err != nil {
						global.GVA_LOG.Error("发送短信给巡检员失败!", zap.Error(err))
						continue
					}
					phoneList = append(phoneList, inspector.PhoneNumber)
				}
				if len(phoneList) > 0 {
					sms["phonelist"] = phoneList
					err = sendSMS(sms)
					if err != nil {
						global.GVA_LOG.Error("发送短信给巡检员失败!", zap.Error(err))
						response.FailWithMessage("发送短信给巡检员失败!", c)
						return
					}
				}
			}
		}

		response.OkWithMessage("巡检员提交巡检任务成功", c)
	}
}

// @Router /task/assignTask [put]
func (taskApi *TaskApi) AssignTask(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.ID == 0 || task.InspectorUsername == "" || task.InspectorName == "" {
		global.GVA_LOG.Error("下派任务失败!请检查请求信息!")
		response.FailWithMessage("下派任务失败!请检查请求信息!", c)
		return
	}
	if err := taskService.AssignTask(task); err != nil {
		global.GVA_LOG.Error("下派任务失败!", zap.Error(err))
		response.FailWithMessage("下派任务失败", c)
	} else {
		//发短信
		_, taskInfo := taskService.GetTask(task.ID)
		err, users := userService.GetMaintainUsers(taskInfo.FactoryName, commval.MaintainUserAuthorityId)
		if err != nil {
			global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
			response.FailWithMessage("发送短信失败!", c)
			return
		}
		_, inspector := inspectorService.GetInspectorByUserName(task.InspectorUsername)
		sms := make(map[string]interface{})
		sms["reason"] = "任务下派"
		sms["areaInfo"] = taskInfo.AreaName
		sms["deviceInfo"] = taskInfo.ItemName + " " + taskInfo.ItemSn
		sms["people"] = users[0].NickName
		var phoneList []string
		phoneList = append(phoneList, inspector.PhoneNumber)
		sms["phonelist"] = phoneList
		err = sendSMS(sms)
		if err != nil {
			global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
			response.FailWithMessage("发送短信失败!", c)
			return
		}
		response.OkWithMessage("下派任务成功", c)
	}
}

// @Router /task/approveTask [put]
func (taskApi *TaskApi) ApproveTask(c *gin.Context) {
	var task safetyReq.TaskApprove
	_ = c.ShouldBindJSON(&task)
	if task.ID == 0 {
		global.GVA_LOG.Error("审批任务失败!请检查请求信息!")
		response.FailWithMessage("审批任务失败!请检查请求信息!", c)
		return
	}
	if err := taskService.ApproveTask(task.Task); err != nil {
		global.GVA_LOG.Error("审批任务失败!", zap.Error(err))
		response.FailWithMessage("审批任务失败", c)
	} else {
		//火警处置确认发短信
		if len(task.PhoneNumberList) > 0 {
			_, taskInfo := taskService.GetTask(task.ID)
			err, users := userService.GetMaintainUsers(taskInfo.FactoryName, commval.MaintainUserAuthorityId)
			if err != nil {
				global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
				response.FailWithMessage("发送短信失败!", c)
				return
			}
			sms := make(map[string]interface{})
			sms["reason"] = "火警处置确认"
			sms["areaInfo"] = taskInfo.AreaName
			sms["deviceInfo"] = taskInfo.ItemName + " " + taskInfo.ItemSn
			sms["people"] = users[0].NickName
			sms["phonelist"] = task.PhoneNumberList
			err = sendSMS(sms)
			if err != nil {
				global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
				response.FailWithMessage("发送短信失败!", c)
				return
			}
		}
		response.OkWithMessage("审批任务成功", c)
	}
}

// @Router /task/rejectTask [put]
func (taskApi *TaskApi) RejectTask(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.ID == 0 || task.AdminComment == ""{
		global.GVA_LOG.Error("拒绝审批任务失败!请检查请求信息!")
		response.FailWithMessage("拒绝审批任务失败!请检查请求信息!", c)
		return
	}

	if err := taskService.RejectTask(task); err != nil {
		global.GVA_LOG.Error("拒绝审批任务失败!", zap.Error(err))
		response.FailWithMessage("拒绝审批任务失败", c)
	} else {
		//发送短信
		_, taskInfo := taskService.GetTask(task.ID)
		err, users := userService.GetMaintainUsers(taskInfo.FactoryName, commval.MaintainUserAuthorityId)
		if err != nil {
			global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
			response.FailWithMessage("发送短信失败!", c)
			return
		}
		_, inspector := inspectorService.GetInspectorByUserName(taskInfo.InspectorUsername)
		sms := make(map[string]interface{})
		sms["reason"] = "任务驳回"
		sms["areaInfo"] = taskInfo.AreaName
		sms["deviceInfo"] = taskInfo.ItemName + " " + taskInfo.ItemSn
		sms["people"] = users[0].NickName
		var phoneList []string
		phoneList = append(phoneList, inspector.PhoneNumber)
		sms["phonelist"] = phoneList
		err = sendSMS(sms)
		if err != nil {
			global.GVA_LOG.Error("发送短信失败!", zap.Error(err))
			response.FailWithMessage("发送短信失败!", c)
			return
		}
		response.OkWithMessage("拒绝审批任务成功", c)
	}
}

// FindTask 用id查询Task
// @Tags Task
// @Summary 用id查询Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Task true "用id查询Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
func (taskApi *TaskApi) FindTask(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindQuery(&task)
	if err, retask := taskService.GetTask(task.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retask": retask}, c)
	}
}

// GetTaskList 分页获取Task列表
// @Tags Task
// @Summary 分页获取Task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.TaskSearch true "分页获取Task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getNotStartTaskList [post]
func (taskApi *TaskApi) GetNotStartTaskList(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusNotStart)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/getFaultTaskList [post]
func (taskApi *TaskApi) GetFaultTaskList(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusReportIssue)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/getAssignTaskList [post]
func (taskApi *TaskApi) GetAssignTaskList(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusAssignTask)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/getApprovalTaskList [post]
func (taskApi *TaskApi) GetApprovalTaskList(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusApproval)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/getFireAlarmTaskList [post]
func (taskApi *TaskApi) GetFireAlarmTaskList(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusFireAlarm)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

func (taskApi *TaskApi) getTaskList(c *gin.Context, taskStatus int)(error, interface{}, int64, safetyReq.TaskSearch) {
	var pageInfo safetyReq.TaskSearch
	_ = c.ShouldBindJSON(&pageInfo)
	appFlag := c.Request.Header.Get("x-app-flag")
	if appFlag != "" {
		if pageInfo.FactoryName == "" {
			return errors.New("工厂名称不能为空"), nil, 0, pageInfo
		}
	} else {
		err, curUser := GetCurUser(c)
		if err != nil {
			return err, nil, 0, pageInfo
		}
		pageInfo.FactoryName = curUser.FactoryName
	}

	curDate := time.Now().Format("2006-01-02")
	pageInfo.TaskStatus = taskStatus
	pageInfo.PlanInspectionDate = curDate
	err, list, total := taskService.GetTaskInfoList(pageInfo)
	newList := taskList2TaskReportList(list.([]safety.Task))
	return err, newList, total, pageInfo
}

// @Router /task/getTaskHistory [post]
func (taskApi *TaskApi) GetTaskHistory(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败!", c)
		return
	}

	_ = c.ShouldBindJSON(&pageInfo)
	pageInfo.FactoryName = curUser.FactoryName
	if pageInfo.TaskStatusStr == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的记录类型!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的记录类型!", c)
		return
	}

	if err, list, total := taskService.GetTaskHistory(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}

// @Router /task/getTimeOutTaskHistory [post]
func (taskApi *TaskApi) GetTimeOutTaskHistory(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	err, curUser := GetCurUser(c)
	if err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败!", c)
		return
	}

	_ = c.ShouldBindJSON(&pageInfo)
	pageInfo.FactoryName = curUser.FactoryName

	if err, list, total := taskService.GetTimeOutTaskHistory(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}


// @Router /task/app/getTaskHistoryByItem [post]
func (taskApi *TaskApi) GetTaskHistoryByItem(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.InspectorUsername == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的巡检员用户名!", c)
		return
	}

	err, inspector := inspectorService.GetInspectorByUserName(pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的巡检员用户名!", c)
		return
	}
	pageInfo.FactoryName = inspector.FactoryName

	if err, list, total := taskService.GetTaskHistoryByItemForInspector(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}

// @Router /task/app/getTaskHistoryByStatus [post]
func (taskApi *TaskApi) GetTaskHistoryByStatus(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.InspectorUsername == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的巡检员用户名!", c)
		return
	}
	if pageInfo.TaskStatus == commval.TaskStatusNotStart {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的记录类型!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的记录类型!", c)
		return
	}

	err, inspector := inspectorService.GetInspectorByUserName(pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检历史记录失败!请输入正确的巡检员用户名!", c)
		return
	}
	pageInfo.FactoryName = inspector.FactoryName

	if err, list, total := taskService.GetTaskHistoryByStatusForInspector(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}

// @Router /task/app/getTaskHistoryByStatusStr [post]
func (taskApi *TaskApi) GetTaskHistoryByStatusStrForAppAdmin(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.TaskStatusStr == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!任务状态不能为空!")
		response.FailWithMessage("获取巡检历史记录失败!任务状态不能为空!", c)
		return
	}
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!工厂名称不能为空!")
		response.FailWithMessage("获取巡检历史记录失败!工厂名称不能为空!", c)
		return
	}

	if err, list, total := taskService.GetTaskHistoryByStatusStrForAppAdmin(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}

// @Router /task/app/inspectorGetTaskHistoryByStatusStr [post]
func (taskApi *TaskApi) GetTaskHistoryByStatusStrForInspector(c *gin.Context) {
	var pageInfo safetyReq.ReqTaskHistory
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.TaskStatusStr == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!任务状态不能为空!")
		response.FailWithMessage("获取巡检历史记录失败!任务状态不能为空!", c)
		return
	}
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!工厂名称不能为空!")
		response.FailWithMessage("获取巡检历史记录失败!工厂名称不能为空!", c)
		return
	}
	if pageInfo.InspectorUsername == "" {
		global.GVA_LOG.Error("获取巡检历史记录失败!巡检员用户名不能为空!")
		response.FailWithMessage("获取巡检历史记录失败!巡检员用户名不能为空!", c)
		return
	}

	if err, list, total := taskService.GetTaskHistoryByStatusStrForInspector(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检历史记录失败!", zap.Error(err))
		response.FailWithMessage("获取巡检历史记录失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskHistoryList2TaskReportList(list.([]safety.TaskHistory)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检历史记录成功", c)
	}
}

// @Router /task/app/getTaskListByArea [post]
func (taskApi *TaskApi) GetTaskListByArea(c *gin.Context) {
	var pageInfo safetyReq.TaskSearch
	_ = c.ShouldBindBodyWith(&pageInfo, binding.JSON)
	if pageInfo.InspectorUsername == ""{
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请检查输入!")
		response.FailWithMessage("获取巡检员巡检任务失败!请检查输入!", c)
		return
	}
	if pageInfo.AreaId == 0 {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请输入正确区域ID!")
		response.FailWithMessage("获取巡检员巡检任务失败!请输入正确区域ID!", c)
		return
	}

	err, inspector := inspectorService.GetInspectorByUserName(pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检员巡检任务失败!请输入正确的巡检员用户名!", c)
		return
	}
	pageInfo.FactoryName = inspector.FactoryName
	curDate := time.Now().Format("2006-01-02")
	pageInfo.PlanInspectionDate = curDate

	if err, list, total := taskService.GetTaskListByAreaForInspector(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员巡检任务失败", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskList2TaskReportList(list.([]safety.Task)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检员巡检任务成功", c)
	}
}

// @Router /task/app/getTaskListByStatus [post]
// 小程序任务中心
func (taskApi *TaskApi) GetTaskListByStatus(c *gin.Context) {
	var pageInfo safetyReq.TaskSearch
	_ = c.ShouldBindBodyWith(&pageInfo, binding.JSON)
	if pageInfo.InspectorUsername == ""{
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请检查输入!")
		response.FailWithMessage("获取巡检员巡检任务失败!请检查输入!", c)
		return
	}

	bodyMap := make(map[string]interface{})
	_ = c.ShouldBindBodyWith(&bodyMap, binding.JSON)
	_, ok := bodyMap["taskStatus"]
	if !ok {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请检查正确任务状态!")
		response.FailWithMessage("获取巡检员巡检任务失败!请检查正确任务状态!", c)
		return
	}

	if pageInfo.TaskStatus == commval.TaskStatusEnd {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请检查正确任务状态!")
		response.FailWithMessage("获取巡检员巡检任务失败!请检查正确任务状态!", c)
		return
	}

	err, inspector := inspectorService.GetInspectorByUserName(pageInfo.InspectorUsername)
	if err != nil {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!请输入正确的巡检员用户名!")
		response.FailWithMessage("获取巡检员巡检任务失败!请输入正确的巡检员用户名!", c)
		return
	}
	pageInfo.FactoryName = inspector.FactoryName
	curDate := time.Now().Format("2006-01-02")
	pageInfo.PlanInspectionDate = curDate

	if err, list, total := taskService.GetTaskListByStatusForInspector(pageInfo); err != nil {
		global.GVA_LOG.Error("获取巡检员巡检任务失败!", zap.Error(err))
		response.FailWithMessage("获取巡检员巡检任务失败", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     taskList2TaskReportList(list.([]safety.Task)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取巡检员巡检任务成功", c)
	}
}

// @Router /task/app/getFaultTaskList [post]
// factoryName in body
func (taskApi *TaskApi) GetFaultTaskListForAppAdmin(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusReportIssue)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/app/getApprovalTaskList [post]
// factoryName in body
func (taskApi *TaskApi) GetApprovalTaskListForAppAdmin(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusApproval)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/app/getAssignTaskList [post]
// factoryName in body
func (taskApi *TaskApi) GetAssignTaskListForAppAdmin(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusAssignTask)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /task/app/getFireAlarmTaskListForAppAdmin [post]
// factoryName in body
func (taskApi *TaskApi) GetFireAlarmTaskListForAppAdmin(c *gin.Context) {
	err, list, total, pageInfo := taskApi.getTaskList(c, commval.TaskStatusFireAlarm)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取任务列表成功", c)
	}
}

// @Router /screen/getNormalTaskCount [post]
func (taskApi *TaskApi) GetNormalTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, total := taskService.GetNormalTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(total, "获取成功", c)
	}
}

// @Router /screen/getPendingTaskCount [post]
func (taskApi *TaskApi) GetPendingTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, total := taskService.GetPendingTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(total, "获取成功", c)
	}
}

// @Router /screen/getFixedTaskCount [post]
func (taskApi *TaskApi) GetFixedTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, total := taskService.GetFixedTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(total, "获取成功", c)
	}
}

// @Router /screen/getNotFixedTaskCount [post]
func (taskApi *TaskApi) GetNotFixedTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, total := taskService.GetNotFixedTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(total, "获取成功", c)
	}
}

// @Router /screen/getTopFailureItems [post]
func (taskApi *TaskApi) GetTopFailureItems(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, list := taskService.GetTopFailureItems(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
		}, "获取成功", c)
	}
}

// @Router /screen/getFixedStatistics [post]
func (taskApi *TaskApi) GetFixedStatistics(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, fixedTotal, notFixedTotal := taskService.GetFixedStatistics(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]int64{"fixedTotal":fixedTotal, "notFixedTotal":notFixedTotal}, "获取成功", c)
	}
}

// @Router /screen/getWeeklyHealthIndex [post]
func (taskApi *TaskApi) GetWeeklyHealthIndex(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, index := taskService.GetWeeklyHealthIndex(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(index, "获取成功", c)
	}
}

// @Router /screen/getWeeklyFixedCurve [post]
func (taskApi *TaskApi) GetWeeklyFixedCurve(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)

	if task.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, curve := taskService.GetWeeklyFixedCurve(task); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(curve, "获取成功", c)
	}
}

// @Router /screen/getStartInspectInfo [post]
func (taskApi *TaskApi) GetStartInspectInfo(c *gin.Context) {
	var pageInfo safetyReq.TaskSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!工厂名称不能为空!")
		response.FailWithMessage("获取失败!工厂名称不能为空!", c)
		return
	}

	if err, list, total := taskService.GetStartInspectInfo(pageInfo); err != nil {
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


// @Router /task/app/getInspectorTimeOutTaskCount [post]
func (taskApi *TaskApi) GetInspectorTimeOutTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.FactoryName == "" || task.InspectorUsername == "" {
		global.GVA_LOG.Error("获取超时任务数量失败!请检查输入!")
		response.FailWithMessage("获取超时任务数量失败!请检查输入!", c)
		return
	}

	if err, total := taskService.GetInspectorTimeOutTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取超时任务数量失败!", zap.Error(err))
		response.FailWithMessage("获取超时任务数量失败", c)
	} else {
		response.OkWithDetailed(total, "获取超时任务数量成功", c)
	}
}

// @Router /task/app/getInspectorTodayInspectTaskCount [post]
func (taskApi *TaskApi) GetInspectorTodayInspectTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.FactoryName == "" || task.InspectorUsername == "" {
		global.GVA_LOG.Error("获取已巡检任务数量失败!请检查输入!")
		response.FailWithMessage("获取已巡检任务数量失败!请检查输入!", c)
		return
	}

	if err, total := taskService.GetInspectorTodayInspectTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取已巡检任务数量失败!", zap.Error(err))
		response.FailWithMessage("获取已巡检任务数量失败", c)
	} else {
		response.OkWithDetailed(total, "获取已巡检任务数量成功", c)
	}
}

// @Router /task/app/getInspectorNotFixedTaskCount [post]
func (taskApi *TaskApi) GetInspectorNotFixedTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.FactoryName == "" || task.InspectorUsername == "" {
		global.GVA_LOG.Error("获取完成任务数量失败!请检查输入!")
		response.FailWithMessage("获取完成任务数量失败!请检查输入!", c)
		return
	}

	if err, total := taskService.GetInspectorNotFixedTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取完成任务数量失败!", zap.Error(err))
		response.FailWithMessage("获取完成任务数量失败", c)
	} else {
		response.OkWithDetailed(total, "获取完成任务数量成功", c)
	}
}

// @Router /task/app/getInspectorTodayNotInspectTaskCount [post]
func (taskApi *TaskApi) GetInspectorTodayNotInspectTaskCount(c *gin.Context) {
	var task safety.Task
	_ = c.ShouldBindJSON(&task)
	if task.FactoryName == "" || task.InspectorUsername == "" {
		global.GVA_LOG.Error("获取未巡检任务数量失败!请检查输入!")
		response.FailWithMessage("获取未巡检任务数量失败!请检查输入!", c)
		return
	}

	if err, total := taskService.GetInspectorTodayNotInspectTaskCount(task); err != nil {
		global.GVA_LOG.Error("获取未巡检任务数量失败!", zap.Error(err))
		response.FailWithMessage("获取未巡检任务数量失败", c)
	} else {
		response.OkWithDetailed(total, "获取未巡检任务数量成功", c)
	}
}

func sendSMS(sms map[string]interface{}) error {
	url := "http://xulaogemeishi.top:9244/api/smsSend/batchSend"
	//sms["phonelist"] = []string{"13770947479"}
	data, err := json.Marshal(sms)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("send sms failed, sms: %v, error: %s", sms, err.Error()))
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("send sms failed, sms: %s, error: %s", data, err.Error()))
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("send sms faild, return %d", res.StatusCode))
	} else {
		return nil
	}
}