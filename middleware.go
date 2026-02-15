package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, claims, err := app.auth.GetTokenFromHeaderAndVerify(c.Writer, c.Request)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errors": []gin.H{
					{
						"message": "Unauthorized",
						"extensions": gin.H{
							"code": "UNAUTHENTICATED",
						},
					},
				},
			})
			return
		}

		c.Set("id", claims.Subject)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
