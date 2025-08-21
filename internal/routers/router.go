package routers

import (
	"go-sample-architecture/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // 1 chức năng trong framework của Gin , được dùng để tạo 1 instance version mặc định
	v1 := r.Group("v1/2025")
	{
		v1.GET("/ping", controller.NewPongController().Pong) // v1/2025/ping
	}
	// Tạo v2 ...
	v2 := r.Group("v2/2025")
	{
		v2.GET("/:name", controller.NewUserController().GetName)     // v2/2025/name
		v2.GET("/id", controller.NewUserController().GetIDQueryUser) // v2/2025/id
	}
	return r
}
