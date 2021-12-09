// 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中
package defer_test

import "testing"

func TestArguDefer(*testing.T) {
	println(f())
}

func f() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}
