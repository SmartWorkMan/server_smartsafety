// 自动生成模板Plan
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Plan 结构体
// 如果含有time.Time 请自行import time包
type Plan struct {
      global.GVA_MODEL
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      FileAddr  string `json:"fileAddr" form:"fileAddr" gorm:"column:file_addr;comment:文件地址;size:1000;"`
      FileName  string `json:"fileName" form:"fileName" gorm:"column:file_name;comment:文件名;size:191;"`
      FileType  string `json:"fileType" form:"fileType" gorm:"column:file_type;comment:文件类型, step2:2, step3:3;size:10;"`
      UploadTime  string `json:"uploadTime" form:"uploadTime" gorm:"column:upload_time;comment:上传时间;size:191;"`
}


// TableName Plan 表名
func (Plan) TableName() string {
  return "safety_plan"
}

