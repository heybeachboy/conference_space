package ws

import (
	"ConferenceSpace/constant"
	"ConferenceSpace/data/ws"
	"ConferenceSpace/util"
)

var HubService *SwSwitch
var OnlineUserMap *util.SafeMap[constant.UID, *ws.UserOnline]

func init() {
	HubService = NewWsSwitch()
	OnlineUserMap = util.NewSafeMap[constant.UID, *ws.UserOnline]()
}
