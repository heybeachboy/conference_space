package route

import (
	"ConferenceSpace/controller"
	"github.com/gin-gonic/gin"
)

func init() {

}
func RegisterRoute() *gin.Engine {
	route := gin.Default()
	base := route.Group("/conference")
	{
		base.GET("/echo", controller.Echo)
		base.POST("/login", controller.Login)
		base.POST("/register", controller.Register)
	}

	api := base.Group("/api/v1", AuthMiddleware())
	{
		api.GET("/ws", controller.WsServer)
	}
	return route
}
