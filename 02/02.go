package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/rdt1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://google.com")
	})

	r.GET("/rdt2", func(c *gin.Context) {
		c.Request.URL.Path = "/rdt"
		r.HandleContext(c)
	})

	r.GET("/rdt", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "return ok",
		})
	})

	r.Run(":8080")
}
