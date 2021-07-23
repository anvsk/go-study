package leetcode

import "strconv"

func Stradd(s1, s2 string) string {
    length := len(s1)
    if len(s2) > len(s1) {
        length = len(s2)
    }
    curjinyi := 0
    res := ""
    for i := length; i > 0; i-- {
        var nextJinyi, up, down int
        if length > len(s1) {
            up = 0
        } else {
            up, _ = strconv.Atoi(string(s1[(i - 1)]))
        }
        if length > len(s2) {
            down = 0
        } else {
            down, _ = strconv.Atoi(string(s2[(i - 1)]))
        }
        tmpres := up + down + curjinyi
        if tmpres > 9 {
            nextJinyi = 1
            tmpres -= 10
        }
        curjinyi = nextJinyi
        res = strconv.Itoa(tmpres) + res
    }
    if curjinyi == 1 {
        res = "1" + res
    }
    return res
}
