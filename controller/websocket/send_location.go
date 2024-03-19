package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
)

// @Summary 用户发送位置变化(10001)
// @Id UserSendLocation
// @Description 用户发送位置变化
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.UserLocation true "请求token刷新"
// @Success 0 {object} ws.UserLocation	"处理成功token字符串放在data字段"
// @Router /user/send/loaction [post]
func UserSendLocation(ctx *gin.Context) {
	Param := new(ws.UserLocation)
	util.NewRespJsonBodyWithSuccess(Param)
}
