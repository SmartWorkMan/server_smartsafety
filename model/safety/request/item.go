package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ItemSearch struct{
    safety.Item
    request.PageInfo
}

type ItemInspector struct {
	InspectorUsername  string `json:"inspectorUsername"`
	InspectorName  string `json:"inspectorName"`
}

type ItemCreate struct{
	safety.Item
	InspectorList []ItemInspector `json:"inspectorList"`
}

type ItemUpdateAndDelete struct{
	ItemCreate
	Force int `json:"force"`
}

type ItemUpdateAndDeleteRes struct{
	TaskExist int `json:"taskExist"`
}