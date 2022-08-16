package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"gorm.io/gorm"
	"errors"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type TrainingService struct {
}

// CreateTraining 创建Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService) CreateTraining(inputTraining safety.Training) (err error, outputTraining safety.Training) {
	var training safety.Training
	if !errors.Is(global.GVA_DB.Where("factory_name = ? AND topic = ? AND finish_flag = 1", inputTraining.FactoryName, inputTraining.Topic).First(&training).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在同名培训"), outputTraining
	}
	err = global.GVA_DB.Create(&inputTraining).Error
	if err != nil {
		return err, outputTraining
	}

	err = global.GVA_DB.Where("factory_name = ? AND topic = ? AND finish_flag = 1", inputTraining.FactoryName, inputTraining.Topic).First(&outputTraining).Error
	return err, outputTraining
}

// DeleteTraining 删除Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService)DeleteTraining(training safety.Training) (err error) {	var query safety.Training
	err = global.GVA_DB.Where("id = ?", training.ID).First(&query).Error
	if err != nil {
		return err
	}
	if query.FinishFlag == 2 {
		return errors.New("培训已结束")
	}

	err = global.GVA_DB.Delete(&training).Error
	return err
}

// DeleteTrainingByIds 批量删除Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService)DeleteTrainingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Training{},"id in ?",ids.Ids).Error
	return err
}

func (trainingService *TrainingService)DeleteTrainingByFactory(factoryName string) (err error) {
	err = global.GVA_DB.Delete(&[]safety.Training{},"factory_name = ?", factoryName).Error
	return err
}

// UpdateTraining 更新Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService)UpdateTraining(training safety.Training) (err error) {
	var query safety.Training
	err = global.GVA_DB.Where("id = ?", training.ID).First(&query).Error
	if err != nil {
		return err
	}
	if query.FinishFlag == 2 {
		return errors.New("培训已结束")
	}

	update := safety.Training{
		TrainingType: training.TrainingType,
		StartTime: training.StartTime,
		FinishTime: training.FinishTime,
		Topic: training.Topic,
		Location: training.Location,
		TrainingKind: training.TrainingKind}
	err = global.GVA_DB.Model(&safety.Training{}).Where("id = ?", training.ID).Updates(update).Error
	return err
}

func (trainingService *TrainingService)CreateQRCode(training safety.Training) (err error) {
	err = global.GVA_DB.Model(&safety.Training{}).Where("id = ?", training.ID).Updates(safety.Training{TrainingParam: training.TrainingParam}).Error
	return err
}

func (trainingService *TrainingService)SighInTraining(param string) (err error) {
	var training safety.Training
	err = global.GVA_DB.Where("training_param = ?", param).First(&training).Error
	if err != nil {
		return err
	}
	if training.FinishFlag == 2 {
		return errors.New("会议已结束")
	}

	err = global.GVA_DB.Model(&safety.Training{}).Where("training_param = ?", param).UpdateColumn("number", gorm.Expr("number + ?", 1)).Error
	return err
}

func (trainingService *TrainingService)SubmitTraining(training safety.Training) (err error) {
	curTime := time.Now().Format("2006-01-02 15:04:05")
	updateTraining := safety.Training{
		Description: training.Description,
		StartTime: training.StartTime,
		FinishTime: training.FinishTime,
		FinishFlag: 2,
		Location: training.Location,
		TrainingPic: training.TrainingPic,
		TrainingVideo: training.TrainingVideo,
		Topic: training.Topic,
		SubmitTime: curTime,
		TrainingKind: training.TrainingKind,
	}
	err = global.GVA_DB.Model(&safety.Training{}).Where("id = ?", training.ID).Updates(updateTraining).Error
	return err
}

// GetTraining 根据id获取Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService)GetTraining(id uint) (err error, training safety.Training) {
	err = global.GVA_DB.Where("id = ?", id).First(&training).Error
	return
}

// GetTrainingInfoList 分页获取Training记录
// Author [piexlmax](https://github.com/piexlmax)
func (trainingService *TrainingService)GetTrainingListByFinishFlag(info safetyReq.TrainingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&safety.Training{})
    var trainings []safety.Training
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("factory_name = ? AND training_type = ? AND finish_flag = ?", info.FactoryName, info.TrainingType, info.FinishFlag).Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Order(clause.OrderByColumn{
		Column: clause.Column{Table: clause.CurrentTable, Name: "submit_time"},
		Desc:   true,
	}).Offset(offset).Find(&trainings, "factory_name = ? AND training_type = ? AND finish_flag = ?", info.FactoryName, info.TrainingType, info.FinishFlag).Error
	return err, trainings, total
}

func (trainingService *TrainingService)GetTrainingListByTopicAndTime(info safetyReq.TrainingSearchRange) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&safety.Training{})
	var trainings []safety.Training

	if info.Topic != "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND finish_flag = 2 AND training_type = ? AND topic like ? AND submit_time >= ? AND submit_time <= ?", info.FactoryName, info.TrainingType, "%"+info.Topic+"%", timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "submit_time"},
			Desc:   true,
		}).Offset(offset).Find(&trainings, "factory_name = ? AND finish_flag = 2 AND training_type = ? AND topic like ? AND submit_time >= ? AND submit_time <= ?", info.FactoryName, info.TrainingType, "%"+info.Topic+"%", timeRange[0], timeRange[1]).Error
	} else if info.Topic != "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND finish_flag = 2 AND training_type = ? AND topic like ?", info.FactoryName, info.TrainingType, "%"+info.Topic+"%").Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "submit_time"},
			Desc:   true,
		}).Offset(offset).Find(&trainings, "factory_name = ? AND finish_flag = 2 AND training_type = ? AND topic like ?", info.FactoryName, info.TrainingType, "%"+info.Topic+"%").Error
	} else if info.Topic == "" && info.TimeRange != "" {
		timeRange := strings.Split(info.TimeRange, "~")
		err = db.Where("factory_name = ? AND finish_flag = 2 AND training_type = ? AND submit_time >= ? AND submit_time <= ?", info.FactoryName, info.TrainingType, timeRange[0], timeRange[1]).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "submit_time"},
			Desc:   true,
		}).Offset(offset).Find(&trainings, "factory_name = ? AND finish_flag = 2 AND training_type = ? AND submit_time >= ? AND submit_time <= ?", info.FactoryName, info.TrainingType, timeRange[0], timeRange[1]).Error
	} else if info.Topic == "" && info.TimeRange == "" {
		err = db.Where("factory_name = ? AND finish_flag = 2 AND training_type = ?", info.FactoryName, info.TrainingType).Count(&total).Error
		if err!=nil {
			return
		}
		err = db.Limit(limit).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: "submit_time"},
			Desc:   true,
		}).Offset(offset).Find(&trainings, "factory_name = ? AND finish_flag = 2 AND training_type = ?", info.FactoryName, info.TrainingType).Error
	}

	return err, trainings, total
}

