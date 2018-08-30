package routers

import (
	"DzGoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/?id", &controllers.UserController{}, "get:GetInfo")
	beego.Router("/user/name/?:name", &controllers.UserController{}, "get:GetUserName")
	beego.Router("/user/*.*", &controllers.UserController{}, "get:GetFileType")

}
