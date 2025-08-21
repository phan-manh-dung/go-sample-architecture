package middleware

import (
	"go-sample-architecture/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, response.ErrInvalidToken, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
