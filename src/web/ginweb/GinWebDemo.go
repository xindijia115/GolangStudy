package main

import "github.com/gin-gonic/gin"

/**
 * @author: DJ
 * @date: 2024-12-14  14:55
 * @description:
 */

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
