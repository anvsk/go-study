package reflecto

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	// time.Now()
	fmt.Println(1 << 3)
	fmt.Println(2 << 3)
	fmt.Println(30 << 4)
}

func populate(f reflect.Value, v string) {
	kind := f.Kind()
	switch kind {
	case reflect.String:
		f.SetString(v)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
		f.SetBool(b)
	case reflect.Int:
		b, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			panic(err)
		}
		f.SetInt(b)
	case reflect.Float64:
		b, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		f.SetFloat(b)

	}
}
