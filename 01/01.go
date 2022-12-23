package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "ok",
			"data": "welcome" + name,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "ok",
			"data": "welcome" + name,
		})
	})
	r.Run(":8080")

}
