package main

import (
    "virtbro/pkg/db"
    "github.com/gin-gonic/gin"
)

func main() {
    db.InitDB()
    router := gin.Default()

    router.GET("/connections", func(c *gin.Context) {
        var connections []db.Connection
        db.DB.Find(&connections)
        c.JSON(200, connections)
    })

    router.POST("/connections", func(c *gin.Context) {
        var connection db.Connection
        if err := c.ShouldBindJSON(&connection); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        db.DB.Create(&connection)
        c.JSON(200, connection)
    })

    router.Run(":8080")
}
