package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	safetyReq "github.com/flipped-aurora/gin-vue-admin/server/model/safety/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NoticeApi struct {
}


// CreateNotice 创建Notice
// @Tags Notice
// @Summary 创建Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Notice true "创建Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/createNotice [post]
func (noticeApi *NoticeApi) CreateNotice(c *gin.Context) {
	var notice safety.Notice
	_ = c.ShouldBindJSON(&notice)
	if notice.Topic == "" || notice.Type == 0 || notice.OrgName == "" {
		global.GVA_LOG.Error("创建失败!请检查输入!")
		response.FailWithMessage("创建失败!请检查输入!", c)
		return
	}

	if err := noticeService.CreateNotice(notice); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Router /notice/readNotice [post]
func (noticeApi *NoticeApi) ReadNotice(c *gin.Context) {
	var noticeRead safety.NoticeRead
	_ = c.ShouldBindJSON(&noticeRead)
	if noticeRead.Username == "" || noticeRead.NoticeId == 0 {
		global.GVA_LOG.Error("读取通知失败!")
		response.FailWithMessage("读取通知失败", c)
		return
	}
	if err := noticeService.ReadNotice(noticeRead); err != nil {
		global.GVA_LOG.Error("读取通知失败!", zap.Error(err))
		response.FailWithMessage("读取通知失败", c)
	} else {
		response.OkWithMessage("读取通知成功", c)
	}
}

// DeleteNotice 删除Notice
// @Tags Notice
// @Summary 删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Notice true "删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
func (noticeApi *NoticeApi) DeleteNotice(c *gin.Context) {
	var notice safety.Notice
	_ = c.ShouldBindJSON(&notice)
	if err := noticeService.DeleteNotice(notice); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNoticeByIds 批量删除Notice
// @Tags Notice
// @Summary 批量删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /notice/deleteNoticeByIds [delete]
func (noticeApi *NoticeApi) DeleteNoticeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := noticeService.DeleteNoticeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNotice 更新Notice
// @Tags Notice
// @Summary 更新Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body safety.Notice true "更新Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notice/updateNotice [put]
func (noticeApi *NoticeApi) UpdateNotice(c *gin.Context) {
	var notice safety.Notice
	_ = c.ShouldBindJSON(&notice)
	if notice.ID == 0 {
		global.GVA_LOG.Error("更新失败!")
		response.FailWithMessage("更新失败", c)
		return
	}
	if err := noticeService.UpdateNotice(notice); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNotice 用id查询Notice
// @Tags Notice
// @Summary 用id查询Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safety.Notice true "用id查询Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notice/findNotice [get]
func (noticeApi *NoticeApi) FindNotice(c *gin.Context) {
	var notice safety.Notice
	_ = c.ShouldBindQuery(&notice)
	if err, renotice := noticeService.GetNotice(notice.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renotice": renotice}, c)
	}
}

// GetNoticeList 分页获取Notice列表
// @Tags Notice
// @Summary 分页获取Notice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query safetyReq.NoticeSearch true "分页获取Notice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/getNoticeList [post]
func (noticeApi *NoticeApi) GetNoticeList(c *gin.Context) {
	var pageInfo safetyReq.NoticeSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if pageInfo.Username == "" || pageInfo.Type == 0 || pageInfo.OrgName == "" {
		global.GVA_LOG.Error("获取失败!请检查输入!")
		response.FailWithMessage("获取失败!请检查输入!", c)
		return
	}

	var reads []safety.NoticeRead
	var noticeInfoReadList []safetyReq.NoticeInfoAndRead

	if err, notices, total := noticeService.GetNoticeInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
		if err, reads = noticeService.GetNoticeReadList(pageInfo); err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
			return
		}

		for _, noticeInfo := range notices {
			var noticeInfoRead safetyReq.NoticeInfoAndRead
			noticeInfoRead.Notice = noticeInfo
			noticeInfoRead.Username = pageInfo.Username
			noticeInfoRead.IsRead = 0
			for _, read := range reads {
				if read.NoticeId == noticeInfo.ID {
					noticeInfoRead.IsRead = 1
					break
				}
			}
			noticeInfoReadList = append(noticeInfoReadList, noticeInfoRead)
		}

		response.OkWithDetailed(response.PageResult{
            List:     noticeInfoReadList,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
