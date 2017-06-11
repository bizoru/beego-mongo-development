package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@hotmail.com"
	c.Data["Name"] = "Steven Sierra"
	c.TplName = "index.tpl"
	// This is a test
}
