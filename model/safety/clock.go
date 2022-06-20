// 自动生成模板Clock
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Clock 结构体
// 如果含有time.Time 请自行import time包
type Clock struct {
      global.GVA_MODEL
      ClockDate  string `json:"clockDate" form:"clockDate" gorm:"column:clock_date;comment:打卡日期;size:191;"`
      ClockInLocation  string `json:"clockInLocation" form:"clockInLocation" gorm:"column:clock_in_location;comment:上班打卡位置;size:191;"`
      ClockInPic  string `json:"clockInPic" form:"clockInPic" gorm:"column:clock_in_pic;comment:上班打卡照片;size:2000;"`
      ClockInTime  string `json:"clockInTime" form:"clockInTime" gorm:"column:clock_in_time;comment:上班打卡时间;size:191;"`
      ClockOutLocation  string `json:"clockOutLocation" form:"clockOutLocation" gorm:"column:clock_out_location;comment:下班打卡位置;size:191;"`
      ClockOutPic  string `json:"clockOutPic" form:"clockOutPic" gorm:"column:clock_out_pic;comment:下班打卡照片;size:2000;"`
      ClockOutTime  string `json:"clockOutTime" form:"clockOutTime" gorm:"column:clock_out_time;comment:下班打卡时间;size:191;"`
      ClockStatus  string `json:"clockStatus" form:"clockStatus" gorm:"column:clock_status;comment:打卡状态, 迟到/早退/按时上班/按时下班;size:191;"`
      Depart  string `json:"depart" form:"depart" gorm:"column:depart;comment:巡检员部门;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      InspectorName  string `json:"inspectorName" form:"inspectorName" gorm:"column:inspector_name;comment:巡检员姓名;size:191;"`
      InspectorUsername  string `json:"inspectorUsername" form:"inspectorUsername" gorm:"column:inspector_username;comment:巡检员用户名;size:191;"`
      Job  string `json:"job" form:"job" gorm:"column:job;comment:巡检员岗位;size:191;"`
}


// TableName Clock 表名
func (Clock) TableName() string {
  return "safety_clock"
}

