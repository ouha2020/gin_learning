package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/service", func(c *gin.Context) {
		url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
		response, err := http.Get(url)

		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		body := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
	})

	r.Run(":8080")
}
