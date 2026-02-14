package main

import (
	"github.com/Amirhos-esm/portfolio/models"
	"github.com/Amirhos-esm/portfolio/views"
	"github.com/Amirhos-esm/portfolio/views/pages"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode("release")
	sv := gin.Default()
	
	sv.Static("/static", "./static")
	data := models.GetMockData()
	sv.GET("/", func(ctx *gin.Context) {
		views.Render(ctx.Writer, ctx.Request, pages.MainPage(&data))
	})
	sv.GET("/project", func(ctx *gin.Context) {
		views.Render(ctx.Writer, ctx.Request, pages.Project(&data.Projects[0]))
	})

	sv.Run(":8080")
}
