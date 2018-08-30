package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("get。。。。。。。。。")
}

func (this *UserController) GetInfo() {
	id := this.Ctx.Input.Param("id") //格式（:变量名 要跟路由里面设置的匹配  /user/?id 如果是/user/?:id,前面应该也是 :id）
	this.Ctx.WriteString("getInfo data,id=" + id)
}

func (this *UserController) GetUserName() {
	name := this.Ctx.Input.Param(":name")
	this.Ctx.WriteString("getInfo data,name=" + name)
}

func (this *UserController) GetFileType() {
	//filetype  :=this.Ctx.Input.Param(":ext")
	//filename :=this.Ctx.Input.Param(":path")
	splat := this.Ctx.Input.Param(":splat")
	//this.Ctx.WriteString("filename="+filename)
	//this.Ctx.WriteString("-----filetype="+filetype)
	this.Ctx.WriteString("------all=" + splat)
}
