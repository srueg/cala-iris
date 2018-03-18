package controllers

import (
	"github.com/kataras/iris/mvc"
)

// ProfileController to control
type ProfileController struct {
}

// User represents a user
type User struct {
	Name string
}

// Get renders profile
func (c *ProfileController) Get() mvc.Result {
	user := User{Name: "test"}

	return mvc.View{
		Name: "profile.hbs",
		Data: map[string]interface{}{
			"title": "Profile of " + user.Name,
			"user":  user,
		},
	}
}
