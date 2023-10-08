package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"cognixus/domain/models"
	"cognixus/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			userId, err := utils.ValidateToken(c, authToken)

			if err != nil {
				c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}
			fmt.Printf("get user id: %v\n", fmt.Sprint(userId))
			c.Request.Header.Set("X-USER-ID", fmt.Sprint(userId))
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
