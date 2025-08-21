package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

// api controller
func (pc *PongController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // map string , cặp khóa key-value
		"message": "Hello, World!"})
}
