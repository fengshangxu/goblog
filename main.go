package main

import (
	"github.com/astaxie/beego"
	_ "goblog/routers"
)

func main() {
	funcConf()
	beego.Run()
}

func funcConf() {
	beego.AddFuncMap("appName", func() string { return "123" })
}
