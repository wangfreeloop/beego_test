package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "hello/routers"
		"fmt"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
