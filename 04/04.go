package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type groupInfo struct {
	Data string
	Path string
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		r1 := v1.Group("/user")
		r1.GET("/login", login) // http://127.0.0.1:8080/v1/user/login
		r2 := r1.Group("add")
		r2.GET("/info", info) // http://127.0.0.1:8080/v1/user/add/info
	}

	v2 := r.Group("/v2")
	{
		r2 := v2.Group("update")
		r2.GET("/info", info) // http://127.0.0.1:8080/v2/update/info
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, groupInfo{"login", c.Request.URL.Path})
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK, groupInfo{"info", c.Request.URL.Path})
}
