package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
)

// @Summary 通知用户下线(10002)
// @Id UserOffline
// @Description 通知用户下线
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.UserOffline true "请求token刷新"
// @Success 0 {object} ws.UserOffline	"处理成功token字符串放在data字段"
// @Router /user/offline [post]
func UserOffline(ctx *gin.Context) {
	Param := new(ws.UserOffline)
	util.NewRespJsonBodyWithSuccess(Param)
}
