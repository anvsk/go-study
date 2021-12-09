package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

//isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool) (int, error) {
	if isSymbol {
		return bytesToIntS(b)
	}
	return bytesToIntU(b)
}

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//整形转换成字节
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

// public static string ByteArrayToAsciiString(byte[] bytes, int StartPos,int length)
//         {
//             string asciiString = String.Empty;
//             if (bytes != null)
//             {
//                 StringBuilder strB = new StringBuilder();
//                 for (int i = 0; i < length; i++)
//                 {

//                     if (bytes[StartPos + i] > 1 && bytes[StartPos + i] < 127)//是ASCII码，直接转
//                     {
//                         strB.Append(((char)bytes[StartPos + i]).ToString());
//                     }
//                     else
//                     {
//                         byte[] tmp = new byte[2];
//                         tmp[0] = bytes[StartPos + i];
//                         tmp[1] = bytes[StartPos + i + 1];
//                         strB.Append(Encoding.Default.GetString(tmp, 0, 2));
//                         i++;
//                     }
//                 }
//                 asciiString = strB.ToString().Trim() ;
//             }
//             return asciiString;
//         }

func ByteArrayToAsciiString(bytes []byte, StartPos, length int) string {
	// if bytes == nil || len(bytes) == 0 {
	// 	return ""
	// }
	// res := ""
	// for i := 0; i < length; i++ {
	// 	//是ASCII码，直接转
	// 	if bytes[StartPos+i] > 1 && bytes[StartPos+i] < 127 {
	// 		res += string(bytes[StartPos+i])
	// 	} else {
	// 		res += string([]byte{bytes[StartPos+i], bytes[StartPos+i+1]})
	// 		i++
	// 	}
	// }
	// return strings.Trim(res, " ")
	return TransferBytes2AsciiCodeString(bytes[StartPos : StartPos+length])
}

func TransferBytes2AsciiCodeString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
