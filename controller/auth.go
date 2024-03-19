package controller

import (
	"ConferenceSpace/constant"
	"ConferenceSpace/data/api"
	"ConferenceSpace/db/model"
	"ConferenceSpace/db/mysql"
	"ConferenceSpace/logger"
	"ConferenceSpace/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary 用户登陆
// @Id Login
// @Description 用于用户登陆
// @Tags 会议注册登陆
// @Accept  json
// @Produce  json
// @Param data body api.RegisterReq true "新用户登陆"
// @Success 0 {object} util.RespJsonBody{data=api.LoginResp}	"处理成功"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /conference/login [post]
func Login(c *gin.Context) {
	param := new(api.RegisterReq)
	if err := c.BindJSON(param); err != nil {
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithBadRequest(err.Error()))
		return
	}
	jwt := new(util.JwtService)
	jwt.Ip = c.RemoteIP()
	resp := new(api.LoginResp)
	user := mysql.GetUserByUsername(param.Username)
	if user.Username == param.Username {
		jwt.Uid = user.Uid
		jwt.UserName = user.Username
		jwt.NickName = user.Nickname
		token, err := jwt.CreateToken()
		if err != nil {
			logger.ErrorF("create user token error : %s", err.Error())
			c.JSON(http.StatusOK, util.NewRespJsonBodyWithInternalError("create token error"))
			return
		}
		resp.Token = token
		resp.UserId = user.Uid
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(resp))
		return
	}
	u := new(model.User)
	//u.UserId = util.GetRandomUserIdString()
	u.Username = param.Username
	u.Password = util.MD5ToString(param.Password)
	u.Gender = param.Gender
	u.Avatar = ""
	u.Nickname = param.Nickname
	u.Email = param.Email
	u.Phone = param.Phone
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = u.CreatedAt
	if err := mysql.CreateUser(u); err != nil {
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithInternalError(err.Error()))
		return
	}
	jwt.Uid = u.Uid
	jwt.UserName = u.Username
	jwt.NickName = u.Nickname
	token, err := jwt.CreateToken()
	if err != nil {
		logger.ErrorF("create user token error : %s", err.Error())
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithInternalError("create token error"))
		return
	}
	resp.Token = token
	resp.UserId = u.Uid
	c.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess(resp))
	return

}

// @Summary 用户注册
// @Id Register
// @Description 用于用户注册
// @Tags 会议注册登陆
// @Accept  json
// @Produce  json
// @Param data body api.RegisterReq true "新用户注册"
// @Success 0 {object} util.RespJsonBody	"处理成功"
// @Failure 4000 {object} util.RespJsonBody "请求参数有误"
// @Failure 4003 {object} util.RespJsonBody "Token处理异常"
// @Failure 5000 {object} util.RespJsonBody "业务内部处理错误"
// @Router /conference/register [post]
func Register(c *gin.Context) {
	param := new(api.RegisterReq)
	if err := c.BindJSON(param); err != nil {
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithBadRequest(err.Error()))
		return
	}

	user := mysql.GetUserByUsername(param.Username)
	if user.Username == param.Username {
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithInternalError("username is exist"))
		return
	}
	u := new(model.User)
	//u.UserId = util.GetRandomUserIdString()
	u.Username = param.Username
	u.Password = util.MD5ToString(param.Password)
	u.Gender = constant.GenderTyp(param.Gender)
	u.Avatar = ""
	u.Nickname = param.Nickname
	u.Email = param.Email
	u.Phone = param.Phone
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = u.CreatedAt
	if err := mysql.CreateUser(u); err != nil {
		c.JSON(http.StatusOK, util.NewRespJsonBodyWithInternalError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.NewRespJsonBodyWithSuccess("+OK"))
	return
}
