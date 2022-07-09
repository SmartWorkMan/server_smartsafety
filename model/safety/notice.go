// 自动生成模板Notice
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Notice 结构体
// 如果含有time.Time 请自行import time包
type Notice struct {
      global.GVA_MODEL
      Content  string `json:"content" form:"content" gorm:"column:content;comment:通知内容;"`
      OrgName  string `json:"orgName" form:"orgName" gorm:"column:org_name;comment:单位名称;size:191;"`
      NoticeTime  string `json:"noticeTime" form:"noticeTime" gorm:"column:notice_time;comment:发布时间;size:191;"`
      Topic  string `json:"topic" form:"topic" gorm:"column:topic;comment:通知标题;size:1000;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:通知类型,1:消防单位通知, 2:工厂通知;size:10;"`
}


// TableName Notice 表名
func (Notice) TableName() string {
  return "safety_notice"
}

type NoticeRead struct {
      global.GVA_MODEL
      NoticeId  uint `json:"noticeId" form:"noticeId" gorm:"column:notice_id;comment:通知ID;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
}

func (NoticeRead) TableName() string {
      return "safety_notice_read"
}

