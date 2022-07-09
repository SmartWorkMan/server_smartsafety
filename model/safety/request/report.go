package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReportSearch struct{
    safety.Report
    request.PageInfo
	TimeRange  string `json:"timeRange"` //时间范围,格式"2022-05-26 15:39:00~2022-05-27 00:00:00"
}

type ReportApply struct{
	safety.Report
	ReportPicList  []string `json:"reportPicList"`
	ReportVideoList  []string `json:"reportVideoList"`
}