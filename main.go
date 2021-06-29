package main

import (
	_ "nomadiclife/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qburst@tcp(localhost:3306)/nomadiclife?charset=utf8")
	orm.RunCommand()
}

func main() {
	orm.Debug = true
	beego.Run()
}
