package ws

import (
	"ConferenceSpace/data/ws"
	"ConferenceSpace/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync/atomic"
	"time"
)

const (
	writeTimeout = 10 * time.Second
	readTimeout  = 10 * time.Second
	maxBuffer    = 1024
	sendBuffer   = 1024
	sendRunning  = 1
	sendStop     = 0
)

func NewWsClient(ctx context.Context, ws *websocket.Conn, c *ws.SessionContext) *Client {
	cli := &Client{
		conn:      ws,
		ctx:       c,
		cancelCtx: ctx,
		sendChan:  make(chan []byte, 1024),
		isSend:    sendRunning,
	}
	cli.service = &ServiceHandler{
		client: cli,
	}
	return cli
}

type Client struct {
	conn      *websocket.Conn
	ctx       *ws.SessionContext
	cancelCtx context.Context
	sendChan  chan []byte
	isSend    uint32
	service   *ServiceHandler
}

func (c *Client) Close() {
	close(c.sendChan)
	atomic.StoreUint32(&c.isSend, 0)
	if err := c.conn.Close(); err != nil {
		logger.ErrorF("ws service quit error : %s", err.Error())
	}
}

func (c *Client) reader() {
	//c.conn.SetReadDeadline(time.Now().Add(readTimeout))
	defer HubService.userOfflineBroadcast(c.ctx.Uid, c.ctx.Username)
	for {

		select {
		case <-c.cancelCtx.Done():
			logger.DebugF("main process quit close : username %s uid %s ip : %d read task", c.ctx.Username, c.ctx.Uid, c.ctx.Ip)
			return

		default:
			_, data, err := c.conn.ReadMessage()
			if err != nil {
				logger.ErrorF("ws server read message error : %s", err.Error())
				return
			}
			//logger.FDebug("read data : %v", data)
			go c.dispatch(data)

		}

	}

}

// 分派ws的业务处理逻辑
func (c *Client) dispatch(data []byte) {
	p := new(ws.Pack)
	if err := json.Unmarshal(data, p); err != nil {
		logger.ErrorF("uid : %s username : %s  ip : %s Unmarshal terminal ws data error : %s",
			c.ctx.Uid, c.ctx.Username, c.ctx.Ip, data)
		HubService.broadcast(data)
		return
	}
	logger.DebugF("receiver user : %s username : %s  ip : %s data packet : %v",
		c.ctx.Uid, c.ctx.Username, c.ctx.Ip, p)

	switch p.Code {
	case ws.SendLocation:
		c.service.SendLocation(p)
	case ws.PingTest:
		c.service.Ping(p)
	case ws.RefreshToken:
		c.service.RefreshToken(p)
	case ws.UserOffLine:
		c.service.UserOffLine(p)
	case ws.CreateRoom:
		c.service.CreatedRoom(p)
	default:
		logger.ErrorF("protocol code unknow : %d ", p.Code)
		logger.DebugF("reply message : %v", p)
		data, _ = json.Marshal(p)
		c.Send(data)
	}
}

func (c *Client) writer() {
	defer HubService.userOfflineBroadcast(c.ctx.Uid, c.ctx.Username)
	for {

		select {
		case <-c.cancelCtx.Done():
			logger.DebugF("main process quit close : username %s uid %s ip : %d write task", c.ctx.Uid, c.ctx.Uid, c.ctx.Ip)
			c.Close()
			return
		case msg, ok := <-c.sendChan:
			fmt.Println("send", len(msg))
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
				logger.ErrorF("ws service write data  set deadline error : %s", err.Error())
				return
			}

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				logger.ErrorF("ws server next writer error : %s", err.Error())
				return
			}
		}
	}

}

func (c *Client) Send(data []byte) {
	if atomic.LoadUint32(&c.isSend) == sendStop {
		logger.ErrorF("send chan close stop send data : %s", data)
		return
	}
	c.sendChan <- data
}

func (c *Client) GetCtx() *ws.SessionContext {
	return c.ctx
}
