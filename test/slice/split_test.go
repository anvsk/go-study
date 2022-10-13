package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s      string
		symbol string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.symbol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit2(t *testing.T) {
	type args struct {
		s      string
		symbol string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split2(tt.args.s, tt.args.symbol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXxx(t *testing.T) {
	// a := []byte{2, 3}
	// r := byte(239)
	// a[0] &= r
	// a[1] |= r
	// a[1] = a[1] | r

	// fmt.Println(a)
	for i := 0; i < 20; i++ {
		r := byte(2)
		rr := byte(i)
		fmt.Printf("%v&%v=%v\r", rr, r, rr&r)
		fmt.Printf("%v|%v=%v\r", rr, r, rr|r)
	}
	return
}
