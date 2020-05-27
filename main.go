package main

import (
	"log"
	"math/rand"
	
	"github.com/gin-gonic/gin"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {

		c.Next()

	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		log.Println("----------------",rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run("0.0.0.0:8080")
}
