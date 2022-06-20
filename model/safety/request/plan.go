package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PlanSearch struct{
    safety.Plan
    request.PageInfo
}
