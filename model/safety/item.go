// 自动生成模板Item
package safety

import (
      "github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Item 结构体
// 如果含有time.Time 请自行import time包
type Item struct {
      global.GVA_MODEL
      AreaName  string `json:"areaName" form:"areaName" gorm:"column:area_name;comment:巡检区域名称;size:191;"`
      ItemName  string `json:"itemName" form:"itemName" gorm:"column:item_name;comment:巡检事项名称;size:191;"`
      InspectorUsername  string `json:"inspectorUsername" form:"inspectorUsername" gorm:"column:inspector_username;comment:巡检员用户名;size:191;"`
      InspectorName  string `json:"inspectorName" form:"inspectorName" gorm:"column:inspector_name;comment:巡检员姓名;size:191;"`
      Period  string `json:"period" form:"period" gorm:"column:period;comment:巡检周期,day,week,month;size:191;"`
      StartTime  string `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;"`
      EndTime  string `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;"`
      FactoryName  string `json:"factoryName" form:"factoryName" gorm:"column:factory_name;comment:工厂名称;size:191;"`
      Standard  string `json:"standard" form:"standard" gorm:"column:standard;comment:检查标准;size:191;"`
      Enable  int `json:"enable" form:"enable" gorm:"column:enable;comment:启用标志,0:停用 1:启用;"`
      AreaId uint `json:"areaId" form:"areaId" gorm:"column:area_id;comment:巡检区域ID;"`
      ItemSn  string `json:"itemSn" form:"itemSn" gorm:"column:item_sn;comment:巡检事项编码;size:191;"`
}

// TableName Item 表名
func (Item) TableName() string {
      return "safety_item"
}

type ItemNextPeriodDate struct {
      global.GVA_MODEL
      Period  string `json:"period" form:"period" gorm:"column:period;comment:巡检周期,day,week,month;size:191;"`
      NextDate  string `json:"nextDate" form:"nextDate" gorm:"column:next_date;comment:下一个周期开始时间;"`
      Interval  int `json:"interval" form:"interval" gorm:"column:interval;comment:周期间隔;"`
}

func (ItemNextPeriodDate) TableName() string {
      return "safety_item_next_date"
}
