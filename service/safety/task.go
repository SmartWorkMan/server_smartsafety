package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type TaskService struct {
}

// CreateTask 创建Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) CreateTask(inputTask safety.Task) (err error) {
	var task safety.Task
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND item_id = ? AND plan_inspection_date = ?", inputTask.FactoryName, inputTask.ItemId, inputTask.PlanInspectionDate).First(&task).Error, gorm.ErrRecordNotFound) {
		return errors.New("巡检任务已创建")
	}
	err = global.GVA_DB.Create(&inputTask).Error
	return err
}

func (taskService *TaskService) CreateTaskHistory(inputTask safety.Task) (err error) {
	if inputTask.TaskStatus == 0 {
		return errors.New("未开始任务不需要记录")
	}

	var inputTaskHistory safety.TaskHistory
	inputTaskHistory.Task = inputTask
	var zeroTime time.Time
	var zeroDelTime gorm.DeletedAt
	inputTaskHistory.CreatedAt = zeroTime
	inputTaskHistory.UpdatedAt = zeroTime
	inputTaskHistory.DeletedAt = zeroDelTime
	curTime := time.Now().Format("2006-01-02 15:04:05")
	inputTaskHistory.ActualInspectionTime = curTime
	inputTaskHistory.TaskId = inputTaskHistory.ID
	inputTaskHistory.ID = 0

	err = global.GVA_DB.Create(&inputTaskHistory).Error
	return err
}

// DeleteTask 删除Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)DeleteTask(task safety.Task) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}

// DeleteTaskByIds 批量删除Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)DeleteTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Task{},"id in ?",ids.Ids).Error
	return err
}

func (taskService *TaskService)DeleteTaskByAreaId(areaId uint) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Task{},"area_id = ?", areaId).Error
	return err
}

func (taskService *TaskService)DeleteTaskByItemId(itemId uint) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Task{},"item_id = ?", itemId).Error
	return err
}

func (taskService *TaskService)DeleteTaskByFactoryName(factoryName string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Task{},"factory_name = ?", factoryName).Error
	return err
}

// UpdateTask 更新Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)ReportTaskResult(task safety.Task) (err error) {
	db := global.GVA_DB.Model(&safety.Task{})
	curTime := time.Now().Format("2006-01-02 15:04:05")
	updateTask := safety.Task{TaskStatus: task.TaskStatus, TaskStatusStr: task.TaskStatusStr, ItemPic: task.ItemPic, ItemDesc: task.ItemDesc, ItemValue: task.ItemValue, FixPic: task.FixPic, FixDesc: task.FixDesc, ActualInspectionTime: curTime}
	err = db.Where("id = ?", task.ID).Updates(updateTask).Error
	if err != nil {
		return err
	}
	_, outTask := taskService.GetTask(task.ID)
	err = taskService.CreateTaskHistory(outTask)
	return err
}

func (taskService *TaskService)AssignTask(task safety.Task) (err error) {
	db := global.GVA_DB.Model(&safety.Task{})
	updateTask := safety.Task{ InspectorUsername: task.InspectorUsername, InspectorName: task.InspectorName, TaskStatus: commval.TaskStatusAssignTask, TaskStatusStr: commval.TaskStatus[commval.TaskStatusAssignTask]}
	err = db.Where("id = ?", task.ID).Updates(updateTask).Error
	if err != nil {
		return err
	}
	_, outTask := taskService.GetTask(task.ID)
	err = taskService.CreateTaskHistory(outTask)
	return err
}

func (taskService *TaskService)ApproveTask(task safety.Task) (err error) {
	db := global.GVA_DB.Model(&safety.Task{})
	updateTask := safety.Task{ TaskStatus: commval.TaskStatusEnd, TaskStatusStr: "审批完成"}
	err = db.Where("id = ?", task.ID).Updates(updateTask).Error
	if err != nil {
		return err
	}
	_, outTask := taskService.GetTask(task.ID)
	err = taskService.CreateTaskHistory(outTask)
	return err
}

// GetTask 根据id获取Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)GetTask(id uint) (err error, task safety.Task) {
	err = global.GVA_DB.Where("id = ?", id).First(&task).Error
	return
}

