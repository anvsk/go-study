package db

import (
	"fmt"
	"go-study/pkg/util"
	"sync"

	"github.com/pieterclaerhout/go-log"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var chdb ClickHouseDB

type ClickHouseDB struct {
	db       *gorm.DB
	onceLock sync.Once
}

func (d *ClickHouseDB) getInstance() *gorm.DB {
	d.onceLock.Do(func() {
		cf := util.Config.Store.DB.Connects["clickhouse"]
		var err error
		dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", cf.Host, cf.Port, cf.DbName, cf.User, cf.Password)
		chdb.db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
		if err != nil {
			log.ErrorDump(err, "clickhouse connect error")
		}
		log.Debug("init ClickHouse success")
	})
	return d.db
}

type TestUser struct {
	ID   int64
	Name string
	Age  int64
}

func TestCH() {
	// Auto Migrate
	ClickHouse.AutoMigrate(&TestUser{})
	// Set table options
	// 插入
	ClickHouse.Create(&TestUser{
		10,
		"andy",
		12,
	})
	user := TestUser{}
	// 查询
	ClickHouse.Find(&user, "id = ?", 10)
	log.DebugDump(user, "dumpClickhouseUser")
}
