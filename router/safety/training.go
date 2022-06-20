package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TrainingRouter struct {
}

// InitTrainingRouter 初始化 Training 路由信息
func (s *TrainingRouter) InitTrainingRouter(Router *gin.RouterGroup) {
	trainingRouter := Router.Group("training").Use(middleware.OperationRecord())
	trainingRouterWithoutRecord := Router.Group("training")
	var trainingApi = v1.ApiGroupApp.SafetyApiGroup.TrainingApi
	{
		trainingRouter.POST("createTraining", trainingApi.CreateTraining)   // 新建Training
		trainingRouter.DELETE("deleteTraining", trainingApi.DeleteTraining) // 删除Training
		//trainingRouter.DELETE("deleteTrainingByIds", trainingApi.DeleteTrainingByIds) // 批量删除Training
		trainingRouter.PUT("updateTraining", trainingApi.UpdateTraining)    // 更新Training
		trainingRouter.PUT("createQRCode", trainingApi.CreateQRCode)    // 更新Training
		trainingRouter.GET("sighInTraining", trainingApi.SighInTraining)    // 更新Training
		trainingRouter.PUT("submitTraining", trainingApi.SubmitTraining)    // 更新Training
	}
	{
		//trainingRouterWithoutRecord.GET("findTraining", trainingApi.FindTraining)        // 根据ID获取Training
		trainingRouterWithoutRecord.POST("getTrainingListByFinishFlag", trainingApi.GetTrainingListByFinishFlag)  // 获取Training列表
		trainingRouterWithoutRecord.POST("getTrainingListByTopicAndTime", trainingApi.GetTrainingListByTopicAndTime)  // 获取Training列表
	}
}
