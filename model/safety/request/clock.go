package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ClockSearch struct{
    safety.Clock
    request.PageInfo
}

type SubmitClock struct {
	ClockPic  string `json:"clockPic"`
	ClockTime  string `json:"clockTime"`
	ClockType  string `json:"clockType"`
	FactoryName  string `json:"factoryName"`
	InspectorUsername  string `json:"inspectorUsername"`
	Location  string `json:"location"`
}
