package db

import (
	"go-study/pkg/util"

	"gorm.io/gorm"
)

var Default *gorm.DB
var Orm *gorm.DB
var ClickHouse *gorm.DB

func InitDB() {
	driver := util.Config.Store.DB.Driver
	for _, v := range driver {
		switch v {
		case "mysql":
			Orm = orm.getInstance()
		case "clickhouse":
			ClickHouse = chdb.getInstance()
		}
	}
	Default = orm.getInstance()
}
func GetMysql() *gorm.DB {
	Orm = orm.getInstance()
	return Orm
}
func GetClickHouse() *gorm.DB {
	ClickHouse = chdb.getInstance()
	return ClickHouse
}
