package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// سرو کردن فایل‌های ویدئو (HLS)
	r.StaticFS("/videos", http.Dir("./videos"))

	r.Run(":8080")
}