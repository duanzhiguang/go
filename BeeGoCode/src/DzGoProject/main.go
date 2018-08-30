package main

import (
	_ "DzGoProject/models"
	_ "DzGoProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//models.InsertUser()
	//
	//models.UpdateUser()
	beego.Run()

}
