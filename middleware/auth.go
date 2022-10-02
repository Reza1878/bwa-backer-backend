package middleware

import (
	"bwa-backer/auth"
	"bwa-backer/helper"
	"bwa-backer/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			helper.ResponseUnAuthorized(c, "Unauthorized", nil)
			return
		}

		arrToken := strings.Split(authHeader, " ")
		var tokenString string
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			helper.ResponseUnAuthorized(c, "Unauthorized", nil)
			c.Abort()
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.ResponseUnAuthorized(c, "Unauthorized", nil)
			c.Abort()
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil || user.Id == 0 {
			helper.ResponseUnAuthorized(c, "Unauthorized", nil)
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
