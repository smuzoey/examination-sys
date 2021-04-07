package main

import (
	_ "examination-sys/internal/routers"
	"github.com/astaxie/beego"
)

func main() {
	//web.BConfig.Listen.ServerTimeOut=0
	beego.Run()
}
