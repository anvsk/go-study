// 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中
package defer_test

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestMutilDefer(*testing.T) {
	defer func() {
		fmt.Println("aa")
	}()
	defer func() {
		panic("bb")
		fmt.Println("bb")
	}()
	defer func() {
		panic("aa")
		fmt.Println("cc")
	}()
}

func TestArguDefer(*testing.T) {
	println(f())
}

func f() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}

func TestRecoverFunc(*testing.T) {
	if err := ffff(); err != nil {
		log.Println("ffff() err:", err)
	}
}

func TestWarpperFunc(*testing.T) {
	newaa := RecoverWarpper(aa)
	// aa()
	newaa()
}

func aa() {
	cc()
}

func cc() {
	dd()
}

func dd() {
	panic("dd")
}

func RecoverWarpper(func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				log.Println(err, "panic", "stack", "...\n"+string(buf))
			}
		}()
		f()
	}
}

func ffff() (err error) {
	defer func() {
		// if err != nil {
		// 	handleError(err)
		// }
		if err := recover(); err != nil {
			// handleError(err)
			err = errors.New(fmt.Sprintln(err))
			return
		}
	}()
	// err = errors.New("aaa")
	// return err
	panic(errors.New("demo-panic"))
}

func handleError(err interface{}) {
	log.Println(err)
	// if err != nil {
	// 	log.Println("common err:", err)
	// 	return
	// }
	if err := recover(); err != nil {
		log.Println("recover err:", err)
		return
	}
}
