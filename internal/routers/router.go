package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v2 := r.Group("/v2/2025")
	{ // /v2/2025/ping
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}
	return r

}

// api or controller
func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "thinh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{ // map string interface
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"thinh", "thinh2"},
	})
}
