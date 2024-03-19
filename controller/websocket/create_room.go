package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 创建房间(10005)
// @Id CreatedRoom
// @Description 用于创建房间
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.NewRoom true "创建房间参数"
// @Success 0 {object} util.RespJsonBody	"处理成功token字符串放在data字段"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /created/room [post]
func CreatedRoom(ctx *gin.Context) {
	param := new(ws.NewRoom)
	ctx.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(param))
}
