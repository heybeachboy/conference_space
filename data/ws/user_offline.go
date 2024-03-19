package ws

import "encoding/json"

type UserOffline struct {
	Uid      uint32 `json:"uid"`      //下线用户uid
	Username string `json:"username"` //下线用户名
	At       int64  `json:"time"`     //下线时间
}

func (p *UserOffline) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *UserOffline) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
