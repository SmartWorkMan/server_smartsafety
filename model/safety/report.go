// 自动生成模板Report
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Report 结构体
// 如果含有time.Time 请自行import time包
type Report struct {
      global.GVA_MODEL
      ApplyTime  string `json:"applyTime" form:"applyTime" gorm:"column:apply_time;comment:提交时间;size:191;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:报告内容;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      ReportPic  string `json:"reportPic" form:"reportPic" gorm:"column:report_pic;comment:报告照片;size:2000;"`
      ReportVideo  string `json:"reportVideo" form:"reportVideo" gorm:"column:report_video;comment:报告视频;size:2000;"`
      Topic  string `json:"topic" form:"topic" gorm:"column:topic;comment:报告标题;size:1000;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:报告类型,1:日报, 2:周报, 3:季报, 4:年报;size:10;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
}


// TableName Report 表名
func (Report) TableName() string {
  return "safety_report"
}

type FormalReport struct {
      global.GVA_MODEL
      ApplyTime  string `json:"applyTime" form:"applyTime" gorm:"column:apply_time;comment:提交时间;size:191;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:报告内容;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      ReportPic  string `json:"reportPic" form:"reportPic" gorm:"column:report_pic;comment:报告照片;size:2000;"`
      ReportVideo  string `json:"reportVideo" form:"reportVideo" gorm:"column:report_video;comment:报告视频;size:2000;"`
      Topic  string `json:"topic" form:"topic" gorm:"column:topic;comment:报告标题;size:1000;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:报告类型,1:日报, 2:周报, 3:季报, 4:年报;size:10;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
}


// TableName Report 表名
func (FormalReport) TableName() string {
      return "safety_formal_report"
}


