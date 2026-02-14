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
