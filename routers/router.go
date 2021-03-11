package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
	"goblog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/index", &admin.LoginController{}, "GET:Index")
}
