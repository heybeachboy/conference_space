package ws

import "encoding/json"

type RtcToken struct {
	UidRtcInt   uint32 `json:"uid"`          //用户userId
	ChannelName string `json:"channel_name"` //频道名
	Role        uint32 `json:"role"`         //角色1发部者,2观众
}

func (p *RtcToken) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *RtcToken) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
