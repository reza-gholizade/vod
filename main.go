package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Serve video files (HLS)
	r.StaticFS("/videos", http.Dir("./videos"))

	// Serve the index.html file at the root route
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.Run(":8080")
}
