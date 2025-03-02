package routers

import (
	"github.com/TS0906/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v2 := r.Group("/v2/2025")
	{
		v2.GET("/ping", controller.NewPongController().Pong)          // /v2/2025/ping
		v2.GET("/user/1", controller.NewUserController().GetUserByID) // /v2/2025/user/1
		// v2.PUT("/ping", Pong)
		// v2.PATCH("/ping", Pong)
		// v2.DELETE("/ping", Pong)
		// v2.HEAD("/ping", Pong)
		// v2.OPTIONS("/ping", Pong)
	}
	return r

}