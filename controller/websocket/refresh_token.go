package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 用于刷新Token(10004)
// @Id RefreshToken
// @Description 用于刷新token
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.RtcToken true "请求token刷新"
// @Success 0 {object} util.RespJsonBody	"处理成功token字符串放在data字段"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /refresh/token [post]
func RefreshToken(ctx *gin.Context) {
	param := new(ws.RtcToken)
	ctx.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(param))
}
