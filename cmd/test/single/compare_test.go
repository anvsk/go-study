package single

import (
	"fmt"
	"testing"
)

type People interface {
	name(name string)
	age(age int)
}

type Student struct {
	score int
}

type Teacher struct {
	score int
}

func (s *Student) name(nm string) {
	fmt.Println("Student name is %s ", nm)
}

func (s *Student) age(age int) {
	fmt.Println("Student age is ", age)
}

func (t *Teacher) name(nm string) {
	fmt.Println("Teacher name is  ", nm)
}

func (t *Teacher) age(age int) {
	fmt.Println("Teacher age is ", age)
}

// 定义interface变量
var conn People

func GetStudentConn() People {
	conn = &Student{score: 100}
	return conn
}

func GetTeacherConn() People {
	conn = &Teacher{score: 100}
	return conn
}

func Compare() {
	stuConn := GetStudentConn()
	// stuConn.age(100)

	stuConn2 := GetStudentConn()
	// stuConn2.age(100)

	teaConn := GetTeacherConn()
	teaConn.name("tea")

	fmt.Println(stuConn == teaConn)
	fmt.Println(stuConn == stuConn2)
}

func TestAaa(t *testing.T) {
	// fmt.Printf("aaa")
}
