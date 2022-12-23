package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userInfo struct {
	name string
	age  int
}

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>hello,world</b>",
		})
	})

	r.GET("/jsonhtml", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello,world</b>",
		})
	})

	r.GET("/xml", func(c *gin.Context) {
		info := userInfo{
			name: "wb",
			age:  3,
		}
		c.XML(http.StatusOK, info)
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"name": "wb",
			"age":  "3",
		})
	})

	r.Run(":8080")
}
