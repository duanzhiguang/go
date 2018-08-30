package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id   int
	Name string `orm:size(100)`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beegoDb?charset=utf8", 30)
	//注册模块
	orm.RegisterModel(new(User))
	//创建table
	orm.RunSyncdb("default", false, true)
}

//插入用户信息
func InsertUser() {
	o := orm.NewOrm()
	user := User{Name: "dzg"}

	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}
func UpdateUser() {
	o := orm.NewOrm()
	user := User{Name: "dzgggg", Id: 1}
	up, err := o.Update(&user)
	fmt.Printf("ID: %d, ERR: %v\n", up, err)

}
