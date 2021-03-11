package admin

import "fmt"

type LoginController struct {
	BaseController
}

func (login *LoginController) Index() {
	fmt.Print(111)
	login.TplName = "admin/user/login.html"
}
