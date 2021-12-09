package time

import (
	"go-ticket/pkg/store/db"
	"time"
)

func TestMain() {
	// println("aa")
	db.InitDB()
	time.Now()
	// db.Orm.Table("")
}
