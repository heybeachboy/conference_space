package ws

import "encoding/json"

type NewRoom struct {
	SpaceId     uint32 `json:"root_id"`
	OwnId       uint32 `json:"own_id"`
	ChannelName string `json:"channel_name"`
	StartAt     int64  `json:"start_at"`
	EndAt       int64  `json:"end_at"`
}

func (p *NewRoom) Marshal() []byte {
	data, _ := json.Marshal(p)
	return data
}

func (p *NewRoom) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
