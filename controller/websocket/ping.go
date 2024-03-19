package websocket

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 用于Ping测试(10000)
// @Id Ping
// @Description 用于Ping测试
// @Tags Websocket
// @Accept  json
// @Produce  json
// @Param payload body ws.Ping true "用于保持网络链路活跃(建议15-60秒发送一次)"
// @Success 0 {object} util.RespJsonBody{data=ws.Ping}	"处理成功"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /ping [post]
func Ping(ctx *gin.Context) {
	param := new(ws.Ping)
	ctx.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(param))
}
