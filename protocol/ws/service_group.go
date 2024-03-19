package ws

import "ConferenceSpace/data/ws"

func (s *ServiceHandler) SendLocation(p *ws.Pack) {
	HubService.broadcast(p.Marshal()) //位置信息广播出去
}

func (s *ServiceHandler) UserOffLine(p *ws.Pack) {
	HubService.broadcast(p.Marshal())
}
