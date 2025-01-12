package main

import (
    "github.com/gin-gonic/gin"
    "mock-ses/handlers"
)

func main() {
    r := gin.Default()
    r.POST("/send-email", handlers.SendEmail)
    r.GET("/statistics", handlers.GetStatistics)
    r.Run(":8080")
}