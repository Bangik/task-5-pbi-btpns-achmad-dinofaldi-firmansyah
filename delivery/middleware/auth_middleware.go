package middleware

import (
	"net/http"
	"strings"

	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/security"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		err := c.ShouldBindHeader(&header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid header " + err.Error(),
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(header.AuthorizationHeader, "Bearer ", "", 1)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token string",
			})
			c.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token " + err.Error(),
			})
			c.Abort()
			return
		}

		if claims["id"] != c.Param("id") {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "You don't have permission to access this resource",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
