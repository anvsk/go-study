package slice

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cases struct {
	in       string
	sep      string
	expected []string
}

var ca = []cases{
	{
		",,lda,,,jskl,dj,,,",
		",",
		[]string{"lda", "jskl", "dj"},
	},
	{
		",,lda,,,j",
		",",
		[]string{"lda", "j"},
	},
	// {
	// 	",,la,,,j",
	// 	",",
	// 	[]string{"la", "j"},
	// },
	// {
	// 	",,,sad,,,32894,dsjf,,,",
	// 	",,",
	// 	[]string{",sad", ",32894,dsjf", ","},
	// },
}

func TestSplitFunc(t *testing.T) {
	{
		s := ",,lda,,,jskl,dj,,,"
		res := Split(s, ",")
		want := []string{"lda", "jskl", "dj"}
		flag := reflect.DeepEqual(res, want)
		fmt.Println(flag)
		if !flag {
			t.Errorf("expected %v,got %v", want, res)
		}
	}
	{
		s := ",,,sad,,,32894,dsjf,,,"
		res := Split2(s, ",,")
		want := []string{",sad", ",32894,dsjf", ","}
		flag := reflect.DeepEqual(res, want)
		fmt.Println(flag)
		if !flag {
			t.Errorf("expected %v,got %v", want, res)
		}
	}

}

func TestSplitFuncByTable(t *testing.T) {

	for _, v := range ca {
		v := v
		t.Run("////Source:"+v.in, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, v.expected, Split(v.in, v.sep))
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for _, v := range ca {
		b.Run(fmt.Sprintf(v.in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Split(v.in, v.sep)
			}
		})
	}
}

func BenchmarkSplit2(b *testing.B) {
	for _, v := range ca {
		b.Run("2==="+fmt.Sprintf(v.in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Split2(v.in, v.sep)
			}
		})
	}
}
