// 自动生成模板SafetyFactory
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SafetyFactory 结构体
// 如果含有time.Time 请自行import time包
type SafetyFactory struct {
      global.GVA_MODEL
      FactoryId  string `json:"factoryId" form:"factoryId" gorm:"column:factory_id;comment:工厂ID;size:64;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
}


// TableName SafetyFactory 表名
func (SafetyFactory) TableName() string {
  return "safety_factory"
}

