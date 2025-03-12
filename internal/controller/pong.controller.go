package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

// api or controller
func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("---> My Handler")
	name := c.DefaultQuery("name", "thinh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{ // map string interface
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"thinh", "thinh2"},
	})
}
