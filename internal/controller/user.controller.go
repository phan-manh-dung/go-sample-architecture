package controller

import (
	"go-sample-architecture/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController { // có thể trong user controller có nhiều hàm service
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service
// ur user repository
// um user model

func (uc *UserController) GetName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"Tên là:": "" + name})
}

func (uc *UserController) GetIDQueryUser(c *gin.Context) {
	id := c.Query("id") // http://localhost:8080/v2/2025/id\?id=999123
	c.JSON(http.StatusOK, gin.H{
		"ID là:":  "" + id,
		"users":   []string{"Hiệp", "Dũng"},
		"message": uc.userService.GetInfoUserService(),
	})
}
