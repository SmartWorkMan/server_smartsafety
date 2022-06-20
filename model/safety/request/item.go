package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ItemSearch struct{
    safety.Item
    request.PageInfo
}

type ItemUpdateAndDelete struct{
	safety.Item
	Force int `json:"force"`
}

type ItemUpdateAndDeleteRes struct{
	TaskExist int `json:"taskExist"`
}