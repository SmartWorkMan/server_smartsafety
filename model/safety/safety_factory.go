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
	  Lat  string `json:"lat" form:"lat" gorm:"column:lat;comment:工厂位置纬度;size:191;"`
	  Lng  string `json:"lng" form:"lng" gorm:"column:lng;comment:工厂位置经度;size:191;"`
	  City  string `json:"city" form:"city" gorm:"column:city;comment:省市;size:191;"`
	  Addr  string `json:"addr" form:"addr" gorm:"column:addr;comment:详细地址;size:191;"`
}


// TableName SafetyFactory 表名
func (SafetyFactory) TableName() string {
  return "safety_factory"
}

