package ws

import (
	"ConferenceSpace/constant"
	"encoding/json"
)

type UserOnline struct {
	Uid      constant.UID `json:"uid"`      //上线用户uid
	Username string       `json:"username"` //上线用户名
	At       int64        `json:"time"`     //上线时间
}

func (p *UserOnline) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *UserOnline) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
