package api

import "ConferenceSpace/constant"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserId uint32 `json:"user_id"` //用户user_id
	Token  string `json:"token"`   //用户登陆token
}

type RegisterReq struct {
	Username        string             `json:"username" binding:"required"`         //用户登陆名
	Nickname        string             `json:"nickname" binding:"required"`         //用户昵称
	Phone           string             `json:"phone" binding:"required"`            //用户手机
	Email           string             `json:"email" binding:"required"`            //用户邮箱
	Password        string             `json:"password" binding:"required"`         //登陆密码
	ConfirmPassword string             `json:"confirm_password" binding:"required"` //确认输入密码
	Gender          constant.GenderTyp `json:"gender"`                              //性别0:female,1:male
}
