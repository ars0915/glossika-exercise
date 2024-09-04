package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ars0915/glossika-exercise/util"
	"github.com/ars0915/glossika-exercise/util/cGin"
)

func jwtCheck(rH HttpHandler) cGin.HandlerFunc {
	return func(c *cGin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
