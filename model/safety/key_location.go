// 自动生成模板KeyLocation
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// KeyLocation 结构体
// 如果含有time.Time 请自行import time包
type KeyLocation struct {
      global.GVA_MODEL
      Attachment  string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:1000;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:191;"`
      DutyOfficer  string `json:"dutyOfficer" form:"dutyOfficer" gorm:"column:duty_officer;comment:责任人;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      LocationImage  string `json:"locationImage" form:"locationImage" gorm:"column:location_image;comment:重点部位照片;size:1000;"`
      LocationName  string `json:"locationName" form:"locationName" gorm:"column:location_name;comment:重点部位名称;size:191;"`
      Place  string `json:"place" form:"place" gorm:"column:place;comment:所在位置;size:191;"`
      AttachmentName  string `json:"attachmentName" form:"attachmentName" gorm:"column:attachment_name;comment:附件名称;size:191;"`
}


// TableName KeyLocation 表名
func (KeyLocation) TableName() string {
  return "safety_key_location"
}

