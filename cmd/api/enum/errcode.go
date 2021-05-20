// 定义返回错误码的常量
package enum

const (
    _   int = 2000 + iota
    TOKEN_MISSING
    TOKEN_EXPIRE
)

var Emap = map[int]string{
    TOKEN_MISSING: "TOKEN_MISSING",
    TOKEN_EXPIRE:  "TOKEN_EXPIRE",
}
