// 自动生成模板LocationLibrary
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LocationLibrary 结构体
// 如果含有time.Time 请自行import time包
type LocationLibrary struct {
      global.GVA_MODEL
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      LocationName  string `json:"locationName" form:"locationName" gorm:"column:location_name;comment:重点部位名称;size:191;"`
}


// TableName LocationLibrary 表名
func (LocationLibrary) TableName() string {
  return "safety_location_library"
}