// GetTaskInfoList 分页获取Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)GetTaskInfoList(info safetyReq.TaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Task{})
    var tasks []safety.Task
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TaskStatus == commval.TaskStatusReportIssue ||
		info.TaskStatus == commval.TaskStatusAssignTask ||
		info.TaskStatus == commval.TaskStatusApproval {
		err = db.Where("factory_name = ? AND task_status = ?", info.FactoryName, info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND task_status = ?", info.FactoryName, info.TaskStatus).Error
	} else {
		err = db.Where("factory_name = ? AND plan_inspection_date = ? AND task_status = ?", info.FactoryName, info.PlanInspectionDate, info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND plan_inspection_date = ? AND task_status = ?", info.FactoryName, info.PlanInspectionDate, info.TaskStatus).Error
	}

	return err, tasks, total
}

func (taskService *TaskService)GetTaskHistory(info safetyReq.ReqTaskHistory) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&safety.TaskHistory{})
	var tasks []safety.TaskHistory

	if info.ItemName != "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND item_name like ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.FactoryName, "%"+info.ItemName+"%", info.TaskStatus, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND item_name like ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.FactoryName, "%"+info.ItemName+"%", info.TaskStatus, timeRange[0], timeRange[1]).Error
	} else if info.ItemName != "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND item_name like ? AND task_status = ?", info.FactoryName, "%"+info.ItemName+"%", info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND item_name like ? AND task_status = ?", info.FactoryName, "%"+info.ItemName+"%", info.TaskStatus).Error
	} else if info.ItemName == "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.FactoryName, info.TaskStatus, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.FactoryName, info.TaskStatus, timeRange[0], timeRange[1]).Error
	} else if info.ItemName == "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND task_status = ?", info.FactoryName, info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "factory_name = ? AND task_status = ?", info.FactoryName, info.TaskStatus).Error
	}
	return err, tasks, total
}

func (taskService *TaskService)GetTaskHistoryByItemForInspector(info safetyReq.ReqTaskHistory) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.TaskHistory{})
	var tasks []safety.TaskHistory
	//zeroTime := time.Now().Format("2006-01-02")
	//zeroTime += " 00:00:00"

	if info.ItemId != 0 && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")

		err = db.Where("item_id = ? AND inspector_username = ? actual_inspection_time >= ? And actual_inspection_time <= ?", info.ItemId, info.InspectorUsername, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks,"item_id = ? AND inspector_username = ? actual_inspection_time >= ? And actual_inspection_time <= ?", info.ItemId, info.InspectorUsername, timeRange[0], timeRange[1]).Error
	} else if info.ItemId != 0 && info.TimeRange == "" {
		err = db.Where("item_id = ? AND inspector_username = ?", info.ItemId, info.InspectorUsername).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "item_id = ? AND inspector_username = ?", info.ItemId, info.InspectorUsername).Error
	} else if info.ItemId == 0 && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")

		err = db.Where("inspector_username = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, timeRange[0], timeRange[1]).Error
	} else if info.ItemId == 0 && info.TimeRange == "" {
		err = db.Where("inspector_username = ?", info.InspectorUsername).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ?", info.InspectorUsername).Error
	}
	return err, tasks, total
}

func (taskService *TaskService)GetTaskHistoryByStatusForInspector(info safetyReq.ReqTaskHistory) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&safety.TaskHistory{})
	var tasks []safety.TaskHistory

	if info.ItemName != "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("inspector_username = ? AND item_name like ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, "%"+info.ItemName+"%", info.TaskStatus, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND item_name like ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, "%"+info.ItemName+"%", info.TaskStatus, timeRange[0], timeRange[1]).Error
	} else if info.ItemName != "" && info.TimeRange == "" {
		err = db.Where("inspector_username = ? AND item_name like ? AND task_status = ?", info.InspectorUsername, "%"+info.ItemName+"%", info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND item_name like ? AND task_status = ?", info.InspectorUsername, "%"+info.ItemName+"%", info.TaskStatus).Error
	} else if info.ItemName == "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("inspector_username = ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, info.TaskStatus, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND task_status = ? And actual_inspection_time >= ? And actual_inspection_time <= ?", info.InspectorUsername, info.TaskStatus, timeRange[0], timeRange[1]).Error
	} else if info.ItemName == "" && info.TimeRange == "" {
		err = db.Where("inspector_username = ? AND task_status = ?", info.InspectorUsername, info.TaskStatus).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND task_status = ?", info.InspectorUsername, info.TaskStatus).Error
	}
	return err, tasks, total
}

