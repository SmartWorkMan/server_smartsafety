package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type ReportService struct {
}

// CreateReport 创建Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService) CreateReport(report safety.Report) (err error, outputReport safety.Report) {
	err = global.GVA_DB.Create(&report).Error
	if err != nil {
		return err, outputReport
	}

	err = global.GVA_DB.Where("factory_name = ? AND username = ? AND apply_time = ?", report.FactoryName, report.Username, report.ApplyTime).First(&outputReport).Error
	return err, outputReport
}

func (reportService *ReportService) ApplyReport(report safety.Report) (err error) {
	var formalReport safety.FormalReport
	formalReport.Content = report.Content
	formalReport.FactoryName = report.FactoryName
	formalReport.ReportPic = report.ReportPic
	formalReport.ReportVideo = report.ReportVideo
	formalReport.Topic = report.Topic
	formalReport.Type = report.Type
	formalReport.Username = report.Username
	formalReport.ApplyTime = time.Now().Format("2006-01-02 15:04:05")

	err = global.GVA_DB.Create(&formalReport).Error
	return err
}

// DeleteReport 删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)DeleteReport(report safety.Report) (err error) {
	err = global.GVA_DB.Delete(&report).Error
	return err
}

// DeleteReportByIds 批量删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)DeleteReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Report{},"id in ?",ids.Ids).Error
	return err
}

func (reportService *ReportService)DeleteReportByFactory(factoryName string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Report{},"factory_name = ?", factoryName).Error
	return err
}

// UpdateReport 更新Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)UpdateReport(report safety.Report) (err error) {
	update := safety.Report{
		Type: report.Type,
		Topic: report.Topic,
		Content: report.Content,
		ApplyTime: report.ApplyTime,
		Username: report.Username,
		ReportPic: report.ReportPic,
		ReportVideo: report.ReportVideo,
	}
	err = global.GVA_DB.Model(&safety.Report{}).Where("id = ?", report.ID).Updates(update).Error
	return err
}

// GetReport 根据id获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)GetReport(id uint) (err error, report safety.Report) {
	err = global.GVA_DB.Where("id = ?", id).First(&report).Error
	return
}

// GetReportInfoList 分页获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)GetReportListByUser(info safetyReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Report{})
    var reports []safety.Report
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? AND username = ?", info.FactoryName, info.Username).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
		Desc:   true,
	}).Offset(offset).Find(&reports, "factory_name = ? AND username = ?", info.FactoryName, info.Username).Error
	return err, reports, total
}

func (reportService *ReportService)GetFormalReportListByUser(info safetyReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.FormalReport{})
	var reports []safety.FormalReport
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? AND username = ?", info.FactoryName, info.Username).Count(&total).Error
	if err!=nil {
		return
	}
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
		Desc:   true,
	}).Offset(offset).Find(&reports, "factory_name = ? AND username = ?", info.FactoryName, info.Username).Error
	return err, reports, total
}

func (reportService *ReportService)GetFormalReportList(info safetyReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.FormalReport{})
	var reports []safety.FormalReport

	if info.Type == 0 && info.Username == "" && info.TimeRange == "" {
		err = db.Where("factory_name = ?", info.FactoryName).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ?", info.FactoryName).Error
	} else if info.Type != 0 && info.Username == "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND type = ?", info.FactoryName, info.Type).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND type = ?", info.FactoryName, info.Type).Error
	} else if info.Type == 0 && info.Username != "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND username = ?", info.FactoryName, info.Username).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND username = ?", info.FactoryName, info.Username).Error
	} else if info.Type == 0 && info.Username == "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, timeRange[0], timeRange[1]).Error
	} else if info.Type != 0 && info.Username != "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND type = ? AND username = ?", info.FactoryName, info.Type, info.Username).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND type = ? AND username = ?", info.FactoryName, info.Type, info.Username).Error
	} else if info.Type != 0 && info.Username == "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND type = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Type, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND type = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Type, timeRange[0], timeRange[1]).Error
	} else if info.Type == 0 && info.Username != "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND username = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Username, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND username = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Username, timeRange[0], timeRange[1]).Error
	}else if info.Type != 0 && info.Username != "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND type = ? AND username = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Type, info.Username, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "apply_time"},
			Desc:   true,
		}).Offset(offset).Find(&reports, "factory_name = ? AND type = ? AND username = ? AND apply_time >= ? AND apply_time <= ?", info.FactoryName, info.Type, info.Username, timeRange[0], timeRange[1]).Error
	}

	return err, reports, total
}
