package controller

import (
	"elderly/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// GetDhtListHandler 查询温湿度列表信息接口
// @Summary 查询温湿度列表信息接口
// @Description 查询温湿度列表
// @Tags 温湿度列表信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} _ResponseDhtList
// @Router /dht22 [get]
func GetDhtListHandler(c *gin.Context) {
	data, err := logic.GetDhtList()
	if err != nil {
		zap.L().Error("logic.GetDhtList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
