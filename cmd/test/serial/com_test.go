// 串口通信
package com_serial

import (
	"bytes"
	"fmt"
	"go-ticket/pkg/util"
	"io"
	"log"
	"testing"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

// func TestBuf(t *testing.T) {
// 	a := []byte{0, 1, 2, 3}
// 	// b := []byte{0, 1, 2}
// 	a = a[:2]
// 	a = append(a, make([]byte, 9)...)
// 	fmt.Println(a)
// 	println(len(a))
// 	// fmt.Println(bytes.Equal(a[:3], b))
// }

var MinPackageLength = 8

var MaxPackageLength = 62

type HxDevice struct {
	Conn                      io.ReadWriteCloser
	Buffer                    *HxBuffer
	DeviceStatusResponseCache []byte
	DevHopeStatusCache        HxHopeDevStatus
}

type HxHopeDevStatus struct {
	HopeDeviceStatus int       //设备期望状态
	FlagTime         time.Time //标记时间
	HandleFlag       bool      //处理控制命令的下发情况
}

func (r HxHopeDevStatus) IsOverTime() bool {
	return time.Since(r.FlagTime) >= 15*time.Second
}

func (r HxHopeDevStatus) ReSet() {
	r.HandleFlag = false
}

func NewSerialPort() *HxDevice {
	return &HxDevice{}
}

func TestConnect(*testing.T) {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "/dev/tty.usbserial-BCDFf103Y23",
		BaudRate:        38400,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
		ParityMode:      0,
		// RTSCTSFlowControl: true,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	fmt.Println(options.PortName, "	serial opened。。。")
	// Make sure to clo,se it later.
	// defer port.Close()
	dev := &HxDevice{
		Conn:   port,
		Buffer: NewHxBuffer(),
	}
	go dev.registRead()
	go dev.registIntervalReqStatus()
	// go func() {
	// 	buf := make([]byte, 1000)
	// 	for {
	// 		n, err := port.Read(buf[:])
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Println(n, "]recv:", buf[:n])
	// 	}
	// }()
	for {
	}
}

// 仅read append
func (h *HxDevice) registRead() {
	// buf := make([]byte, 100000)
	for {
		// time.Sleep(50 * time.Millisecond)
		n, err := h.Conn.Read(h.Buffer.Buffer[h.Buffer.ValidLength:])
		if err != nil {
			panic(err)
		}
		fmt.Println("接收到字节数", n)
		// n, err := h.Conn.Read(buf[:])
		h.Buffer.ValidLength += n
		fmt.Println("ValidLength", h.Buffer.ValidLength)

		fmt.Println("Received:", h.Buffer.Buffer[:h.Buffer.ValidLength])

		if err != nil {
			h.Close()
			println("hx设备断开连接")
		}
		if h.Buffer.ValidLength < 8 {
			continue
		}
		// fmt.Println("Before:AfterReceivePortBytes接收:", n, h.Buffer.Buffer[h.Buffer.ValidLength:h.Buffer.ValidLength+n])

		h.AfterReceivePortBytes()
	}
}

func (h *HxDevice) registIntervalReqStatus() {
	for {
		b := []byte{0x05, 0x44, 0x43, 0x52, 0x5A, 0x31, 0x36, 0x39, 0x46, 0x04}
		n, err := h.Conn.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
		fmt.Println("Wrote", n, "bytes:", b)
		time.Sleep(1000 * time.Millisecond)
	}
}

// 解析-分发handler
func (h *HxDevice) AfterReceivePortBytes() {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Println("Hx AfterReceiveSocketBytes error:", err)
		}
	}()
	dealLength := 0
	buffer := h.Buffer
	for i := 0; i < buffer.ValidLength-MinPackageLength; {

		//实时状态包
		if buffer.ValidLength-i >= 62 && bytes.Equal(buffer.Buffer[i:i+5], []byte{6, 68, 67, 82, 90}) && (i+62) <= buffer.ValidLength && buffer.Buffer[i+61] == 3 {
			fmt.Println("接收到实时状态包:", len(buffer.Buffer))
			iCurLength := 62
			// standNo := util.TransferBytes2AsciiCodeString(buffer.Buffer[5:7])
			realWithBuffer := make([]byte, 62)
			copy(realWithBuffer[:], buffer.Buffer[i:i+62])
			// Array.Copy(buffer.Buffer, i, realWithBuffer, 0, 62);

			// 缓存当前状态
			// SaveDeviceStatus(standNo, realWithBuffer);
			h.DeviceStatusResponseCache = realWithBuffer
			state := realWithBuffer[24]

			if h.DevHopeStatusCache.HandleFlag {
				// 	//如果状态等于预期状态，则清空标记，同时也上抛数据
				if h.DevHopeStatusCache.HopeDeviceStatus == int(state) || h.DevHopeStatusCache.IsOverTime() {
					h.DevHopeStatusCache.ReSet()
					// DealDataCommandAck(standNo, realWithBuffer)
				}
				// 	//如果状态不符合预期，则不处理任何事情
			} else {
				// DealDataCommandAck(standNo, realWithBuffer)

			}
			// else
			// {
			// 	DealDataCommandAck(standNo, realWithBuffer);
			// }
			//SetDeviceStatusResponse(standNo);
			i += iCurLength
			dealLength = i
			//工艺上传的启动帧
		} else if buffer.ValidLength-i >= 24 && buffer.Buffer[i] == 6 && buffer.Buffer[i+3] == 82 && buffer.Buffer[i+4] == 80 && (i+24) <= buffer.ValidLength && buffer.Buffer[i+23] == 3 {
			iCurLength := 24
			// int standNo = 0;
			// int craftNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			// Logger.Info($"工艺检验，接收到站号{standNo}的上传启动帧报文");
			// standNo = GetCurrSendCraftStationNo(craftNo.ToString());
			// int count = 7;
			// for (count = 7; count < 18; count++)
			// {
			// 	if (buffer.Buffer[i + count] == 32)
			// 		break;
			// }
			// string craftName = util.ByteArrayToAsciiString(buffer.Buffer, i + 7, count - 6);
			// if (UploadingCraftNodeCache.ContainsKey(standNo))
			// {
			// 	var rptHeaderNode = UploadingCraftNodeCache[standNo].FirstOrDefault(x => x.IsRptHeader == true);
			// 	if (rptHeaderNode != null && rptHeaderNode.VerCraftCode == standNo.ToString().PadLeft(2, '0') && rptHeaderNode.VerCraftName == craftName)
			// 		rptHeaderNode.IsVerify = true;
			// }

			i += iCurLength
			dealLength = i
			// 下载工艺启动帧和结束帧
		} else if buffer.ValidLength-i >= 8 && buffer.Buffer[i] == 0x06 && buffer.Buffer[i+3] == 0x57 && buffer.Buffer[i+4] == 0x50 && buffer.ValidLength >= i+8 && buffer.Buffer[i+7] == 0x03 {
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();

			i += iCurLength
			dealLength = i
			// 下载工艺工序帧响应
		} else if buffer.ValidLength-i >= 8 && buffer.Buffer[i] == 0x06 && buffer.Buffer[i+3] == 0x57 && buffer.Buffer[i+4] == 0x53 && buffer.ValidLength >= i+8 && buffer.Buffer[i+7] == 0x03 {
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			// if (DownCraftNodeResponseEvent.ContainsKey(standNo))
			// {
			// 	DownCraftNodeResponseEvent[standNo].Set();
			// }

			i += iCurLength
			dealLength = i
			//工艺上传的工序帧
		} else if buffer.ValidLength-i >= 24 && buffer.Buffer[i] == 6 && buffer.Buffer[i+1] == 82 && buffer.Buffer[i+2] == 83 && (i+24) <= buffer.ValidLength && buffer.Buffer[i+23] == 3 {
			iCurLength := 24

			// int craftNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 3, 2).To<int>();

			// int standNo = GetCurrSendCraftStationNo(craftNo.ToString());

			// var craftOperationId = util.ByteArrayToAsciiString(buffer.Buffer, i + 5, 2).To<int>();
			// string functionNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 7, 2);
			// string p1 = util.ByteArrayToAsciiString(buffer.Buffer, i + 9, 4).To<int>().ToString();
			// string p2 = util.ByteArrayToAsciiString(buffer.Buffer, i + 13, 4).To<int>().ToString();
			// string p3 = util.ByteArrayToAsciiString(buffer.Buffer, i + 17, 4).To<int>().ToString();
			// Logger.Info($"工艺检验，接收到站号{standNo},工序为{craftOperationId},功能码为{functionNo}，参数为{p1 + p2 + p3}的上传工序帧报文");
			// if (UploadingCraftNodeCache.ContainsKey(standNo))
			// {
			//     var rptNode = UploadingCraftNodeCache[standNo].FirstOrDefault(x => x.VerCraftStepIndex == craftOperationId);
			//     if (rptNode != null && rptNode.VerCraftCode == functionNo && rptNode.VerParameterValue1 == p1 && rptNode.VerParameterValue2 == p2 && rptNode.VerParameterValue3 == p3)
			//         rptNode.IsVerify = true;
			// }
			i += iCurLength
			dealLength = i
			//复位
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WQ" && (i+8 <= buffer.ValidLength) {

			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();

			// //先回应控制指令
			// DealCommandAck(standNo, Command.Reset);

			// //再修改实时状态数据
			// DevHopeStatusCache[standNo].Set(56);
			// Logger.Info($"期望值,{standNo},{56}");
			// DealHopeDataCommandAck(standNo, DeviceStatus.Awaiting);

			i += iCurLength
			dealLength = i
			// 暂停
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WR" && (i+8 <= buffer.ValidLength) {
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			// DealCommandAck(standNo, Command.Pause);

			// //再修改实时状态数据
			// DevHopeStatusCache[standNo].Set(65);
			// Logger.Info($"期望值,{standNo},{65}");
			// DealHopeDataCommandAck(standNo, DeviceStatus.Pause);
			i += iCurLength
			dealLength = i
			// 运行
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WT" && (i+8 <= buffer.ValidLength) {
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			// DealCommandAck(standNo, Command.Run);

			// //再修改实时状态数据
			// DevHopeStatusCache[standNo].Set(66);
			// Logger.Info($"期望值,{standNo},{66}");
			// DealHopeDataCommandAck(standNo, DeviceStatus.Run);
			i += iCurLength
			dealLength = i
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WM" && (i+8 <= buffer.ValidLength) {
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			// DealCommandAck(standNo, Command.Manual);

			i += iCurLength
			dealLength = i
		} else {
			fmt.Println("else...")
			i++
		}

	}

	if dealLength > 0 {
		h.Buffer.Clear(dealLength)
	}
	if h.Buffer.ValidLength > MaxPackageLength {
		h.Buffer.Clear(h.Buffer.ValidLength - MaxPackageLength)
	}
	fmt.Println("ValidLength", h.Buffer.ValidLength)
}

func (h *HxDevice) Close() {
	h.Conn.Close()
}
