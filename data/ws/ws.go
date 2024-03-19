package ws

import (
	logger2 "ConferenceSpace/logger"
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
)

type SessionContext struct {
	Uid      uint32
	RoomId   string
	Username string
	Ip       string
	Token    string
}

func (s *SessionContext) Marshal() string {
	buf := new(bytes.Buffer)
	en := gob.NewEncoder(buf)
	if err := en.Encode(s); err != nil {
		logger2.ErrorF("session ctx encode error : %s", err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func (s *SessionContext) Unmarshal(ctx string) error {
	byts, err := base64.StdEncoding.DecodeString(ctx)

	if err != nil {
		return err
	}

	de := gob.NewDecoder(bytes.NewBuffer(byts))
	if err = de.Decode(s); err != nil {
		return err
	}

	return nil
}

type PackKind uint8 //数据包类型

const (
	SingleChatPack PackKind = 1 //个人消息数据包
	GroupChatPack  PackKind = 2 //群组消息数据包
)

type MsgCode int

const (
	PingTest          MsgCode = 10000 //Ping测试
	SendLocation      MsgCode = 10001 //发送位置
	UserOffLineNotify MsgCode = 10002 //用户离线
	KickOfflineNotify MsgCode = 10003 //踢下线
	RefreshToken      MsgCode = 10004 //刷新token
	CreateRoom        MsgCode = 10005 //创建文件
	UserOnlineNotify  MsgCode = 10006 //用户上线
	UserOnlineList    MsgCode = 10007 //所有在线用户
)

type ContentKind uint8

const (
	Text     ContentKind = 111 //文本消息
	Picture  ContentKind = 112 //图片消息
	Voice    ContentKind = 113 //语音消息
	Video    ContentKind = 114 //视频消息
	File     ContentKind = 115 //发送文件消息
	Location ContentKind = 116 //发送位置消息
	Typing   ContentKind = 117 //对方正在输入消息提示
	Quote    ContentKind = 118 //消息Quote
	Card     ContentKind = 119 //个人名片消息
	AtText   ContentKind = 120 //AT消息
)

type Pack struct {
	From    uint32          `json:"from"`    //发送人的userId
	To      uint32          `json:"to"`      //到地方人的userId如果*表示广播
	Kind    PackKind        `json:"kind"`    //1表示以送个人消息，2表示发送群组消息
	Code    MsgCode         `json:"code"`    //消息code
	Payload json.RawMessage `json:"payload"` //消息载荷，具体类型由消息而定
}

func (p *Pack) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *Pack) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

type PushPack struct {
	Uid uint32
	Msg *Pack
}

type Content struct {
	//Kind    ContentKind     `json:"kind"`    //消息内容类型113:语音消息,112:图片消息,113:语音消息,114:视频消息,115:发送文件消息,116:发送位置消息,117对方正在输入消息提示,118:消息Quote,119:个人名片消息,120:AT消息
	Content json.RawMessage `json:"content"` //消息内容实体结构，具体根据业务需要定义
}

type WsProtocol struct {
	From    uint32   `json:"from"`    //发送人的userId
	To      uint32   `json:"to"`      //到地方人的userId如果*表示广播
	Kind    PackKind `json:"kind"`    //1表示以送个人消息，2表示发送群组消息
	Code    MsgCode  `json:"code"`    //消息类型
	Payload Content  `json:"payload"` //消息载体
}
