package controller

import (
	"elderly/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// GetLedStateListHandler 查询烟雾灯光列表信息接口
// @Summary 查询烟雾灯光列表信息接口
// @Description 查询烟雾灯光列表
// @Tags 烟雾灯光列表信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} _ResponseLedList
// @Router /smooth_led [get]
func GetLedStateListHandler(c *gin.Context) {
	data, err := logic.GetLedStateListHandler()
	if err != nil {
		zap.L().Error("logic.GetDhtList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
