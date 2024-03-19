package ws

var HubService *SwSwitch

func init() {
	HubService = NewWsSwitch()
}
