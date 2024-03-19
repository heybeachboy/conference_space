package controller

import (
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EchoReq struct {
	Code    int    `json:"code"`
	EchoMsg string `json:"echo_msg"`
}

// @Summary ECHO
// @Id CreateAccount
// @Description 用于测试响应是否正常
// @Tags Reference Space
// @Accept  json
// @Produce  json
// @Param data body EchoReq true "code和echo_msg必填"
// @Success 0 {object} util.RespJsonBody	"处理成功"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /echo [post]
func Echo(c *gin.Context) {
	resp := struct {
		Code int
		Msg  string
	}{0, "Health!"}
	c.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(resp))
}
