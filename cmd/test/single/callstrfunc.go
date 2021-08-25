package single

import (
	"log"
	"reflect"
)

type Animal struct {
	name string
}

type Num struct {
	num1 int
	nm2  int
}

func (a *Animal) Jiao() Num {
	log.Println(" hellow! -- with method ", a.name)
	return Num{
		1,
		23,
	}
}

func Jiao() {
	log.Println(" hellow! -- with func")
}

func Callstrfunc() {
	str := "Jiao"
	res := reflect.ValueOf(&Animal{name: "222"}).MethodByName(str).Call([]reflect.Value{})
	log.Println(reflect.ValueOf(res))
	log.Println(reflect.TypeOf(res))
}
