package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/file", func(c *gin.Context) {
		path, _ := os.Getwd()
		fileNm := path + "/" + c.Query("name")
		c.File(fileNm)
	})

	r.Run(":8080")
}
