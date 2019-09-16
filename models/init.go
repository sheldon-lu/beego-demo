package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


var (
	o orm.Ormer
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println(dsn)

	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(RbacAccess), new(RbacApi), new(RbacRole), new(User))
	o = orm.NewOrm()
}

// TableName 统一表名管理
func TableName(tableName string) string {
	return beego.AppConfig.String("db.prefix") + tableName
}

type M struct {
}

func (this *M) Object() orm.Ormer {
	return o
}

