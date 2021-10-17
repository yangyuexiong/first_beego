package main

import (
	"github.com/astaxie/beego"
	_ "new_beego/routers"
)

// go build -mod=mod

func main() {
	beego.Run()
}
