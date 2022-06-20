package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type TrainingApi struct {
}


// CreateTraining 创建Training
// @Tags Training
// @Summary 创建Training
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Training true "创建Training"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /training/createTraining [post]
func (trainingApi *TrainingApi) CreateTraining(c *gin.Context) {
	var training safety.Training
	_ = c.ShouldBindJSON(&training)
	if training.FactoryName == "" || training.Topic == "" || training.TrainingType == 0 {
		global.GVA_LOG.Error("创建培训失败!请检查输入!")
		response.FailWithMessage("创建培训失败!请检查输入!", c)
		return
	}
	training.FinishFlag = 1
	if err, created := trainingService.CreateTraining(training); err != nil {
        global.GVA_LOG.Error("创建培训失败!", zap.Error(err))
		response.FailWithMessage("创建培训失败", c)
	} else {
		response.OkWithDetailed(created, "创建培训成功", c)
	}
}

// DeleteTraining 删除Training
// @Tags Training
// @Summary 删除Training
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Training true "删除Training"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /training/deleteTraining [delete]
func (trainingApi *TrainingApi) DeleteTraining(c *gin.Context) {
	var training safety.Training
	_ = c.ShouldBindJSON(&training)
	if err := trainingService.DeleteTraining(training); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTrainingByIds 批量删除Training
// @Tags Training
// @Summary 批量删除Training
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Training"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /training/deleteTrainingByIds [delete]
func (trainingApi *TrainingApi) DeleteTrainingByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := trainingService.DeleteTrainingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTraining 更新Training
// @Tags Training
// @Summary 更新Training
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Training true "更新Training"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /training/updateTraining [put]
func (trainingApi *TrainingApi) UpdateTraining(c *gin.Context) {
	var training safety.Training
	_ = c.ShouldBindJSON(&training)
	if training.ID == 0 || training.Topic == "" || training.TrainingType == 0 {
		global.GVA_LOG.Error("修改培训失败!请检查输入!")
		response.FailWithMessage("修改培训失败!请检查输入!", c)
		return
	}
	if err := trainingService.UpdateTraining(training); err != nil {
        global.GVA_LOG.Error("修改培训失败!", zap.Error(err))
		response.FailWithMessage("修改培训失败", c)
	} else {
		response.OkWithMessage("修改培训成功", c)
	}
}

// @Router /training/createQRCode [put]
func (trainingApi *TrainingApi) CreateQRCode(c *gin.Context) {
	var training safety.Training
	_ = c.ShouldBindJSON(&training)
	if training.ID == 0 || training.TrainingParam == "" {
		global.GVA_LOG.Error("生成二维码失败!")
		response.FailWithMessage("生成二维码失败", c)
		return
	}
	if err := trainingService.CreateQRCode(training); err != nil {
		global.GVA_LOG.Error("生成二维码失败!", zap.Error(err))
		response.FailWithMessage("生成二维码失败", c)
	} else {
		response.OkWithMessage("生成二维码成功", c)
	}
}

// @Router /training/sighInTraining [get]
func (trainingApi *TrainingApi) SighInTraining(c *gin.Context) {
	param := c.Query("param")
	if param == "" {
		global.GVA_LOG.Error("签到失败!请输入有效参数!")
		//response.FailWithMessage("签到失败!请输入有效参数!", c)
		response.Redirect(commval.SighInRedirectFailed, "签到失败!请输入有效参数!", c)
		return
	}
	if err := trainingService.SighInTraining(param); err != nil {
		global.GVA_LOG.Error("签到失败!", zap.Error(err))
		if err.Error() == "会议已结束" {
			response.Redirect(commval.SighInRedirectOver, "会议已结束", c)
		} else {
			response.Redirect(commval.SighInRedirectFailed, "签到失败!", c)
		}
	} else {
		//response.OkWithMessage("签到成功", c)
		response.Redirect(commval.SighInRedirectSuccess, "签到成功", c)
	}
}

// @Router /training/submitTraining [put]
func (trainingApi *TrainingApi) SubmitTraining(c *gin.Context) {
	var trainingSubmit safetyReq.TrainingSubmit
	_ = c.ShouldBindJSON(&trainingSubmit)
	if trainingSubmit.ID == 0 {
		global.GVA_LOG.Error("提交培训失败!培训ID不能为空!")
		response.FailWithMessage("提交培训失败!培训ID不能为空!", c)
		return
	}
	if err := trainingService.SubmitTraining(trainingSubmit2Training(trainingSubmit)); err != nil {
		global.GVA_LOG.Error("提交培训失败!", zap.Error(err))
		response.FailWithMessage("提交培训失败", c)
	} else {
		response.OkWithMessage("提交培训成功", c)
	}
}

func trainingSubmit2Training(trainingSubmit safetyReq.TrainingSubmit) safety.Training {
	var training safety.Training
	training = trainingSubmit.Training
	if len(trainingSubmit.TrainingPicList) == 0 {
		training.TrainingPic = ""
	} else {
		training.TrainingPic = ""
		for i := 0; i < len(trainingSubmit.TrainingPicList); i++ {
			if i != len(trainingSubmit.TrainingPicList) - 1 {
				training.TrainingPic += trainingSubmit.TrainingPicList[i] + ","
			} else {
				training.TrainingPic += trainingSubmit.TrainingPicList[i]
			}
		}
	}

	if len(trainingSubmit.TrainingVideoList) == 0 {
		training.TrainingVideo = ""
	} else {
		training.TrainingVideo = ""
		for i := 0; i < len(trainingSubmit.TrainingVideoList); i++ {
			if i != len(trainingSubmit.TrainingVideoList) - 1 {
				training.TrainingVideo += trainingSubmit.TrainingVideoList[i] + ","
			} else {
				training.TrainingVideo += trainingSubmit.TrainingVideoList[i]
			}
		}
	}

	training.FinishFlag = 2

	return training
}

func training2TrainingSubmit (trainingList []safety.Training) []safetyReq.TrainingSubmit {
	var trainingSubmitList []safetyReq.TrainingSubmit
	for _, training := range trainingList {
		var trainingSubmit safetyReq.TrainingSubmit
		trainingSubmit.Training = training
		trainingSubmit.TrainingVideo = ""
		trainingSubmit.TrainingPic = ""
		videoList := strings.Split(training.TrainingVideo, ",")
		PicList := strings.Split(training.TrainingPic, ",")
		trainingSubmit.TrainingVideoList = videoList
		trainingSubmit.TrainingPicList = PicList
		trainingSubmitList = append(trainingSubmitList, trainingSubmit)
	}
	return trainingSubmitList
}

// FindTraining 用id查询Training
// @Tags Training
// @Summary 用id查询Training
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Training true "用id查询Training"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /training/findTraining [post]
func (trainingApi *TrainingApi) FindTraining(c *gin.Context) {
	var training safety.Training
	_ = c.ShouldBindJSON(&training)
	if err, retraining := trainingService.GetTraining(training.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"training": retraining}, c)
	}
}

