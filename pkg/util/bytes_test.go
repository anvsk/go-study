package util

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestByteArrayToAsciiString(*testing.T) {
	buf := []byte{168, 100, 100, 100}
	// res := ByteArrayToAsciiString(buf, 0, 4)
	// fmt.Println(res)
	fmt.Println(BytesToString(buf))
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
