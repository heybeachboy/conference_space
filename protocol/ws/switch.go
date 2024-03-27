package ws

import (
	"ConferenceSpace/constant"
	"ConferenceSpace/data/ws"
	"ConferenceSpace/logger"
	"context"
	"encoding/json"
	"sync"
	"time"
)

const PushMessageBufferSize = 1000
const NewClientBufferSize = 200

func NewWsSwitch() *SwSwitch {
	ctx, cancel := context.WithCancel(context.Background())
	return &SwSwitch{
		ctx:        ctx,
		cancel:     cancel,
		PushChan:   make(chan *ws.PushPack, PushMessageBufferSize),
		newChan:    make(chan *Client, NewClientBufferSize),
		lock:       new(sync.Mutex),
		sessionMap: make(map[uint32]*Client),
	}
}

type SwSwitch struct {
	ctx        context.Context
	cancel     context.CancelFunc
	PushChan   chan *ws.PushPack
	newChan    chan *Client
	lock       *sync.Mutex
	sessionMap map[uint32]*Client
}

func (s *SwSwitch) NewSession(c *Client) {
	s.newChan <- c
}

func (s *SwSwitch) RemoveSession(c *Client) {
	s.lock.Lock()
	defer s.lock.Unlock()
	conn, ok := s.sessionMap[c.ctx.Uid]
	if !ok {
		return
	}
	conn.Close()
	delete(s.sessionMap, c.ctx.Uid)
}

func (s *SwSwitch) Run() {
	for {
		select {
		case task := <-s.PushChan:
			s.pushMessageToTerminal(task)
		case cl := <-s.newChan:
			if conn, ok := s.sessionMap[cl.ctx.Uid]; ok {
				logger.WaringF("username : %s uid : %s repeat login, kick offline")
				conn.Send(s.getKickOfflineData(conn.ctx.Uid, conn.ctx.Username)) //通知客户端被踢下线
				conn.Close()
			} else {
				HubService.broadcast(s.getOnlineData(cl.ctx.Uid, cl.ctx.Username))
			}
			s.sessionMap[cl.ctx.Uid] = cl
			go cl.reader()
			go cl.writer()
			OnlineUserMap.Set(constant.UID(cl.ctx.Uid), &ws.UserOnline{
				Uid:      constant.UID(cl.ctx.Uid),
				Username: cl.ctx.Username,
				At:       time.Now().Unix(),
			})
			cl.sendOnlineUserList()
		}

	}

}

func (s *SwSwitch) getKickOfflineData(uid uint32, username string) []byte {
	payload := new(ws.UserOffline)
	payload.Uid = uid
	payload.Username = username
	payload.At = time.Now().Unix()
	pack := new(ws.Pack)
	pack.To = uid
	pack.Code = ws.KickOfflineNotify
	pack.Kind = ws.SingleChatPack
	pack.Payload = payload.Marshal()
	return pack.Marshal()
}

func (s *SwSwitch) getOnlineData(uid uint32, username string) []byte {
	payload := new(ws.UserOnline)
	payload.Uid = constant.UID(uid)
	payload.Username = username
	payload.At = time.Now().Unix()
	pack := new(ws.Pack)
	pack.To = uid
	pack.Code = ws.UserOnlineNotify
	pack.Kind = ws.SingleChatPack
	pack.Payload = payload.Marshal()
	return pack.Marshal()
}

func (s *SwSwitch) userOfflineBroadcast(uid uint32, username string) {
	OnlineUserMap.Delete(constant.UID(uid))
	payload := new(ws.UserOffline)
	payload.Uid = uid
	payload.Username = username
	payload.At = time.Now().Unix()
	pack := new(ws.Pack)
	pack.To = uid
	pack.Code = ws.UserOffLineNotify
	pack.Kind = ws.SingleChatPack
	pack.Payload = payload.Marshal()
	s.broadcast(pack.Marshal())
}

func (s *SwSwitch) pushMessageToTerminal(t *ws.PushPack) {
	s.lock.Lock()
	defer s.lock.Unlock()
	connMap, ok := s.sessionMap[t.Uid]
	if !ok {
		logger.DebugF("uid %s is not terminal online payload : %v", t.Uid, t.Msg)
		return
	}
	data, err := json.Marshal(t.Msg)
	if err != nil {
		logger.ErrorF("pushMessageToTerminal error: %v", err)
		return
	}
	connMap.Send(data)
}

/*
*
@Comment 把收到的消息包广播给连接所有终端用户
*/
func (s *SwSwitch) PushToTerminal(p *ws.Pack) error {
	data, err := json.Marshal(p)

	if err != nil {
		logger.ErrorF("push to terminal json marshal packet error : %s", err.Error())
		return err
	}
	s.broadcast(data) //把消息广播给所有连接的终端用户
	return nil
}

func (s *SwSwitch) broadcast(data []byte) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, connMap := range s.sessionMap {
		connMap.Send(data)
	}
}

func (s *SwSwitch) Free() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.cancel()
	for k, _ := range s.sessionMap {
		delete(s.sessionMap, k)
	}

}

func (s *SwSwitch) GetCancelCtx() context.Context {
	return s.ctx
}
