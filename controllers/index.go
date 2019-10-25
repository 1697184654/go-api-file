package controllers

import (
	"github.com/astaxie/beego"
	"go-api-file/requests"
	"go-api-file/responses"
)

type IndexController struct {
	beego.Controller
}

// @TitLe 首页
// @Description 首页简介
// @router / [get]
func (c *IndexController) Index() {
	c.Ctx.WriteString("API")
	return
}

// @Title 报名
// @Description 用户报名
// @Param user_name
// @Param email
// @Param tel
// @Success 200 {object} responses.IndexSignUp
// @Failure 500 {object} responses.IndexSignUp
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

// @Title 记录用户访问数
// @router /visit-count [get]
func (c *IndexController) VisitCount() {

	visitCountResponse := responses.IndexVisitCount{}
	c.Data["json"] = visitCountResponse.Process()
	c.ServeJSON()
}
