package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	v1 := engine.Group("/v1")
	v1.GET("/api1", api1Handle)

	engine.Run(":3000")
}

func api1Handle(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