// GetTrainingList 分页获取Training列表
// @Tags Training
// @Summary 分页获取Training列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.TrainingSearch true "分页获取Training列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /training/getTrainingListByFinishFlag [post]
func (trainingApi *TrainingApi) GetTrainingListByFinishFlag(c *gin.Context) {
	var pageInfo safetyReq.TrainingSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" || pageInfo.TrainingType == 0 || pageInfo.FinishFlag == 0 {
		global.GVA_LOG.Error("查询培训失败!请检查输入!")
		response.FailWithMessage("查询培训失败!请检查输入!", c)
		return
	}
	if err, list, total := trainingService.GetTrainingListByFinishFlag(pageInfo); err != nil {
	    global.GVA_LOG.Error("查询培训失败!", zap.Error(err))
        response.FailWithMessage("查询培训失败", c)
    } else {
		newList := training2TrainingSubmit(list.([]safety.Training))
        response.OkWithDetailed(response.PageResult{
            List:     newList,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "查询培训成功", c)
    }
}

// @Router /training/getTrainingListByTopicAndTime [post]
func (trainingApi *TrainingApi) GetTrainingListByTopicAndTime(c *gin.Context) {
	var pageInfo safetyReq.TrainingSearchRange
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.FactoryName == "" || pageInfo.TrainingType == 0{
		global.GVA_LOG.Error("查询培训失败!请检查输入!")
		response.FailWithMessage("查询培训失败!请检查输入!", c)
		return
	}
	if err, list, total := trainingService.GetTrainingListByTopicAndTime(pageInfo); err != nil {
		global.GVA_LOG.Error("查询培训失败!", zap.Error(err))
		response.FailWithMessage("查询培训失败", c)
	} else {
		newList := training2TrainingSubmit(list.([]safety.Training))
		response.OkWithDetailed(response.PageResult{
			List:     newList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "查询培训成功", c)
	}
}
