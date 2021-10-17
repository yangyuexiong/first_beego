package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"new_beego/tools"
)

type UserModel struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	// 获取url参数
	UserId := c.GetString("id")

	// 返回字符串
	c.Ctx.WriteString("User:" + UserId)
}

func (c *UserController) Post() {

	// 接收json
	var user UserModel
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	// 返回json
	c.Data["json"] = tools.ResultApi(200, "操作成功", user)
	c.ServeJSON()
}

func (c *UserController) Put() {

	// form data
	username := c.GetString("username")
	password := c.GetString("password")
	var user1 UserModel
	if err := c.ParseForm(&user1); err != nil {
		c.Data["json"] = tools.ResultApi(400, "参数异常", nil)
		c.ServeJSON()
	}

	// 返回json
	var userJson = UserModel{
		Username: username,
		Password: password,
	}
	c.Data["json"] = userJson
	c.ServeJSON()
}

func (c *UserController) Delete() {
	// 返回json
	c.Data["json"] = tools.ResultApi(200, "操作成功", nil)
	c.ServeJSON()
}
