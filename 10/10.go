package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(login())
	r.GET("/login", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, "login success"+user)
	})
	r.Run(":8080")

}

func login() gin.HandlerFunc {
	// user and password
	account := gin.Accounts{
		"admin": "123",
		"user":  "456",
	}
	auth := gin.BasicAuth(account)
	return auth
}
