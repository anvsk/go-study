package util

import (
    "time"
)

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

// GetBeforeDate 获取n天前的时间
func GetDateFromNow(n int) time.Time {
    timer, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
    if n == 0 {
        return timer
    }
    return timer.AddDate(0, 0, n)
}

// 获取本周星期N的Date
func WeekDate(n int) string {
    if n == 0 {
        n = 7
    }
    now := time.Now()
    offset := (n - int(now.Weekday()))
    if offset > (n - 1) { //判断周日
        offset = (n - 7)
    }
    return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).
        AddDate(0, 0, offset).Format("2006-01-02")
}
