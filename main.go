package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/rxps")
	{
		api.GET("/timezone", tzHandler)
	}
	router.Run()
}

func tzHandler(ctx *gin.Context) {
	log.Printf("Context headers %#v\n", ctx.Request.Header)
	log.Printf("Context path params: %#v\n", ctx.Params)
	log.Printf("Context query params: %#v\n", ctx.Request.URL.Query())
	ctx.JSON(200, gin.H{
		"ping": "pong",
	})
}
