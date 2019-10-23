package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

// @router / [get]
func (c *IndexController) Index() {
	c.Ctx.WriteString("API")
	return
}