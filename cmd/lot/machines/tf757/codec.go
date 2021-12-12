package tf757

import (
	"encoding/hex"
	"math"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util/convert"
)

// 天富独有--编码规则

// 可变参数
// param1:字节
// param2:index起始位
// param3:index结束位
func PrePacketByStdCRC(args ...interface{}) []byte {
	bytes := args[0].([]byte)
	// 默认值
	var (
		start  = 0
		end    = len(bytes) - 1
		bytes2 = []byte{}
	)
	// 有传参则覆盖
	switch len(args) {
	case 2:
		start = args[1].(int)
	case 3:
		start = args[1].(int)
		end = args[2].(int)
		if end > len(bytes)-1 {
			end = len(bytes) - 1
		}
	}
	if end >= start {
		bytes2 = bytes[start : end+1]
	}
	a := ToCRC16(bytes2, false)
	high, low := ToHighLowByte(a)
	return []byte{high, low}
}

func ToCRC16(bytes []byte, incCRC bool) int {
	num := len(bytes)
	if incCRC {
		num -= 2
	}
	num2 := 65535
	for i := 0; i < num; i++ {
		num2 = num2 ^ int(bytes[i])
		for j := 1; j <= 8; j++ {
			num3 := num2 & 1
			num2 >>= 1
			if num3 > 0 {
				num2 ^= 0xA001
			}
		}
	}
	num4 := num2 >> 8
	num2 = (num2 << 8) | num4
	num2 = num2 & 0xFFFF
	return num2
}

func ToHighLowByte(a int) (high, low byte) {
	high = (byte)(a >> 8)
	low = (byte)(a << 8 >> 8)
	return
}

func PrePacketExplainRule5(value interface{}) []byte {
	doubleValue := convert.ToFloat(value)
	bytes2 := util.IntToBytesLittle(int(doubleValue * 10))
	return []byte{bytes2[1], bytes2[0]}
}

func PrePacketExplainRule6(value interface{}) []byte {
	doubleValue := convert.ToFloat(value)
	bytes2 := util.IntToBytesLittle(int(doubleValue))
	return []byte{bytes2[1], bytes2[0]}
}

func PrePacketExplainRule3(value string, iLimitSize int) []byte {
	if iLimitSize == 0 {
		iLimitSize = 48
	}
	vbyte := []byte(value)
	res := make([]byte, iLimitSize)
	copy(res[:], vbyte)
	return res[:]
}

func PrePacketExplainRule4(value string) []byte {
	res, err := hex.DecodeString(value)
	if err != nil {
		return []byte{}
	}
	return res
}

func PrePacketExplainRule1(value int) []byte {
	return []byte{byte(value)}
}

func UnPackByExplainRule1(bytes []byte, iStartPos, iEndPos int) float64 {
	return float64(int(bytes[iStartPos])*256 + int(bytes[iEndPos]))
}

func UnPackByExplainRule3(bytes []byte, iStartPos, iEndPos int) float64 {
	return float64(int(bytes[iStartPos])*256+int(bytes[iEndPos])) / 10
}

func UnPackByExplainRule18(bytes []byte, iStartPos, iEndPos int) int {
	return int(float64(int(bytes[iStartPos])*256+int(bytes[iEndPos])) / 100)
}

func UnPackByExplainRule22(bytes []byte, iStartPos, iEndPos int) float64 {
	var (
		p1 float64 = 0
		p2 float64 = 0
		p3 float64 = 0
		p4 float64 = 0
		p5 float64 = 0
		p6 float64 = 0
		p7 float64 = 0
		p8 float64 = 0
	)
	if bytes[iStartPos] != 0 {
		p1 = float64(bytes[iStartPos]) + math.Pow(256, 7)
	}
	if bytes[iStartPos+1] != 0 {
		p2 = float64(bytes[iStartPos+1]) + math.Pow(256, 6)
	}
	if bytes[iStartPos+2] != 0 {
		p3 = float64(bytes[iStartPos+2]) + math.Pow(256, 5)
	}
	if bytes[iStartPos+3] != 0 {
		p4 = float64(bytes[iStartPos+3]) + math.Pow(256, 4)
	}
	if bytes[iStartPos+4] != 0 {
		p5 = float64(bytes[iStartPos+4]) + math.Pow(256, 3)
	}
	if bytes[iStartPos+5] != 0 {
		p6 = float64(bytes[iStartPos+5]) + math.Pow(256, 2)
	}
	if bytes[iStartPos+6] != 0 {
		p7 = float64(bytes[iStartPos+6]) + math.Pow(256, 1)
	}
	p8 = float64(bytes[iStartPos+7])
	return (p1 + p2 + p3 + p3 + p4 + p5 + p6 + p7 + p8) / 1000000
}
