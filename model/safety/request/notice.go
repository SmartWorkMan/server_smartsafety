package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type NoticeSearch struct{
    safety.Notice
    request.PageInfo
	Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
}

type NoticeInfoAndRead struct{
	NoticeCreate
	Username  string `json:"username" form:"username" gorm:"column:username;comment:登录用户名;size:191;"`
	IsRead  uint `json:"IsRead"`
}

type NoticeCreate struct{
	safety.Notice
	AttachmentList  []string `json:"attachmentList"`
}