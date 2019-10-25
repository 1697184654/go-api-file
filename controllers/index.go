package controllers

import (
	"github.com/astaxie/beego"
	"go-api-file/requests"
	"go-api-file/responses"
)

type IndexController struct {
	beego.Controller
}

// @router / [get]
func (c *IndexController) Index() {
	c.Ctx.WriteString("API")
	return
}

// @Param user_name
// @Param email
// @Param tel
// @router /sign-up [post]
func (c *IndexController) SignUp() {
	userInfo := requests.UserInfo{}
	userInfo.UserName = c.GetString("user_name")
	userInfo.Email = c.GetString("email")
	userInfo.Tel = c.GetString("tel")
	signUpResponse := responses.IndexSignUp{}
	c.Data["json"] = signUpResponse.Process(userInfo)
	c.ServeJSON()
}

// @router /visit-count [get]
func (c *IndexController) VisitCount() {

	visitCountResponse := responses.IndexVisitCount{}
	c.Data["json"] = visitCountResponse.Process()
	c.ServeJSON()
}
