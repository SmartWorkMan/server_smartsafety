// 自动生成模板BuildingInfo
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BuildingInfo 结构体
// 如果含有time.Time 请自行import time包
type BuildingInfo struct {
      global.GVA_MODEL
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      InfoJson  string `json:"infoJson" form:"infoJson" gorm:"column:info_json;comment:基本信息JSON;"`
}


// TableName BuildingInfo 表名
func (BuildingInfo) TableName() string {
  return "safety_building_info"
}

