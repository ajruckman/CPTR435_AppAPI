package web

import "github.com/gin-gonic/gin"

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
