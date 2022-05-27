// 自动生成模板Task
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Task 结构体
// 如果含有time.Time 请自行import time包
type Task struct {
      global.GVA_MODEL
      ActualInspectionTime  string `json:"actualInspectionTime" form:"actualInspectionTime" gorm:"column:actual_inspection_time;comment:实际巡检时间;size:191;"`
      AreaId  uint `json:"areaId" form:"areaId" gorm:"column:area_id;comment:巡检区域ID;"`
      AreaName  string `json:"areaName" form:"areaName" gorm:"column:area_name;comment:巡检区域名称;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      FixDesc  string `json:"fixDesc" form:"fixDesc" gorm:"column:fix_desc;comment:整改情况描述;size:191;"`
      FixPic  string `json:"fixPic" form:"fixPic" gorm:"column:fix_pic;comment:整改照片视频;size:191;"`
      InspectorName  string `json:"inspectorName" form:"inspectorName" gorm:"column:inspector_name;comment:巡检员姓名;size:191;"`
      InspectorUsername  string `json:"inspectorUsername" form:"inspectorUsername" gorm:"column:inspector_username;comment:巡检员用户名;size:191;"`
      ItemDesc  string `json:"itemDesc" form:"itemDesc" gorm:"column:item_desc;comment:现场情况描述;size:191;"`
      ItemName  string `json:"itemName" form:"itemName" gorm:"column:item_name;comment:巡检事项名称;size:191;"`
      ItemPic  string `json:"itemPic" form:"itemPic" gorm:"column:item_pic;comment:现场照片视频;size:191;"`
      ItemSn  string `json:"itemSn" form:"itemSn" gorm:"column:item_sn;comment:巡检事项编码;size:191;"`
      ItemValue  string `json:"itemValue" form:"itemValue" gorm:"column:item_value;comment:现场设备显示值;size:191;"`
      Period  string `json:"period" form:"period" gorm:"column:period;comment:巡检周期,重复日,重复周,重复月;size:191;"`
      PlanInspectionDate  string `json:"planInspectionDate" form:"planInspectionDate" gorm:"column:plan_inspection_date;comment:计划巡检日期;size:191;"`
      Standard  string `json:"standard" form:"standard" gorm:"column:standard;comment:检查标准;size:191;"`
      TaskStatus  int `json:"taskStatus" form:"taskStatus" gorm:"column:task_status;comment:任务状态,0:未开始 1:巡检员上报隐患 2:管理员已下派任务 3:巡检员处理下派任务后上报审批 4:任务完成;"`
      ItemId  uint `json:"itemId" form:"itemId" gorm:"column:item_id;comment:巡检事项ID;"`
      TaskStatusStr  string `json:"taskStatusStr" form:"taskStatusStr" gorm:"column:task_status_str;comment:任务状态,0:未开始 1:巡检员上报隐患 2:管理员已下派任务 3:巡检员处理下派任务后上报审批 4:任务完成;size:191;"`
}


// TableName Task 表名
func (Task) TableName() string {
  return "safety_task"
}

