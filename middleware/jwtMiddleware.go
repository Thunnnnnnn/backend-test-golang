package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "missing token",
			})
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil {
			fmt.Println("JWT Error:", err)

			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "token invalid",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID, _ := claims["user_id"].(string)

		c.Set("userId", userID)

		c.Next()
	}
}
