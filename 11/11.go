package main

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserInfo struct {
	Id   string `json:"Id" validate:"required"`
	Name string `json:"name" validate:"ckNm"`
	Age  uint8  `json:"age" validate:"max=120"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("ckNm", ckNmF)
}

func ckNmF(field validator.FieldLevel) bool {
	count := utf8.RuneCountInString(field.Field().String())
	if count >= 2 && count <= 15 {
		return true
	}
	return false
}

func main() {
	r := gin.Default()
	var user UserInfo
	r.POST("/validate", func(c *gin.Context) {
		err := c.Bind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}
		err = validate.Struct(user)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				fmt.Println(e.Field())
				fmt.Println(e.Value())
			}
			c.JSON(http.StatusBadRequest, "error")
			return
		}
		c.JSON(http.StatusOK, "login success")
	})
	r.Run(":8080")

}
