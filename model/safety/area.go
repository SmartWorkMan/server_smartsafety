// 自动生成模板Area
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Area 结构体
// 如果含有time.Time 请自行import time包
type Area struct {
      global.GVA_MODEL
      AreaName  string `json:"areaName" form:"areaName" gorm:"column:area_name;comment:巡检区域名称;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      ParentId int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点ID;"`
}


// TableName Area 表名
func (Area) TableName() string {
  return "safety_area"
}

