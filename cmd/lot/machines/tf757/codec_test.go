package tf757

import (
	"bytes"
	"fmt"
	"testing"
)

// 校验位编码
func TestPrePacketByStdCRC(t *testing.T) {
	res := PrePacketByStdCRC([]byte{1, 2, 3, 4, 5}, 2, 5)
	fmt.Println(res)
	if bytes.Compare(res, []byte{67, 3}) != 0 {
		t.Fatal("Method Error！")
	}
	res2 := PrePacketByStdCRC([]byte{1, 2, 3, 4, 5}, 3)
	fmt.Println(res2)
	if bytes.Compare(res2, []byte{195, 115}) != 0 {
		t.Fatal("Method Error！")
	}
	res3 := PrePacketByStdCRC([]byte{1, 2, 3, 4, 5})
	fmt.Println(res3)
	if bytes.Compare(res3, []byte{42, 187}) != 0 {
		t.Fatal("Method Error！")
	}
	// 站号2的实时状态请求包
	res4 := PrePacketByStdCRC([]byte{2, 3, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(res4)
	if bytes.Compare(res3, []byte{169, 252}) != 0 {
		t.Fatal("Method Error！")
	}
}

// 测试跟C#输出一致
func TestPrePacketExplainRule3(t *testing.T) {
	res := PrePacketExplainRule3("工艺名一", 0)
	fmt.Println(res)
	if bytes.Compare(res[:12], []byte{229, 183, 165, 232, 137, 186, 229, 144, 141, 228, 184, 128}) != 0 {
		t.Fatal("Method Error！")
	}
}
