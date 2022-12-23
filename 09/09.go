package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(login())
	r.GET("/login", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"name": name,
			"age":  age,
		})
	})
	r.Run(":8080")

}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ageStr := c.Query("age")
		name := c.Query("name")
		if name != "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "age error")
			return
		}
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "age error")
			return
		}
		if age < 0 || age > 120 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "age error [0-120]")
			return
		}
		c.Next()
	}
}
