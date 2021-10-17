package routers

import (
	"github.com/astaxie/beego"
	"new_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "get:Get;post:Post;put:Put;delete:Delete")

}
