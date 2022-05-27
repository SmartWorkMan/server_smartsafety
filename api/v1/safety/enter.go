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
}

var (
	safetyFactoryService = service.ServiceGroupApp.SafetyServiceGroup.SafetyFactoryService
	userService          = service.ServiceGroupApp.SystemServiceGroup.UserService
    inspectorService     = service.ServiceGroupApp.SafetyServiceGroup.InspectorService
    areaService          = service.ServiceGroupApp.SafetyServiceGroup.AreaService
    itemService          = service.ServiceGroupApp.SafetyServiceGroup.ItemService
    taskService          = service.ServiceGroupApp.SafetyServiceGroup.TaskService
)
