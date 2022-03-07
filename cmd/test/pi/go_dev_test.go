package pi

import (
	"fmt"
	"math"
	"testing"
)

func TestOfficialPi(*testing.T) {

	fmt.Println(pi(5000))
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go term2(ch, float64(k))
	}
	f := 0.0
	for k := 0; k < n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

func term2(ch chan float64, v float64) {
	k := int64(1) * n / int64(v+1)
	x := h * (float64(k) + 0.5)
	res := f(x) * h
	ch <- res
}
