package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-auth/config"
	"github.com/go-auth/controllers"
)

func main() {
    config.ConnectDB()

    r := gin.Default()

    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/register", controllers.Register)
        authRoutes.POST("/login", controllers.Login)
    }

    r.Run(":8080") // Runs on port 8080
}