package ws

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/db/model"
	"ConferenceSpace/db/mysql"
	"ConferenceSpace/libs/agora"
	"ConferenceSpace/util"
	"time"
)

type PayloadIf interface {
	Marshal() []byte
	Unmarshal(data []byte) error
}

type ServiceHandler struct {
	client *Client
}

func (s *ServiceHandler) Ping(p *ws.Pack) {
	ping := new(ws.Ping)
	ping.Time = time.Now().Unix()
	s.reply(p, util.NewRespJsonBodyWithSuccess(ping))
}

func (s *ServiceHandler) RefreshToken(p *ws.Pack) {
	param := new(ws.RtcToken)
	if err := param.Unmarshal(p.Payload); err != nil {
		s.reply(p, util.NewRespJsonBodyWithBadRequest(err.Error()))
		return
	}

	if param.Role < 1 || param.Role > 2 {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("role is illegal and in(1,2)"))
		return
	}

	if param.UidRtcInt < 1 {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("uid is illegal"))
		return
	}

	if param.ChannelName == "" {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("channel name is null"))
		return
	}

	token, err := agora.GenerateRtcToken(param)
	if err != nil {
		s.reply(p, util.NewRespJsonBodyWithInternalError(err.Error()))
		return
	}
	s.reply(p, util.NewRespJsonBodyWithSuccess(token))
}

func (s *ServiceHandler) CreatedRoom(p *ws.Pack) {
	param := new(ws.NewRoom)
	if err := param.Unmarshal(p.Payload); err != nil {
		s.reply(p, util.NewRespJsonBodyWithBadRequest(err.Error()))
		return
	}
	if param.OwnId == s.client.ctx.Uid {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("own id is illegal"))
		return
	}

	if param.ChannelName == "" {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("channel name is null"))
		return
	}

	if param.StartAt < 1 || param.EndAt < 1 || (param.EndAt < param.StartAt) {
		s.reply(p, util.NewRespJsonBodyWithBadRequest("start time is illegal or end time"))
		return
	}

	room := new(model.Room)
	room.OwnId = param.OwnId
	room.ChannelName = param.ChannelName
	room.StartAt = param.StartAt
	room.EndAt = param.EndAt
	room.SpaceId = param.SpaceId
	room.CreatedAt = time.Now().Unix()
	room.UpdatedAt = room.CreatedAt
	if err := mysql.NewRoom(room); err != nil {
		s.reply(p, util.NewRespJsonBodyWithInternalError(err.Error()))
		return
	}
	s.reply(p, util.NewRespJsonBodyWithSuccess("OK"))
}

func (s *ServiceHandler) reply(p *ws.Pack, pay PayloadIf) {
	p.Payload = pay.Marshal()
	s.client.Send(p.Marshal())
}