func (taskService *TaskService)GetTaskHistoryByStatusStrForAppAdmin(info safetyReq.ReqTaskHistory) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&safety.TaskHistory{})
	var tasks []safety.TaskHistory

	err = db.Where("factory_name = ? AND task_status_str = ?", info.FactoryName, info.TaskStatusStr).Count(&total).Error
	if err!=nil {
		return
	}
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
		Desc:   true,
	}).Offset(offset).Find(&tasks, "factory_name = ? AND task_status_str = ?", info.FactoryName, info.TaskStatusStr).Error

	return err, tasks, total
}

func (taskService *TaskService)GetTaskListByAreaForInspector(info safetyReq.TaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Task{})
	var tasks []safety.Task

	err = db.Where("inspector_username = ? AND area_id = ? AND (((task_status = 0 OR task_status = 4) AND plan_inspection_date = ?) OR (task_status != 0 AND task_status != 4))", info.InspectorUsername, info.AreaId, info.PlanInspectionDate).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
		Desc:   true,
	}).Offset(offset).Find(&tasks, "inspector_username = ? AND area_id = ? AND (((task_status = 0 OR task_status = 4) AND plan_inspection_date = ?) OR (task_status != 0 AND task_status != 4))", info.InspectorUsername, info.AreaId, info.PlanInspectionDate).Error


	return err, tasks, total
}

func (taskService *TaskService)GetTaskListByStatusForInspector(info safetyReq.TaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Task{})
	var tasks []safety.Task

	if info.TaskStatus == commval.TaskStatusNotStart {
		err = db.Where("inspector_username = ? AND task_status = 0 AND plan_inspection_date = ?", info.InspectorUsername, info.PlanInspectionDate).Count(&total).Error
		if err != nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND task_status = 0 AND plan_inspection_date = ?", info.InspectorUsername, info.PlanInspectionDate).Error
	} else {
		err = db.Where("inspector_username = ? AND task_status = ?", info.InspectorUsername, info.TaskStatus).Count(&total).Error
		if err != nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
			Desc:   true,
		}).Offset(offset).Find(&tasks, "inspector_username = ? AND task_status = ?", info.InspectorUsername, info.TaskStatus).Error
	}

	return err, tasks, total
}

func (taskService *TaskService)GetTaskListByInspectorForAreaTree(info safetyReq.TaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Task{})
	var tasks []safety.Task
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? AND inspector_username = ? AND (((task_status = 0 OR task_status = 4) AND plan_inspection_date = ?) OR (task_status != 0 AND task_status != 4))", info.FactoryName, info.InspectorUsername, info.PlanInspectionDate).Count(&total).Error
	if err!=nil {
		return
	}
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "actual_inspection_time"},
		Desc:   true,
	}).Offset(offset).Find(&tasks, "factory_name = ? AND inspector_username = ? AND (((task_status = 0 OR task_status = 4) AND plan_inspection_date = ?) OR (task_status != 0 AND task_status != 4))", info.FactoryName, info.InspectorUsername, info.PlanInspectionDate).Error
	return err, tasks, total
}

func (taskService *TaskService)GetTaskByItem(itemId uint) (err error, list []safety.Task) {
	db := global.GVA_DB.Model(&safety.Task{})
	var tasks []safety.Task
	err = db.Find(&tasks, "item_id = ?", itemId).Error
	return err, tasks
}

func (taskService *TaskService) IsExistActiveTask(inputTask safety.Task) bool {
	var task safety.Task
	if errors.Is(global.GVA_DB.Where("factory_name = ? AND item_id = ? AND (task_status = ? OR task_status = ? OR task_status = ?)", inputTask.FactoryName, inputTask.ItemId, commval.TaskStatusReportIssue, commval.TaskStatusAssignTask, commval.TaskStatusApproval).First(&task).Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
