package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type LawsSearch struct{
    safety.Laws
    request.PageInfo
	Sort  string `json:"sort"`
}
