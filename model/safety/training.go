// 自动生成模板Training
package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Training 结构体
// 如果含有time.Time 请自行import time包
type Training struct {
      global.GVA_MODEL
      Description  string `json:"description" form:"description" gorm:"column:description;comment:培训描述;size:191;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      FinishFlag  int `json:"finishFlag" form:"finishFlag" gorm:"column:finish_flag;comment:结束标志,1:未结束, 2:已结束;size:10;"`
      FinishTime  string `json:"finishTime" form:"finishTime" gorm:"column:finish_time;comment:结束时间;size:191;"`
      Location  string `json:"location" form:"location" gorm:"column:location;comment:培训地点;size:191;"`
      Number  int `json:"number" form:"number" gorm:"column:number;comment:参会人数;size:10;"`
      StartTime  string `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;size:191;"`
      Topic  string `json:"topic" form:"topic" gorm:"column:topic;comment:主题;size:191;"`
      TrainingParam  string `json:"trainingParam" form:"trainingParam" gorm:"column:training_param;comment:培训代码;size:191;"`
      TrainingPic  string `json:"trainingPic" form:"trainingPic" gorm:"column:training_pic;comment:培训照片;size:2000;"`
      TrainingType  int `json:"trainingType" form:"trainingType" gorm:"column:training_type;comment:培训类型,1:消防培训, 2:消防演练;size:10;"`
      TrainingVideo  string `json:"trainingVideo" form:"trainingVideo" gorm:"column:training_video;comment:培训视频;size:2000;"`
      SubmitTime  string `json:"submitTime" form:"submitTime" gorm:"column:submit_time;comment:提交时间;size:191;"`
      TrainingKind  int `json:"trainingKind" form:"trainingKind" gorm:"column:training_kind;comment:培训演练种类,1:新员工, 2:定期;size:10;"`
}


// TableName Training 表名
func (Training) TableName() string {
  return "safety_training"
}

