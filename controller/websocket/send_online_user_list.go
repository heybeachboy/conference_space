package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
)

type UserOnlineList []*ws.UserOnline

// @Summary 发送所有的在线用户列表(10007)
// @Id UsersOnlineList
// @Description 发送所有的在线用户列表
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body UserOnlineList true "请求token刷新"
// @Success 0 {object} UserOnlineList	"处理成功token字符串放在data字段"
// @Router /user/online/list [post]
func UsersOnlineList(ctx *gin.Context) {
	Param := new([]*ws.UserOnline)
	util.NewRespJsonBodyWithSuccess(Param)
}
