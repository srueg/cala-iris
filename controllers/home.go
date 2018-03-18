package controllers

import (
	"github.com/kataras/iris/mvc"
)

// HomeController to control
type HomeController struct {
}

// Get renders home
func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home.hbs",
		Data: map[string]interface{}{
			"title": "Calabouxi",
		},
	}
}
