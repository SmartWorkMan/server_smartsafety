// 自动生成模板Inspector
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Inspector 结构体
// 如果含有time.Time 请自行import time包
type Inspector struct {
      global.GVA_MODEL
      Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:登录密码;size:191;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:191;"`
      PhoneNumber  string `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;comment:手机号;size:191;"`
      Depart  string `json:"depart" form:"depart" gorm:"column:depart;comment:部门;size:191;"`
      Job  string `json:"job" form:"job" gorm:"column:job;comment:岗位;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
}


// TableName Inspector 表名
func (Inspector) TableName() string {
  return "safety_inspector"
}

