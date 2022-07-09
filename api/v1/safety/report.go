package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type ReportApi struct {
}


// CreateReport 创建Report
// @Tags Report
// @Summary 创建Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Report true "创建Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/app/createReport [post]
func (reportApi *ReportApi) CreateReport(c *gin.Context) {
	var report safetyReq.ReportApply
	_ = c.ShouldBindJSON(&report)
	if report.FactoryName == "" || report.Type == 0 || report.Username == "" || report.ApplyTime == ""{
		global.GVA_LOG.Error("创建失败!请检查输入!")
		response.FailWithMessage("创建失败!请检查输入!", c)
		return
	}
	if err, output := reportService.CreateReport(reportApply2Report(report)); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(output.ID,"创建成功", c)
	}
}

// @Router /report/app/applyReport [post]
func (reportApi *ReportApi) ApplyReport(c *gin.Context) {
	var report safetyReq.ReportApply
	_ = c.ShouldBindJSON(&report)
	if report.FactoryName == "" || report.Type == 0 || report.Username == ""{
		global.GVA_LOG.Error("创建失败!请检查输入!")
		response.FailWithMessage("创建失败!请检查输入!", c)
		return
	}
	if err := reportService.ApplyReport(reportApply2Report(report)); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


func reportApply2Report(reportApply safetyReq.ReportApply) safety.Report {
	var report safety.Report
	report = reportApply.Report
	if len(reportApply.ReportPicList) == 0 {
		report.ReportPic = ""
	} else {
		report.ReportPic = ""
		for i := 0; i < len(reportApply.ReportPicList); i++ {
			if i != len(reportApply.ReportPicList) - 1 {
				report.ReportPic += reportApply.ReportPicList[i] + ","
			} else {
				report.ReportPic += reportApply.ReportPicList[i]
			}
		}
	}

	if len(reportApply.ReportVideoList) == 0 {
		report.ReportVideo = ""
	} else {
		report.ReportVideo = ""
		for i := 0; i < len(reportApply.ReportVideoList); i++ {
			if i != len(reportApply.ReportVideoList) - 1 {
				report.ReportVideo += reportApply.ReportVideoList[i] + ","
			} else {
				report.ReportVideo += reportApply.ReportVideoList[i]
			}
		}
	}

	return report
}

func report2ReportApply (reportList []safety.Report) []safetyReq.ReportApply {
	var reportApplyList []safetyReq.ReportApply
	for _, report := range reportList {
		var reportApply safetyReq.ReportApply
		reportApply.Report = report
		reportApply.ReportVideo = ""
		reportApply.ReportPic = ""
		videoList := strings.Split(report.ReportVideo, ",")
		PicList := strings.Split(report.ReportPic, ",")
		reportApply.ReportVideoList = videoList
		reportApply.ReportPicList = PicList
		reportApplyList = append(reportApplyList, reportApply)
	}
	return reportApplyList
}

func formalReport2ReportApply (formalReportList []safety.FormalReport) []safetyReq.ReportApply {
	var reportApplyList []safetyReq.ReportApply
	for _, formalReport := range formalReportList {
		var reportApply safetyReq.ReportApply
		var report safety.Report
		report.GVA_MODEL = formalReport.GVA_MODEL
		report.Content = formalReport.Content
		report.FactoryName = formalReport.FactoryName
		report.ReportPic = formalReport.ReportPic
		report.ReportVideo = formalReport.ReportVideo
		report.Topic = formalReport.Topic
		report.Type = formalReport.Type
		report.Username = formalReport.Username
		report.ApplyTime = formalReport.ApplyTime

		reportApply.Report = report
		reportApply.ReportVideo = ""
		reportApply.ReportPic = ""
		videoList := strings.Split(formalReport.ReportVideo, ",")
		PicList := strings.Split(formalReport.ReportPic, ",")
		if len(videoList) == 1 && videoList[0] == "" {

		} else {
			reportApply.ReportVideoList = videoList
		}
		if len(PicList) == 1 && PicList[0] == "" {

		} else {
			reportApply.ReportPicList = PicList
		}
		reportApplyList = append(reportApplyList, reportApply)
	}
	return reportApplyList
}

// DeleteReport 删除Report
// @Tags Report
// @Summary 删除Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Report true "删除Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /report/app/deleteReport [delete]
func (reportApi *ReportApi) DeleteReport(c *gin.Context) {
	var report safety.Report
	_ = c.ShouldBindJSON(&report)
	if report.ID == 0 {
		global.GVA_LOG.Error("删除失败!报告ID不能为空!")
		response.FailWithMessage("删除失败!报告ID不能为空!", c)
		return
	}
	if err := reportService.DeleteReport(report); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReportByIds 批量删除Report
// @Tags Report
// @Summary 批量删除Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /report/deleteReportByIds [delete]
func (reportApi *ReportApi) DeleteReportByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := reportService.DeleteReportByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReport 更新Report
// @Tags Report
// @Summary 更新Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Report true "更新Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /report/app/updateReport [put]
func (reportApi *ReportApi) UpdateReport(c *gin.Context) {
	var report safetyReq.ReportApply
	_ = c.ShouldBindJSON(&report)
	if report.FactoryName == "" || report.Type == 0 || report.Username == ""{
		global.GVA_LOG.Error("创建失败!请检查输入!")
		response.FailWithMessage("创建失败!请检查输入!", c)
		return
	}
	if err := reportService.UpdateReport(reportApply2Report(report)); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReport 用id查询Report
// @Tags Report
// @Summary 用id查询Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Report true "用id查询Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /report/findReport [get]
func (reportApi *ReportApi) FindReport(c *gin.Context) {
	var report safety.Report
	_ = c.ShouldBindQuery(&report)
	if err, rereport := reportService.GetReport(report.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rereport": rereport}, c)
	}
}

// GetReportList 分页获取Report列表
// @Tags Report
// @Summary 分页获取Report列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.ReportSearch true "分页获取Report列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/app/getReportListByUser [post]
func (reportApi *ReportApi) GetReportListByUser(c *gin.Context) {
	var pageInfo safetyReq.ReportSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" || pageInfo.Username == "" {
		global.GVA_LOG.Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	if err, list, total := reportService.GetReportListByUser(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     report2ReportApply(list.([]safety.Report)),
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// @Router /report/app/getFormalReportListByUser [post]
func (reportApi *ReportApi) GetFormalReportListByUser(c *gin.Context) {
	var pageInfo safetyReq.ReportSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" || pageInfo.Username == "" {
		global.GVA_LOG.Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	if err, list, total := reportService.GetFormalReportListByUser(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     formalReport2ReportApply(list.([]safety.FormalReport)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}


// @Router /report/getFormalReportList [post]
func (reportApi *ReportApi) GetFormalReportList(c *gin.Context) {
	var pageInfo safetyReq.ReportSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" {
		global.GVA_LOG.Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	if err, list, total := reportService.GetFormalReportList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     formalReport2ReportApply(list.([]safety.FormalReport)),
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

