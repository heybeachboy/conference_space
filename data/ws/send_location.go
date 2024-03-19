package ws

type UserLocation struct {
	Uid uint32  `json:"uid"` //用户uid
	X   float64 `json:"x"`   //用户x坐标
	Y   float64 `json:"y"`   //用户y坐标
	Z   float64 `json:"z"`   //用户z坐标
}
