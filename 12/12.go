package main

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Addr struct {
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UserInfo struct {
	Id      string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"ckNm"`
	Age     uint8  `json:"age" validate:"max=120"`
	Address []Addr `json:"address" validate:"dive"`
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
	r.POST("/validatedive", func(c *gin.Context) {
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
