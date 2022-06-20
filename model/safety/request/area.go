package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AreaSearch struct{
    safety.Area
    request.PageInfo
}

type AreaByInspector struct{
	InspectorUsername  string `json:"inspectorUsername"`
}

