package test

import (
    "fmt"
    "reflect"

    uuid "github.com/satori/go.uuid"
)

func Flect() {
    // u := User{1, "zs", 20}
    // Poni(u)

    // 创建
    u1 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u1)

    // 解析
    u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
    if err != nil {
        fmt.Printf("Something gone wrong: %s", err)
        return
    }
    fmt.Printf("Successfully parsed: %s", u2)
}

// 定义结构体
type User struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 绑方法
func (u User) Hello() {
    fmt.Println("Hello")
}

// 传入interface{}
func Poni(o interface{}) {
    t := reflect.TypeOf(o)
    fmt.Println("类型：", t)
    fmt.Println("字符串类型：", t.Name())
    // 获取值
    v := reflect.ValueOf(o)
    fmt.Println(v)
    // 可以获取所有属性
    // 获取结构体字段个数：t.NumField()
    for i := 0; i < t.NumField(); i++ {
        // 取每个字段
        f := t.Field(i)
        fmt.Printf("%s : %v : %s", f.Name, f.Type, f.Tag)
        // 获取字段的值信息
        // Interface()：获取字段对应的值
        val := v.Field(i).Interface()
        fmt.Println("\nval :", val)
    }
    fmt.Println("=================方法====================")
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)
        fmt.Println(m.Name)
        fmt.Println(m.Type)
    }

}
