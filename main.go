package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {

	app := gin.Default()

	// routes.
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pong",
		})
	  })

	app.Run(":8090")
}