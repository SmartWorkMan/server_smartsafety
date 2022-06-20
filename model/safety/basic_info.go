// 自动生成模板BasicInfo
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BasicInfo 结构体
// 如果含有time.Time 请自行import time包
type BasicInfo struct {
      global.GVA_MODEL
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      InfoJson  string `json:"infoJson" form:"infoJson" gorm:"column:info_json;comment:基本信息JSON;"`
}


// TableName BasicInfo 表名
func (BasicInfo) TableName() string {
  return "safety_basic_info"
}

