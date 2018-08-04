package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "hello/routers"
	"log"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		log.Fatal("Fail to create xorm db engine : %v\n", err)
	}
	beego.Run()
}
