package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/updatefile", func(c *gin.Context) {
		file, err := c.FormFile("fileNm")
		if err != nil {
			c.String(http.StatusBadRequest, "file update error")
		}
		path, _ := os.Getwd()
		c.SaveUploadedFile(file, path+"/"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s update success", file.Filename))
	})
	r.Run(":8080")

	//Content-Type multipart/form-data
}
