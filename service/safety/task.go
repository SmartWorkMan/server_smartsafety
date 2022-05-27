package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
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

// UpdateTask 更新Task记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService)UpdateTask(task safety.Task) (err error) {
	err = global.GVA_DB.Save(&task).Error
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
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&tasks).Error
	return err, tasks, total
}
