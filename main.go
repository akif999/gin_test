package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// API の定義
	v1 := engine.Group("/v1")
	v1.GET("/api1", api1Handle)
	v1.GET("/api2", api2Handle)

	engine.Run(":3000")
}

// クエリの値は利用しない API
func api1Handle(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// クエリの値を利用する API
func api2Handle(c *gin.Context) {
	hoge := c.Query("hoge")
	c.JSON(http.StatusCreated, gin.H{
		"status": hoge,
	})
}
