package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	test()
	r.Run("127.0.0.1:5000")
}
