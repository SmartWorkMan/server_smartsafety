package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct{
    safety.Task
    request.PageInfo
}

type ReqTaskHistory struct{
	request.PageInfo
	ItemId  uint `json:"itemId" form:"itemId" gorm:"column:item_id;comment:巡检事项ID;"`
	InspectorUsername  string `json:"inspectorUsername" form:"inspectorUsername" gorm:"column:inspector_username;comment:巡检员用户名;size:191;"`
	FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
	TimeRange  string `json:"timeRange"` //时间范围,格式"2022-05-26 15:39:00~2022-05-27 00:00:00"
	ItemName  string `json:"itemName" form:"itemName" gorm:"column:item_name;comment:巡检事项名称;size:191;"`
	TaskStatus  int `json:"taskStatus" form:"taskStatus" gorm:"column:task_status;comment:任务状态,0:未开始 1:巡检员上报隐患 2:管理员已下派任务 3:巡检员处理下派任务后上报审批 4:任务完成;"`
	TaskStatusStr  string `json:"taskStatusStr" form:"taskStatusStr" gorm:"column:task_status_str;comment:任务状态,0:未开始 1:巡检员上报隐患 2:管理员已下派任务 3:巡检员处理下派任务后上报审批 4:任务完成;size:191;"`
}

type TaskReport struct{
	safety.Task
	ItemPicList  []string `json:"itemPicList"`
	FixPicList  []string `json:"fixPicList"`
}

type TopItem struct {
	ItemName  string `json:"itemName" form:"itemName" gorm:"column:item_name;comment:巡检事项名称;size:191;"`
	Number int `json:"num" gorm:"column:num;"`
}

type TaskApprove struct{
	safety.Task
	PhoneNumberList  []string `json:"phoneNumberList"`
}
