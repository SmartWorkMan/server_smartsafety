package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type InspectorSearch struct{
    safety.Inspector
    request.PageInfo
}

type InspectorLogin struct{
	Role  int `json:"role"` //1: inspector 2:maintain user 3: factory user
	FactoryName  string `json:"factoryName"`
}

type InspectorCreate struct{
	safety.Inspector
	CertList  []string `json:"certList"`
}