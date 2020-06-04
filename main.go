package main

import (
	"./troubles"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/images", "./images")
	router.Static("/locales/ja", "./locales/ja")
	router.Static("/locales/en", "./locales/en")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/iof", func(ctx *gin.Context) {
		troubles.GetIntOverflow(ctx)
	})
	router.POST("/iof", func(ctx *gin.Context) {
		troubles.PostIntOverflow(ctx)
	})

	router.Run()
}
