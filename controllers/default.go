package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type SeconController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "zilla.org"
	c.Data["Email"] = "zilla@gmail.com"
	c.TplName = "index.tpl"
}

func (c *SeconController) Get() {
	c.Data["name"] = "zilla"
	c.TplName = "second.tpl"
}
