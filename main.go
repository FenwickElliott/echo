package main

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.Writer.Write([]byte(`{"ping":"pong"}`))
	})
	r.GET("/echo", func(c *gin.Context) {
		io.Copy(c.Writer, c.Request.Body)
	})

	log.Fatal(r.Run())
}
