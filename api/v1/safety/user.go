package safety

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetCurUser(c *gin.Context) (error, *system.SysUser) {
	selfUserInfo := utils.GetUserInfo(c)
	var curUser *system.SysUser
	var err error
	err, curUser = userService.FindUserById(int(selfUserInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("获取当前用户失败!", zap.Error(err))
		response.FailWithMessage("获取当前用户失败", c)
		return err, nil
	}

	if curUser.FactoryName == "" {
		global.GVA_LOG.Error("获取当前用户失败!当前用户工厂名称为空")
		response.FailWithMessage("获取当前用户失败!当前用户工厂名称为空", c)
		return err, nil
	}

	return nil, curUser
}