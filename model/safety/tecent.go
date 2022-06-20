// 自动生成模板Tecent
package safety

// Tecent 结构体
// 如果含有time.Time 请自行import time包
type Tecent struct {
      TencentId  string `json:"tencentId" form:"tencentId" gorm:"column:tencent_id;comment:ID;size:191;"`
      TencentKey  string `json:"tencentKey" form:"tencentKey" gorm:"column:tencent_key;comment:KEY;size:191;"`
      TencentType  string `json:"tencentType" form:"tencentType" gorm:"column:tencent_type;comment:类型;size:191;"`
}


// TableName Tecent 表名
func (Tecent) TableName() string {
  return "tecent"
}

