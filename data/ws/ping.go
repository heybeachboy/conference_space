package ws

import "encoding/json"

type Ping struct {
	Time int64 `json:"time"` //发送时间
}

func (p *Ping) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *Ping) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
