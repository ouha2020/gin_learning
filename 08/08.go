package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/updatefiles", func(c *gin.Context) {

		mForm, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "file update error")
		}
		path, _ := os.Getwd()
		files := mForm.File["file_key"]

		for _, file := range files {
			c.SaveUploadedFile(file, path+"/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d update success", len(files)))
	})
	r.Run(":8080")

	//Content-Type multipart/form-data
}
