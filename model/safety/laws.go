// 自动生成模板Laws
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Laws 结构体
// 如果含有time.Time 请自行import time包
type Laws struct {
      global.GVA_MODEL
      LawName  string `json:"lawName" form:"lawName" gorm:"column:law_name;comment:法律名称;size:191;"`
      LawStatus  string `json:"lawStatus" form:"lawStatus" gorm:"column:law_status;comment:法律状态;size:191;"`
      LawType  string `json:"lawType" form:"lawType" gorm:"column:law_type;comment:法律类型;size:191;"`
      ReleaseTime  string `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;comment:发布时间;size:191;"`
      StoreAddr  string `json:"storeAddr" form:"storeAddr" gorm:"column:store_addr;comment:存储地址;size:191;"`
}


// TableName Laws 表名
func (Laws) TableName() string {
  return "safety_laws"
}

