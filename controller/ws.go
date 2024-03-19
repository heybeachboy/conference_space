package controller

import (
	"ConferenceSpace/constant"
	ws2 "ConferenceSpace/data/ws"
	logger2 "ConferenceSpace/logger"
	"ConferenceSpace/protocol/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// @Summary 会议websocket
// @Id WsServer
// @Description 用于创建websocket连接/conference/api/v1/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOiIyNzQyNzk5MDcxODgxODEyIiwiTmlja05hbWUiOiJQUFAiLCJVc2VyTmFtZSI6Ikxlb24iLCJJcCI6IjEyNy4wLjAuMSIsImV4cCI6MTcwOTk3MzA1NCwibmJmIjoxNzA5ODg2NjU0LCJpYXQiOjE3MDk4ODY2NTR9.o7vFddZLZ662lr2e5gbe_umohvbdtRa1T1MNuCWiN5Q
// @Tags 会议API
// @Accept  json
// @Produce  json
// @Param data body ws2.WsProtocol true "新用户登陆"
// @Success 0 {object} ws2.WsProtocol	"处理成功"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /conference/api/v1/ws [get]
func WsServer(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ctxStr := ctx.Request.Header.Get(constant.SessionCtxHeaderKey)
	if ctxStr == "" {
		logger2.ErrorF("get session ctx str is null")
		http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sessionCtx := new(ws2.SessionContext)
	if err := sessionCtx.Unmarshal(ctxStr); err != nil {
		logger2.ErrorF("Unmarshal session ctx str error : %s", err.Error())
		http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	logger2.DebugF("output session context: %v", *sessionCtx)

	/*if (sessionCtx.Terminal < common.PcWin || sessionCtx.Terminal > common.TestWeb) ||
		sessionCtx.SessionId == "" || sessionCtx.Uid == "" {
		logger.FError("session ctx exception : %v", *sessionCtx)
		http.Error(ctx.Writer, "Bad Request", http.StatusBadRequest)
		return
	}*/

	upgrader := &websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  2048,
		CheckOrigin: func(r *http.Request) bool {
			return true

		},
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		logger2.ErrorF("http upgrade protocol error : %s", err.Error())
		http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//ctx, cancel := context.WithCancel(context.Background())
	cli := ws.NewWsClient(ws.HubService.GetCancelCtx(), conn, sessionCtx)
	if cli == nil {
		http.Error(ctx.Writer, "internal server error", http.StatusInternalServerError)
		conn.Close()
		return
	}
	ws.HubService.NewSession(cli)
}
