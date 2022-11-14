package reflecto

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type M interface {
	Say() reflect.Value
}

type Aa struct {
	C int
}

func (*Aa) Say() reflect.Value {
	return reflect.ValueOf(Aa{})
}

type Aa2 struct {
	C int
}

func (*Aa2) Say() reflect.Value {
	return reflect.ValueOf(Aa2{})
}

// 26. 两个 interface 可以比较吗？
// 判断类型是否一样
// reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind()

// 判断两个interface{}是否相等
// reflect.DeepEqual(a, b interface{})

// 将一个interface{}赋值给另一个interface{}
// reflect.ValueOf(a).Elem().Set(reflect.ValueOf(b))

func TestRef(t *testing.T) {
	a := map[string]M{
		"a1": &Aa{},
		"a2": &Aa2{},
	}

	inter := a["a1"]
	inter2 := a["a2"]
	_ = inter2
	// reflect.TypeOf(inter)
	// fmt.Println(reflect.TypeOf(inter).Kind() == reflect.TypeOf(inter2).Kind())

	// var ii interface{}
	// ii = reflect.ValueOf()
	reflect.ValueOf(inter).Elem().Set(reflect.ValueOf(Aa2{C: 3}))
	fmt.Println(reflect.TypeOf(inter).Kind())

	fmt.Printf("inter====%#v", inter)
}

type TestStructMember struct {
	Sli    []int
	Mmap   map[string]string
	Iint   *int
	String *int
}

func TestXxx(t *testing.T) {
	// time.Now()
	fmt.Println(1 << 3)
	fmt.Println(2 << 3)
	fmt.Println(30 << 4)
}

func populate(f reflect.Value, v string) {
	kind := f.Kind()
	switch kind {
	case reflect.String:
		f.SetString(v)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
		f.SetBool(b)
	case reflect.Int:
		b, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			panic(err)
		}
		f.SetInt(b)
	case reflect.Float64:
		b, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		f.SetFloat(b)

	}
}

type Table interface {
	TableName() string
}

type demo111 struct {
	ID int `gorm:"primaryKey"`
	Aa int
	Bb int
	Cc string
}

func (*demo111) TableName() string {
	return "demo111"
}

type demo222 struct {
	ID int `gorm:"primaryKey"`
	Aa int
	Bb int
	Cc string
}

func (*demo222) TableName() string {
	return "demo222"
}

var mmap = map[string]models{}
var db *gorm.DB
var err error

type models struct {
	Name string
	Typ  reflect.Type
	Val  reflect.Value
	Ins  Table
}

func TestXxxModel(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/demo1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&demo111{})
	db.AutoMigrate(&demo222{})
	db.Create(&demo111{})
	db.Create(&demo222{})

	register(&demo111{})
	register(&demo222{})
	call("demo111")
	fmt.Println("<<<<<>>>>>>>")
	call("demo222")
	// reflect.New(mmap["demo111"].Typ)
}

func register(i Table) {
	aa := models{}
	aa.Name = reflect.Indirect(reflect.ValueOf(i)).Type().Name()
	aa.Typ = reflect.TypeOf(i)
	aa.Val = reflect.ValueOf(i)
	aa.Ins = i
	mmap[aa.Name] = aa
}

func call(s string) {
	for _, v := range mmap {
		if v.Name == s {
			_ = v
			mm := v.Ins.TableName()
			m := []map[string]interface{}{}
			errs := db.Table(mm).Scan(&m).Error
			fmt.Println("=====", errs, m)
			return
		}
	}
}
