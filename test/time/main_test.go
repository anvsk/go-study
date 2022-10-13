package time2

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {

	tt, err := time.Parse("20060102", "20221011")
	if err != nil {
		panic(err)
	}
	fmt.Println(tt.Format("20060102"))

	days := CalHowMuchDaysInYear(tt)

	fmt.Printf("是今年的第 %d 天!\n", days)
}

// 计算当前是一年中的第几天
func CalHowMuchDaysInYear(t time.Time) (days int) {
	y := t.Year()
	m := t.Month()
	d := t.Day()
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	return days
}

// 计算入职日期算起的年假天数和日期
func TestJIaqi(t *testing.T) {
	// 入职日期
	indate := "20220104"
	in, _ := time.Parse("20060302", indate)

	// 入职天数
	diffDays := CalInDutyDays(in)

	// c天一次假
	c := float64(365 / 5)

	var jia int64 = 0

	times := []string{}

	// 算出有几天年假和下一次年假日期
	for diffDays > 0 {
		if diffDays > c {
			jia++
			times = append(times, in.Add(time.Duration(24*float64(jia)*c)*time.Hour).Format("2006-01-02"))
		}
		diffDays = diffDays - c
	}
	fmt.Println("可请假天数=", jia, "日期=", times)
	for 5-jia > 0 {
		fmt.Println("下次请假日期", in.Add(time.Duration(24*float64(jia+1)*c)*time.Hour).Format("2006-01-02"))
		jia++
	}
}

// 计算在职多少天
func CalInDutyDays(t time.Time) float64 {
	h := time.Since(t).Hours() / 24
	return h
}
