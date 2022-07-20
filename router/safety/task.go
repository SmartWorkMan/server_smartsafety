package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
}

// InitTaskRouter 初始化 Task 路由信息
func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("task")
	var taskApi = v1.ApiGroupApp.SafetyApiGroup.TaskApi
	{
		//taskRouter.POST("createTask", taskApi.CreateTask)   // 新建Task
		taskRouter.DELETE("deleteTask", taskApi.DeleteTask) // 删除Task
		taskRouter.DELETE("deleteTaskByIds", taskApi.DeleteTaskByIds) // 批量删除Task
		taskRouter.PUT("assignTask", taskApi.AssignTask)    // 维保管理员下派任务
		taskRouter.PUT("approveTask", taskApi.ApproveTask)    // 巡检管理员审批任务
		taskRouter.PUT("rejectTask", taskApi.RejectTask)    // 巡检管理员拒绝审批任务
	}
	{
		taskRouterWithoutRecord.GET("findTask", taskApi.FindTask)        // 根据ID获取Task
		taskRouterWithoutRecord.POST("getNotStartTaskList", taskApi.GetNotStartTaskList)  // 获取检查任务列表
		taskRouterWithoutRecord.POST("getFaultTaskList", taskApi.GetFaultTaskList)  // 获取故障任务列表
		taskRouterWithoutRecord.POST("getAssignTaskList", taskApi.GetAssignTaskList)  // 获取下派任务列表
		taskRouterWithoutRecord.POST("getApprovalTaskList", taskApi.GetApprovalTaskList)  // 获取审批任务列表
		taskRouterWithoutRecord.POST("getTaskHistory", taskApi.GetTaskHistory)  // 巡检管理员获取巡检任务历史记录
		taskRouterWithoutRecord.POST("getTimeOutTaskHistory", taskApi.GetTimeOutTaskHistory)

		taskRouterWithoutRecord.PUT( "app/reportTaskResult", taskApi.ReportTaskResult)    // 巡检员提交巡检结果
		taskRouterWithoutRecord.POST("app/getTaskHistoryByItem", taskApi.GetTaskHistoryByItem)  // 巡检员获取巡检任务历史记录
		taskRouterWithoutRecord.POST("app/getTaskHistoryByStatus", taskApi.GetTaskHistoryByStatus)  // 巡检员获取巡检任务历史记录
		taskRouterWithoutRecord.POST("app/getTaskListByArea", taskApi.GetTaskListByArea)  // 巡检员获取指定区域巡检任务
		taskRouterWithoutRecord.POST("app/getTaskListByStatus", taskApi.GetTaskListByStatus)  // 巡检员获取指定区域巡检任务
		taskRouterWithoutRecord.POST("app/getTaskHistoryByStatusStr", taskApi.GetTaskHistoryByStatusStrForAppAdmin)  // 巡检员获取指定区域巡检任务
		taskRouterWithoutRecord.POST("app/inspectorGetTaskHistoryByStatusStr", taskApi.GetTaskHistoryByStatusStrForInspector)
		taskRouterWithoutRecord.POST("app/getFaultTaskList", taskApi.GetFaultTaskListForAppAdmin)  // 巡检员获取指定区域巡检任务
		taskRouterWithoutRecord.POST("app/getApprovalTaskList", taskApi.GetApprovalTaskListForAppAdmin)  // 巡检员获取指定区域巡检任务
		taskRouterWithoutRecord.POST("app/getAssignTaskList", taskApi.GetAssignTaskListForAppAdmin)  // 巡检员获取指定区域巡检任务

		taskRouterWithoutRecord.POST("app/getInspectorTimeOutTaskCount", taskApi.GetInspectorTimeOutTaskCount)
		taskRouterWithoutRecord.POST("app/getInspectorTodayInspectTaskCount", taskApi.GetInspectorTodayInspectTaskCount)
		taskRouterWithoutRecord.POST("app/getInspectorNotFixedTaskCount", taskApi.GetInspectorNotFixedTaskCount)
		taskRouterWithoutRecord.POST("app/getInspectorTodayNotInspectTaskCount", taskApi.GetInspectorTodayNotInspectTaskCount)

		taskRouterWithoutRecord.POST("temp/createTask", taskApi.TempCreateTask)   // 临时创建巡检任务(此API只为测试使用,上线后删除)
	}
}
