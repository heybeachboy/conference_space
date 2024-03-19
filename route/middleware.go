package route

import (
	"ConferenceSpace/constant"
	"ConferenceSpace/data/ws"
	"ConferenceSpace/logger"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//处理接口认证逻辑

		token := strings.TrimSpace(context.Request.Header.Get("token"))

		if strings.Contains(context.Request.URL.Path, `/conference/api/v1/ws`) && token == "" {
			token = context.DefaultQuery("token", "")
		}
		if token == "" {
			logger.InfoF("request token is empty")
			context.JSON(http.StatusForbidden, util.NewRespJsonBodyWithTokenError("request token is empty"))
			context.Abort()
			return
		}
		t := new(util.JwtService)
		err := t.ParseToken(token)

		if err != nil {
			logger.ErrorF("parse token %s error: %s", token, err.Error())
			context.JSON(http.StatusForbidden, util.NewRespJsonBodyWithTokenError("request token parse error"))
			context.Abort()
			return
		}

		if t.ExpiresAt.Unix() < time.Now().Unix() {
			logger.InfoF("Uid : %s token %s expire : %d", t.Uid, token, t.ExpiresAt.Unix())
			context.JSON(http.StatusForbidden, util.NewRespJsonBodyWithTokenError("token is expire"))
			context.Abort()
			return
		}
		logger.DebugF("token : %v uuid : %s", *t, t.Uid)
		ctx := new(ws.SessionContext)
		ctx.Uid = t.Uid
		ctx.Username = t.UserName
		ctx.Token = token
		ctx.Ip = t.Ip
		ctxStr := ctx.Marshal()

		if ctxStr == "" {
			logger.ErrorF("Marshal ctx session str error : %v", ctx)
			context.JSON(http.StatusForbidden, util.NewRespJsonBodyWithTokenError("create ctx session failed"))
			context.Abort()
		}
		context.Request.Header.Set(constant.SessionCtxHeaderKey, ctxStr)
		context.Next()
	}
}
