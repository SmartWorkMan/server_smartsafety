package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SafetyFactoryApi
	InspectorApi
	AreaApi
	ItemApi
	TaskApi
	LawsApi
	TecentApi
	ClockApi
	BasicInfoApi
	BuildingInfoApi
	PlanApi
	LocationLibraryApi
	KeyLocationApi
	TrainingApi
	NoticeApi
	ReportApi
}

var (
	safetyFactoryService   = service.ServiceGroupApp.SafetyServiceGroup.SafetyFactoryService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
    inspectorService       = service.ServiceGroupApp.SafetyServiceGroup.InspectorService
    areaService            = service.ServiceGroupApp.SafetyServiceGroup.AreaService
    itemService            = service.ServiceGroupApp.SafetyServiceGroup.ItemService
    taskService            = service.ServiceGroupApp.SafetyServiceGroup.TaskService
    lawsService            = service.ServiceGroupApp.SafetyServiceGroup.LawsService
    tecentService          = service.ServiceGroupApp.SafetyServiceGroup.TecentService
    clockService           = service.ServiceGroupApp.SafetyServiceGroup.ClockService
    basicInfoService       = service.ServiceGroupApp.SafetyServiceGroup.BasicInfoService
    buildingInfoService    = service.ServiceGroupApp.SafetyServiceGroup.BuildingInfoService
    planService            = service.ServiceGroupApp.SafetyServiceGroup.PlanService
    locationLibraryService = service.ServiceGroupApp.SafetyServiceGroup.LocationLibraryService
    keyLocationService     = service.ServiceGroupApp.SafetyServiceGroup.KeyLocationService
    trainingService        = service.ServiceGroupApp.SafetyServiceGroup.TrainingService
    noticeService          = service.ServiceGroupApp.SafetyServiceGroup.NoticeService
    reportService          = service.ServiceGroupApp.SafetyServiceGroup.ReportService
)
