package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
)

// @Summary 通知用户上线(10006)
// @Id UserOnline
// @Description 通知用户上线
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.UserOnline true "请求token刷新"
// @Success 0 {object} ws.UserOnline	"处理成功token字符串放在data字段"
// @Router /user/online [post]
func UserOnline(ctx *gin.Context) {
	Param := new(ws.UserOnline)
	util.NewRespJsonBodyWithSuccess(Param)
}
