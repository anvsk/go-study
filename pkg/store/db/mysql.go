package db

import (
	"database/sql"
	"fmt"
	"go-ticket/pkg/util"
	llog "log"
	"os"
	"sync"
	"time"

	"github.com/pieterclaerhout/go-log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

type mysqlDB struct {
    db       *gorm.DB
    onceLock sync.Once
}

var orm mysqlDB

func (d *mysqlDB) getInstance() *gorm.DB {
    d.onceLock.Do(func() {
        cf := util.Config.Store.DB.Connects["mysql"]
        // dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cf.User, cf.Password, cf.Host, cf.Port, cf.DbName)
        // dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cf.User, cf.Password, cf.Host, cf.Port, cf.DbName)
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cf.User, cf.Password, cf.Host, cf.Port, cf.DbName)
        sqlDB, _ := sql.Open("mysql", dsn)
        sqlDB.SetMaxIdleConns(10)
        sqlDB.SetMaxOpenConns(100)
        sqlDB.SetConnMaxLifetime(time.Minute)
        logLev := logger.Silent
        if util.Config.Debug {
            logLev = logger.Info
        }
        var err error
        d.db, err = gorm.Open(mysql.New(mysql.Config{
            Conn: sqlDB,
        }), &gorm.Config{
            Logger: logger.New(
                llog.New(os.Stdout, "\r\n", llog.LstdFlags), // io writer
                logger.Config{
                    SlowThreshold:             100 * time.Millisecond, // Slow SQL threshold
                    LogLevel:                  logLev,                 // Log level
                    IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
                    Colorful:                  true,                   // Disable color
                },
            ),
        })
        if err != nil {
            log.ErrorDump(err, "mysql init error")
        }
        d.db.Use(prometheus.New(prometheus.Config{
            DBName:          "DB_" + cf.DbName, // 使用 `DBName` 作为指标 label
            RefreshInterval: 15,                // 指标刷新频率（默认为 15 秒）
            PushAddr:        "",                // 如果配置了 `PushAddr`，则推送指标
            StartServer:     false,             // 启用一个 http 服务来暴露指标
            HTTPServerPort:  9002,              // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
            MetricsCollector: []prometheus.MetricsCollector{
                &prometheus.MySQL{
                    Prefix:        "gorm_status_",
                    Interval:      100,
                    VariableNames: []string{"Threads_running"},
                },
            },  // 用户自定义指标
        }))
        log.Debug("init Mysql success")
    })
    return d.db
}

func TestMysql() {
    Orm.AutoMigrate(&TestUser{})
    // Set table options
    // 插入
    Orm.Create(&TestUser{
        10,
        "andy",
        12,
    })
    user := TestUser{}
    // 查询
    Orm.Find(&user, "id = ?", 10)
    log.DebugDump(user, "dumpMysqlUser")
}
