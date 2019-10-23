package controllers

import (
	"github.com/astaxie/beego"
	"go-api-file/services"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @router /signup [get]
func (u *UserController) SignUp() {
	u.Data["json"] = "sign up"
	u.ServeJSON()
}

// @router /signup-count [get]
func (u *UserController) GetSignUpCount() {
	file := services.Files{}
	file.FileName = "/data/www/go_file.txt"
	u.Data["json"] = services.GetVisitCount()
	u.ServeJSON()
}
