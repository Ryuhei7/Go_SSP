package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("vies/*.html")
	router.Static("/assets", "./assets")

	router.Get("/", func(ctx *gin.Context){
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run(":8081")
}
