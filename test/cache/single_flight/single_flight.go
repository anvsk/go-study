// 防止缓存击穿：缓存失效的时候大量打到DB
// 查看标记位状态：可用；直接return
// 				不可用；一个线程去获取数据、获取到后更改标记位
//			   		   其他线程等待、chan通知、超时
package single_flight

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const key string = "pages"

var redisCli redis.Client
var mysqlcCli *gorm.DB

var cacheIns map[string]interface{}

var conctrolLocks map[string]*LockChan

type LockChan struct {
	Done chan struct{}
	sync.Mutex
	Waits map[string]struct{}
}

// var glock sync.Mutex
var glock2 sync.Mutex

// var llllockchan LockChan

// var lock

func InitEnv() {
	var err error
	dsn := "root:@tcp(:3306)/demo1?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlcCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	redisCli := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisCli.Ping().Result()
	if err != nil {
		log.Println(err)
		return
	}

	cacheIns = make(map[string]interface{}, 0)
	conctrolLocks = make(map[string]*LockChan)
	// llllockchan = LockChan{
	// 	make(chan struct{}),
	// 	make(chan struct{}),
	// 	sync.Mutex{},
	// }
}

func Get(i int, key string) interface{} {
	val, ok := cacheIns[key]
	if !ok {
		loc, ok := getLock2(key)
		if ok {
			fmt.Println(i, "getFromDB")
			val = getFromDB(key)
			cacheIns[key] = val
			close(loc.Done)
		} else {
			fmt.Println(i, "wait <-done")
			for {
				select {
				case <-loc.Done:
					val, ok = cacheIns[key]
					if !ok {
						fmt.Println(i, "cache nil")
						return nil
					}
					return val
				case <-time.After(1 * time.Second):
					fmt.Println(i, "timeout 3 sec")
					return nil
				default:
				}
			}

		}
	}
	return val
}

type Test struct {
	ID   int
	Name string
}

func getLock2(key string) (LockChan, bool) {
	glock2.Lock()
	// defer glock2.Unlock()
	lll, ok := conctrolLocks[key]
	if !ok {
		conctrolLocks[key] = &LockChan{
			make(chan struct{}),
			sync.Mutex{},
			make(map[string]struct{}),
		}
		lll = conctrolLocks[key]
	}
	glock2.Unlock()
	return *lll, lll.TryLock()
}

func getFromDB(key string) interface{} {
	var aa Test
	if err := mysqlcCli.Where("id", key).First(&aa).Error; err != nil {
		fmt.Println(err)
	}

	return aa.Name
}
