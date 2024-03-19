package util

import "encoding/json"

type RespCode int

const (
	NormalCode     RespCode = 0    //正常业务响应码
	BadRequestCode RespCode = 4000 //请求参数异常
	TokenErrorCode RespCode = 4003 //token提交异常
	ErrorCode      RespCode = 5000 //异常请求响应码
)

type RespJsonBody struct {
	Code    RespCode `json:"code"`
	Message string   `json:"message"`
	Data    any      `json:"data"`
}

func (p *RespJsonBody) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *RespJsonBody) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

func NewRespJsonBody(code RespCode, msg string, data interface{}) *RespJsonBody {
	var resp = &RespJsonBody{Code: code, Message: msg, Data: data}

	return resp
}

func NewRespJsonBodyWithSuccess(data interface{}) *RespJsonBody {
	var resp = &RespJsonBody{Code: NormalCode, Message: "Successful", Data: data}

	return resp
}

func NewRespJsonBodyWithInternalError(msg string) *RespJsonBody {
	var resp = &RespJsonBody{Code: ErrorCode, Message: msg, Data: "internal error"}
	return resp
}

func NewRespJsonBodyWithBadRequest(msg string) *RespJsonBody {
	var resp = &RespJsonBody{Code: BadRequestCode, Message: msg, Data: "bad request"}
	return resp
}
func NewRespJsonBodyWithTokenError(msg string) *RespJsonBody {
	var resp = &RespJsonBody{Code: TokenErrorCode, Message: msg, Data: "token error"}
	return resp
}
