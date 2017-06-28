package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"encoding/gob"
	"blog/models"
)

func main() {
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@tcp(localhost:3306)/blog?charset=utf8",50)

	gob.Register(map[string]models.User{})
}