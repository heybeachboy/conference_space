// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/conference/api/v1/ws": {
            "get": {
                "description": "用于创建websocket连接/conference/api/v1/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOiIyNzQyNzk5MDcxODgxODEyIiwiTmlja05hbWUiOiJQUFAiLCJVc2VyTmFtZSI6Ikxlb24iLCJJcCI6IjEyNy4wLjAuMSIsImV4cCI6MTcwOTk3MzA1NCwibmJmIjoxNzA5ODg2NjU0LCJpYXQiOjE3MDk4ODY2NTR9.o7vFddZLZ662lr2e5gbe_umohvbdtRa1T1MNuCWiN5Q",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "会议API"
                ],
                "summary": "会议websocket",
                "operationId": "WsServer",
                "parameters": [
                    {
                        "description": "新用户登陆",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.WsProtocol"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功",
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.WsProtocol"
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/conference/login": {
            "post": {
                "description": "用于用户登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "会议注册登陆"
                ],
                "summary": "用户登陆",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "新用户登陆",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_api.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.RespJsonBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ConferenceSpace_data_api.LoginResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/conference/register": {
            "post": {
                "description": "用于用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "会议注册登陆"
                ],
                "summary": "用户注册",
                "operationId": "Register",
                "parameters": [
                    {
                        "description": "新用户注册",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_api.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/created/room": {
            "post": {
                "description": "用于创建房间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "创建房间(10005)",
                "operationId": "CreatedRoom",
                "parameters": [
                    {
                        "description": "创建房间参数",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.NewRoom"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/echo": {
            "post": {
                "description": "用于测试响应是否正常",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reference Space"
                ],
                "summary": "ECHO",
                "operationId": "CreateAccount",
                "parameters": [
                    {
                        "description": "code和echo_msg必填",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.EchoReq"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/ping": {
            "post": {
                "description": "用于Ping测试",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "用于Ping测试(10000)",
                "operationId": "Ping",
                "parameters": [
                    {
                        "description": "用于保持网络链路活跃(建议15-60秒发送一次)",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.Ping"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.RespJsonBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ConferenceSpace_data_ws.Ping"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/refresh/token": {
            "post": {
                "description": "用于刷新token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "用于刷新Token(10004)",
                "operationId": "RefreshToken",
                "parameters": [
                    {
                        "description": "请求token刷新",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.RtcToken"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4000": {
                        "description": "请求参数有误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "4003": {
                        "description": "Token处理异常",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    },
                    "5000": {
                        "description": "业务内部处理错误",
                        "schema": {
                            "$ref": "#/definitions/util.RespJsonBody"
                        }
                    }
                }
            }
        },
        "/user/offline": {
            "post": {
                "description": "通知用户下线",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "通知用户下线(10002)",
                "operationId": "UserOffline",
                "parameters": [
                    {
                        "description": "请求token刷新",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserOffline"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserOffline"
                        }
                    }
                }
            }
        },
        "/user/online": {
            "post": {
                "description": "通知用户上线",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "通知用户上线(10006)",
                "operationId": "UserOnline",
                "parameters": [
                    {
                        "description": "请求token刷新",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserOnline"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserOnline"
                        }
                    }
                }
            }
        },
        "/user/online/list": {
            "post": {
                "description": "发送所有的在线用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "发送所有的在线用户列表(10007)",
                "operationId": "UsersOnlineList",
                "parameters": [
                    {
                        "description": "请求token刷新",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ConferenceSpace_data_ws.UserOnline"
                            }
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ConferenceSpace_data_ws.UserOnline"
                            }
                        }
                    }
                }
            }
        },
        "/user/send/loaction": {
            "post": {
                "description": "用户发送位置变化",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "用户发送位置变化(10001)",
                "operationId": "UserSendLocation",
                "parameters": [
                    {
                        "description": "请求token刷新",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserLocation"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "处理成功token字符串放在data字段",
                        "schema": {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.UserLocation"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ConferenceSpace_data_api.LoginResp": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "用户登陆token",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户user_id",
                    "type": "integer"
                }
            }
        },
        "ConferenceSpace_data_api.RegisterReq": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "nickname",
                "password",
                "phone",
                "username"
            ],
            "properties": {
                "confirm_password": {
                    "description": "确认输入密码",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "gender": {
                    "description": "性别0:female,1:male",
                    "allOf": [
                        {
                            "$ref": "#/definitions/constant.GenderTyp"
                        }
                    ]
                },
                "nickname": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "password": {
                    "description": "登陆密码",
                    "type": "string"
                },
                "phone": {
                    "description": "用户手机",
                    "type": "string"
                },
                "username": {
                    "description": "用户登陆名",
                    "type": "string"
                }
            }
        },
        "ConferenceSpace_data_ws.Content": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "Kind    ContentKind     ` + "`" + `json:\"kind\"` + "`" + `    //消息内容类型113:语音消息,112:图片消息,113:语音消息,114:视频消息,115:发送文件消息,116:发送位置消息,117对方正在输入消息提示,118:消息Quote,119:个人名片消息,120:AT消息",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "ConferenceSpace_data_ws.MsgCode": {
            "type": "integer",
            "enum": [
                10000,
                10001,
                10002,
                10003,
                10004,
                10005,
                10006,
                10007
            ],
            "x-enum-comments": {
                "CreateRoom": "创建文件",
                "KickOfflineNotify": "踢下线",
                "PingTest": "Ping测试",
                "RefreshToken": "刷新token",
                "SendLocation": "发送位置",
                "UserOffLineNotify": "用户离线",
                "UserOnlineList": "所有在线用户",
                "UserOnlineNotify": "用户上线"
            },
            "x-enum-varnames": [
                "PingTest",
                "SendLocation",
                "UserOffLineNotify",
                "KickOfflineNotify",
                "RefreshToken",
                "CreateRoom",
                "UserOnlineNotify",
                "UserOnlineList"
            ]
        },
        "ConferenceSpace_data_ws.NewRoom": {
            "type": "object",
            "properties": {
                "channel_name": {
                    "type": "string"
                },
                "end_at": {
                    "type": "integer"
                },
                "own_id": {
                    "type": "integer"
                },
                "root_id": {
                    "type": "integer"
                },
                "start_at": {
                    "type": "integer"
                }
            }
        },
        "ConferenceSpace_data_ws.PackKind": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-comments": {
                "GroupChatPack": "群组消息数据包",
                "SingleChatPack": "个人消息数据包"
            },
            "x-enum-varnames": [
                "SingleChatPack",
                "GroupChatPack"
            ]
        },
        "ConferenceSpace_data_ws.Ping": {
            "type": "object",
            "properties": {
                "time": {
                    "description": "发送时间",
                    "type": "integer"
                }
            }
        },
        "ConferenceSpace_data_ws.RtcToken": {
            "type": "object",
            "properties": {
                "channel_name": {
                    "description": "频道名",
                    "type": "string"
                },
                "role": {
                    "description": "角色1发部者,2观众",
                    "type": "integer"
                },
                "uid": {
                    "description": "用户userId",
                    "type": "integer"
                }
            }
        },
        "ConferenceSpace_data_ws.UserLocation": {
            "type": "object",
            "properties": {
                "uid": {
                    "description": "用户uid",
                    "type": "integer"
                },
                "x": {
                    "description": "用户x坐标",
                    "type": "number"
                },
                "y": {
                    "description": "用户y坐标",
                    "type": "number"
                },
                "z": {
                    "description": "用户z坐标",
                    "type": "number"
                }
            }
        },
        "ConferenceSpace_data_ws.UserOffline": {
            "type": "object",
            "properties": {
                "time": {
                    "description": "下线时间",
                    "type": "integer"
                },
                "uid": {
                    "description": "下线用户uid",
                    "type": "integer"
                },
                "username": {
                    "description": "下线用户名",
                    "type": "string"
                }
            }
        },
        "ConferenceSpace_data_ws.UserOnline": {
            "type": "object",
            "properties": {
                "time": {
                    "description": "上线时间",
                    "type": "integer"
                },
                "uid": {
                    "description": "上线用户uid",
                    "type": "integer"
                },
                "username": {
                    "description": "上线用户名",
                    "type": "string"
                }
            }
        },
        "ConferenceSpace_data_ws.WsProtocol": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "消息类型",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.MsgCode"
                        }
                    ]
                },
                "from": {
                    "description": "发送人的userId",
                    "type": "integer"
                },
                "kind": {
                    "description": "1表示以送个人消息，2表示发送群组消息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.PackKind"
                        }
                    ]
                },
                "payload": {
                    "description": "消息载体",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ConferenceSpace_data_ws.Content"
                        }
                    ]
                },
                "to": {
                    "description": "到地方人的userId如果*表示广播",
                    "type": "integer"
                }
            }
        },
        "constant.GenderTyp": {
            "type": "integer",
            "enum": [
                1,
                0
            ],
            "x-enum-varnames": [
                "Male",
                "Female"
            ]
        },
        "controller.EchoReq": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "echo_msg": {
                    "type": "string"
                }
            }
        },
        "util.RespCode": {
            "type": "integer",
            "enum": [
                0,
                4000,
                4003,
                5000
            ],
            "x-enum-comments": {
                "BadRequestCode": "请求参数异常",
                "ErrorCode": "异常请求响应码",
                "NormalCode": "正常业务响应码",
                "TokenErrorCode": "token提交异常"
            },
            "x-enum-varnames": [
                "NormalCode",
                "BadRequestCode",
                "TokenErrorCode",
                "ErrorCode"
            ]
        },
        "util.RespJsonBody": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/util.RespCode"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8900",
	BasePath:         "/conference",
	Schemes:          []string{},
	Title:            "Reference Square API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
