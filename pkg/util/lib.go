package util

import (
    "encoding/json"
    "reflect"
    "time"
)

func InArray(need interface{}, haystack interface{}) (exists bool, index int) {
    exists = false
    index = -1
    switch reflect.TypeOf(haystack).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(haystack)
        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(need, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }
    return
}

// 比较时间
func CompareTime(time1, time2 string) (flag bool, err error) {
    if time1 == time2 {
        flag = true
        return
    }
    t1, err := time.Parse("15:04", time1)
    if err != nil {
        return
    }
    t2, err := time.Parse("15:04", time2)
    if err != nil {
        return
    }
    if t1.Before(t2) {
        return true, nil
    }
    return
}

// 类型互转，通过json编解码
func Type2type(from interface{}, to interface{}) {
    bytes, err := json.Marshal(from)
    if err != nil {
        return
    }
    json.Unmarshal(bytes, to)
}
