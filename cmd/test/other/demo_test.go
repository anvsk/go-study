package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"sync"
	"syscall"
	"time"
	"unsafe"

	_ "net/http/pprof"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
var wg sync.WaitGroup

func main() {
	// println("dasda")
	// s := fmt.Sprintln("aaaa", wg, rr)
	// fmt.Println(s)
	// TestMostT()
}

func main2() {

	// util.InitUtil()

	// db.InitDB()
	// var res AnswerQuestion
	// db.Orm.First(&res)
	// plog.InfoDump(res, "res")
	// plog.Info(res.CreatedAt)
	// plog.Info(res.CreatedAt.Unix())
	// plog.Info(time.Unix(res.CreatedAt.Unix(), 0))

	// s := cuetomMap{}
	// s.Store("a", "a")
	// s.Store("b", "b")
	// fmt.Println(s.Load("a"))
	// fmt.Println(s.Lists()...)
	// aaa := &aacd{"1"}

	// s := sync.Map{}
	// s.Store("a", aaa)
	// // v, ok := s.LoadOrStore("a", 2)

	// // fmt.Println(v)
	// // fmt.Println(ok)
	// // fmt.Println(s.Load("a"))

	// aaa.name = "2"
	// fmt.Println(aaa)
	// change(aaa)
	// fmt.Println(aaa)
	// fmt.Println(s.Load("a"))
	// b := *aaa
	// fmt.Println(b)
	// change2(aaa)
	// fmt.Println(aaa)
	// fmt.Println(b)
	// data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13}
	// // tipBytes := make([]byte, len(data)-8)
	// tipBytes := []byte{}
	// fmt.Println(tipBytes)
	// // for i := 0; i < len(data)-8; i++ {
	// // 	tipBytes[i] =
	// // }
	// tipBytes = append(tipBytes, data[5:(5+len(data)-8)]...)
	// fmt.Println(tipBytes)

	a := 88

	fmt.Println(a)
	fmt.Println(time.Now())
	// fmt.Println(a == nil)
}

func test1() *aacd {
	return nil
}

func change(a *aacd) {
	a.name = "3"

}

func change2(a *aacd) {
	a.name = "333"

}

type aacd struct {
	name string
}

type cuetomMap struct {
	sync.Map
}

func (s *cuetomMap) Lists() (res []interface{}) {
	s.Range(func(key, value interface{}) bool {
		res = append(res, value)
		return true
	})
	return
}

type AnswerQuestion struct {
	gorm.Model
	UserID int
}

func (*AnswerQuestion) TableName() string {
	return "answer_question"
}

// func main() {
// 	ch2 := make(chan int, 1)
// 	fmt.Println("发送前")
// 	ch2 <- 10
// 	fmt.Println("发送成功")
// }

type sst struct {
	s [3]int
}

func (i *sst) ss() {
	i.s[1] = 99
	// i.s = [3]int{1, 1, 1}
}

func TransferBytes2AsciiCodeString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func Utf8ToGbk(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}

func DebugHex(modelData []byte) {
	fmt.Println("hexData----------->")
	for i := 0; i < len(modelData); i++ {
		fmt.Printf("%s ", strconv.FormatInt(int64(modelData[i]), 16))
	}
	fmt.Println("hexData----------->")
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func IntToBytesLittle(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

// func IntToBytes(n int) []byte {
// 	x := int32(n)
// 	bytesBuffer := bytes.NewBuffer([]byte{})
// 	binary.Write(bytesBuffer, binary.BigEndian, x)
// 	return bytesBuffer.Bytes()
// }

func PadLeft(oriStr string, targetLength int, padChar string) string {
	length := len(oriStr)
	if length < targetLength {
		differLength := targetLength - length
		var i int
		var buffer bytes.Buffer
		for i = 0; i < differLength; i++ {
			buffer.WriteString(padChar)
		}
		buffer.WriteString(oriStr)
		return buffer.String()
	}

	return oriStr
}

func GetAsciiStr(input string, ModelLen, ModelNo int) []byte {
	var res []byte
	b := Transfer2AsciiCodeByte(input)
	if len(b) == 0 {
		for i := 0; i < ModelLen; i++ {
			res = append(res, 32)
		}
	} else {
		for i := 0; i < ModelLen; i++ {
			if i < len(b) {
				res = append(res, b[i])
			} else {
				if ModelNo != 0 {
					res[i] = 48
				} else {
					res[i] = 32
				}
			}
		}
	}
	return res
}

func buildSend2DeviceDataForCraftCommandCraftNames(s string) []byte {
	var result []byte
	/**
	工艺不足20字符的，右侧补空格
	转assic码，取前20个字节
	统计前20个字节中大于127的字节个数
	大于127的字节个数为奇数个，则将末位变成32
	*/
	str1 := PadRight(s, 20, " ")

	result = append(result, []byte(str1)[:20]...)
	num1 := 0
	for index := 0; index < len(result); index++ {
		if result[index] > 127 {
			num1++
		}
	}
	if num1%2 > 0 {
		result[19] = 32
	}
	return result
}

func buildSend2DeviceDataForCraftCommandCraftNames2(s string) []byte {
	var result []byte
	/**
	工艺不足20字符的，右侧补空格
	转assic码，取前20个字节
	统计前20个字节中大于127的字节个数
	大于127的字节个数为奇数个，则将末位变成32
	*/
	str1 := PadRight(s, 20, " ")

	result = append(result, []byte(str1)[:20]...)
	num1 := 0
	for index := 0; index < len(result); index++ {
		if result[index] > 127 {
			num1++
		}
	}
	if num1%2 > 0 {
		result[19] = 32
	}
	return result
}

func UTF82GB2312(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func PadRight(oriStr string, targetLength int, padChar string) string {
	length := len(oriStr)
	var buffer bytes.Buffer
	buffer.WriteString(oriStr)
	if length < targetLength {
		differLength := targetLength - length
		var i int
		for i = 0; i < differLength; i++ {
			buffer.WriteString(padChar)
		}
		return buffer.String()
	}

	return oriStr
}

// public static byte[] ByteToCharArray(byte val)
//         {
//             byte[] Model = new byte[2];
//             byte ModelPriority = (byte)(val >> 4 & 15);
//             if (!(ModelPriority > 9))
//             {
//                 Model[0] = (byte)(48 + ModelPriority);
//             }
//             else if ((ModelPriority > 9 && ModelPriority < 16))
//             {
//                 Model[0] = (byte)(55 + ModelPriority);
//             }
//             ModelPriority = (byte)(val & 15);
//             if (!(ModelPriority > 9))
//             {
//                 Model[1] = (byte)(48 + ModelPriority);
//             }
//             else if ((ModelPriority > 9 && ModelPriority < 16))
//             {
//                 Model[1] = (byte)(55 + ModelPriority);
//             }
//             return Model;
//         }

func ByteToCharArray(val byte) []byte {
	Model := [2]byte{}
	ModelPriority := byte(val >> 4 & 15)
	if ModelPriority <= 9 {
		Model[0] = byte(48 + ModelPriority)
	} else if ModelPriority > 9 && ModelPriority < 16 {
		Model[0] = byte(55 + ModelPriority)
	}
	ModelPriority = byte(val & 15)
	if ModelPriority <= 9 {
		Model[1] = byte(48 + ModelPriority)
	} else if ModelPriority > 9 && ModelPriority < 16 {
		Model[1] = byte(55 + ModelPriority)
	}
	res := make([]byte, 2)
	copy(res, Model[:])
	return res
}

func DoubleByteToCharArray(val uint16) []byte {
	DataList := [4]byte{}
	charArray := ByteToCharArray(byte(val >> 8))
	DataList[0] = charArray[0]
	DataList[1] = charArray[1]
	charArray = ByteToCharArray(byte(val))
	DataList[2] = charArray[0]
	DataList[3] = charArray[1]
	res := make([]byte, len(DataList))
	copy(res, DataList[:])
	return res
}

//         public static byte[] DoubleByteToCharArray(ushort val)
//         {
//             var DataList = new byte[4];
//             byte[] charArray = ByteToCharArray((byte)(val >> 8));
//             DataList[0] = charArray[0];
//             DataList[1] = charArray[1];
//             charArray = ByteToCharArray((byte)val);
//             DataList[2] = charArray[0];
//             DataList[3] = charArray[1];
//             return DataList;
//         }

// func DoubleByteToCharArray(val uint16) []byte {
// 	DataList := [4]byte{}
// 	charArray := ByteToCharArray(byte(val >> 8))
// 	DataList[0] = charArray[0]
// 	DataList[1] = charArray[1]
// 	charArray = ByteToCharArray(byte(val >> 8))
// 	DataList[2] = charArray[0]
// 	DataList[3] = charArray[1]
// 	res := make([]byte, len(DataList))
// 	copy(res, DataList[:])
// 	return res
// }

func Transfer2AsciiCodeByte(input string) []byte {
	var result []byte
	for _, char := range []rune(input) {
		result = append(result, byte(char))
	}
	return result
}

type s1 struct {
	Ctype string
	Value interface{}
}

type s2 struct {
	Ctype string
	Value Vv
}

type Vv struct {
	A string
	B Cc
}

type Cc struct {
	I int
}

const (
	GameCallScore int = 9

	GameWaitting = iota
	GamePlaying
	GameEnd
)

const (
	GameCallScore2 = iota
	// GameWaitting2  int = 9
	bb
	GamePlaying2
	aa
	GameEnd2
)

const (
	GameWaitting3  int = 9
	GameCallScore3 int = 99
	GamePlaying3       = iota
)

func gogogo() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func gogogo2() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func gogogo3() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}

func partition(array []int, i int, j int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[i]
	for i < j {
		for j > i && array[j] > pivot {
			j--
		}
		if j > i {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < pivot {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = pivot
	return i
}

func quicksort(array []int, low int, high int) {
	var pivotPos int //划分基准元素索引
	if low < high {
		pivotPos = partition(array, low, high)
		quicksort(array, low, pivotPos-1)
		quicksort(array, pivotPos+1, high)
	}
}

func quick_sort(li []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	li[i], li[r] = li[r], li[i]
	tmp := li[i]
	for i < j {
		for i < j && li[j] >= tmp {
			j--
		}
		li[i] = li[j]
		for i < j && li[i] <= tmp {
			i++
		}
		li[j] = li[i]
	}
	li[i] = tmp
	quick_sort(li, left, i-1)
	quick_sort(li, i+1, right)
}

func bubble_sort(li []int) {
	for i := 0; i < len(li)-1; i++ {
		exchange := false
		fmt.Println("i=", i)

		for j := 0; j < len(li)-i-1; j++ {
			fmt.Println("j", j)

			if li[j] > li[j+1] {
				fmt.Println("===li[j]")
				fmt.Println(li[j])
				fmt.Println(li[j+1])
				fmt.Println("===li[j]")

				li[j], li[j+1] = li[j+1], li[j]
				exchange = true
			}
		}
		fmt.Println(exchange)

		if !exchange {
			return
		}
	}
}

func getslicecap() {
	a := []byte{}
	a = append(a, 1)
	fmt.Println("cap of a is ", cap(a))

	b := []int{23, 51}
	b = append(b, 4, 5, 6)
	fmt.Println("cap of b is ", cap(b))

	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Println("cap of c is ", cap(c))

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}

	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"})
	fmt.Println("cap of d is ", cap(d))
}

func testclosechan() {
	// ch1 := make(chan int, 30)
	// go func() {
	// 	for v := range ch1 {
	// 		time.Sleep(time.Second)
	// 		println("aaaa", v)
	// 	}
	// }()
	// ch1 <- 11
	// ch1 <- 11
	// ch1 <- 11
	// close(ch1)
	// for {
	// }

	// reschan如果没有close会一直阻塞
	// reschan := make(chan int, 10)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	reschan <- 999
	// 	// close(reschan)
	// }()

	// for v := range reschan {
	// 	println(v, " task finished ")
	// }
	// println(" script finished ")
}

type smallobj struct {
	arr [1 << 10]byte
	i   int
}

func consume(n int) {
	ch := make(chan int, n)
	go limitPerSecond(ch, n)
	for i := 1; i < 10000; i++ {
		ch <- i
	}
	for {
	}
}

// 生产者、限流器、消费者模型
// 限流器,分发函数
// 每秒通过的请求数量10？
func limitPerSecond(rec chan int, limit int) {
	for {
		<-time.Tick(time.Second)
		i := 0
		for {
			go myworker(<-rec)
			i++
			if i == limit {
				break
			}
		}
	}
}

func myworker(i int) {
	// time.Sleep(200 * time.Millisecond)
	println(i)
}

func serveHttp() {
	http.HandleFunc("/dhaha", func(rw http.ResponseWriter, r *http.Request) {
		b := "haha"
		rw.Write([]byte(b))
	})

	http.Get("aa")
	http.ListenAndServe(":9999", nil)
	for {
	}
}

func testContext() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go func() {
		<-time.After(20 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		println("aasdsad")
		tt, bool := ctx.Deadline()
		fmt.Println(tt)
		println(bool)
		fmt.Println(ctx.Err())
	}
}

func createM() {
	ch := make(chan int)
	for i := 0; i < 300; i++ {
		go func() {
			time.Sleep(5 * time.Second)
			ch <- 1
		}()
		go func() {
			println("aaa")
		}()
		go func(i int) {
			println(i)
			time.Sleep(3000 * time.Second)
			println(i, "end")

		}(i)
		time.Sleep(80 * time.Millisecond)
	}
	<-ch
	<-time.After(5 * time.Second)
}

func SignalHandle() {
	for {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT)
		signal.Notify(ch, syscall.SIGKILL)
		sig := <-ch
		fmt.Printf("收到信号：%d %s\n", sig, sig.String())
		switch sig {
		case syscall.SIGINT:
			println("我要开始休眠三秒sigint")
			<-time.After(3 * time.Second)
		case syscall.SIGKILL:
			println("我要开始休眠三秒sigkill")
			<-time.After(3 * time.Second)
			// os.Exit(1)
		default:
			println("我进来default了")
		}
	}
}

func testblock(to time.Duration) int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		ch <- 1
	}()
	select {
	case <-time.After(to):
		return -1
	case res := <-ch:
		return res
	}

}

// 三个打印函数，要求用三个携程顺序打印各个N次
func sortprintln(n int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	closech := make(chan bool)
	go pcat(ch1, ch3, n)
	go pdog(ch2, ch1)
	go pbird(ch3, ch2, closech)

	<-closech
}

func pcat(selfchan, prevchan chan int, times int) {
	i := 0
	for {
		i++
		if i == 1 {
			println("begining... ...")
			println("cat:%d", i, "==")
			selfchan <- 1
			continue
		}
		<-prevchan

		println("cat:%d", i, "==")

		if i == times {
			println("finished... ...ch1")
			close(selfchan)
			break
		}
		selfchan <- 1
	}
}

func pdog(selfchan, prevchan chan int) {
	i := 0
	for {

		tmp := <-prevchan
		if tmp == 0 {
			println("finished... ...ch2")
			close(selfchan)
			break
		}
		i++
		println("dog:%d", i, "==")

		selfchan <- 1
	}
}

func pbird(selfchan, prevchan chan int, closech chan bool) {
	i := 0
	for {

		tmp := <-prevchan
		if tmp == 0 {
			println("finished... ...ch3")
			closech <- true
			break
		}
		i++
		println("bird:%d", i, "==")

		selfchan <- 1
	}
}

func plast(ch1, ch2, ch3 chan struct{}) {
	<-ch2
	println("2:last")
	ch3 <- struct{}{}
}

func slicereal() {
	x := []int{1, 2, 3, 4, 5, 6} //cap 6
	y := x[4:5]                  //cap 2 value[5,6]
	log.Println(cap(y))          //
	log.Println((y))             //
	y[0] = 99
	y = append(y, 10, 10, 10, 10)

	log.Println(x) // 99 ,2,3,4,5,6
	log.Println(y) // 5

	// y = append(y, 10)
	// y[0] = 88
	// log.Println(x) // 88 ,2,3,4,5,6

	// y = append(y, 10)
	// y[0] = 77
	// log.Println(x) // 88 ,2,3,4,5,6

	// y = append(y, 10)

	// y = append(y, 50)// [5,6,50]
	// log.Println(x)//
	// y = append(y, 60)//[5,6,50,60]
	// log.Println(x)//
	// log.Println(y)//
	// y[0] = 20
	// log.Println(y)
}

func sliceBiggerCapNums() {
	a := []byte{1, 0}
	a = append(a, 1, 1, 1)
	fmt.Println("cap of a is ", cap(a))

	b := []int{23, 51}
	b = append(b, 4, 1, 1)
	// b = append(b, 4)
	// b = append(b, 4)
	fmt.Println("cap of b is ", cap(b))

	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Println("cap of c is ", cap(c))

	// e := []float64{1, 23}
	// e = append(e, 2, 5, 6)
	// fmt.Println("cap of e is ", cap(e))

	type D struct {
		age  byte
		name string
	}
	d := []D{
		{1, "123"},
		{2, "234"},
	}

	d = append(d, D{4, "456"}, D{5, "567"}, D{6, "678"}, D{6, "678"})
	fmt.Println("cap of d is ", cap(d))

	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))

	f := []int{1, 2, 3}
	fmt.Println("cap of f before:", cap(f))
	f = append(f, 4)
	fmt.Println("cap of f after:", cap(f))
}

func myrand() int {
	return rr.Intn(9999999)
}

// 测试有缓冲相关
func nocachechan() {
	ch1 := make(chan int, 90)
	log.Println(<-ch1)
	log.Println(<-ch1)
	log.Println(<-ch1)
}

// 仅用chan实现 同时N线程下载效果
func forrgetchannel() {
	ch1 := make(chan int, 5)
	down := make(chan int)
	go func() {
		log.Println("<-ch1begin")
		<-time.After(1 * time.Second)
		log.Println("<-ch1end")
		for i := 0; i < 10; i++ {
			log.Printf("开启第%d个消费携程", i)
			go func(i int, wg2 *sync.Mutex) {
				for {
					wg2.Lock()

					tmp := <-ch1
					if tmp == 0 {
						log.Println("==消费g", i, "downed")
						down <- 1
						break
					}
					log.Println("==consume==", tmp)
					<-time.After(2 * time.Second)
					wg2.Unlock()
				}
			}(i, &sync.Mutex{})
		}

	}()

	go func() {
		for i := 1; i < 5000; i++ {
			ch1 <- i
			log.Println("==put==", i)
		}
		close(ch1)
	}()
	<-down
}

func selectblock4() {
	ch1 := make(chan int, 0)
	n := 10
	count := 1

	go func() {
		log.Printf("begin handle task... ...")
		<-time.After(1 * time.Second)
		for {
			log.Println("begin for")
			<-time.After(1 * time.Second)

			tmp := <-ch1
			log.Println(tmp)

			if tmp == 0 {
				log.Printf("closed channel")
				break
			}
			log.Println("end for")
			count++
			if count > 15 {
				break
			}
		}
	}()
	for i := 1; i < n; i++ {
		ch1 <- i
	}
	// close(ch1)
	<-time.After(8000 * time.Millisecond)
	log.Printf("finished")
}

func selectblock2() {
	ch1 := make(chan int, 0)

	go func() {
		log.Printf("begin handle task... ...")
		<-time.After(1 * time.Second)
		for {
			log.Println("begin for")

			tmp := <-ch1
			log.Println(tmp)

			if tmp == 0 {
				log.Printf("closed channel")
				break
			}
			log.Println(<-ch1)
			log.Println("next for")
		}
	}()
	for i := 1; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	<-time.After(8000 * time.Millisecond)
	log.Printf("finished")
}

func selectblock3() {
	ch1 := make(chan int, 0)

	ch1 <- 1
	go func() {
		time.After(2 * time.Second)
		fmt.Println(<-ch1)
	}()

	time.After(1 * time.Second)
}

func selectblock() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(00)
	go func() {
		ch2 <- 2
		ch1 <- 1
		ch2 <- 2
		<-time.After(4 * time.Second)

		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			goto aa
		case <-ch1:
			fmt.Println(11)
		case <-ch2:
			fmt.Println(22)
		default:
			fmt.Println("default")
			<-time.After(500 * time.Millisecond)
		}
	}
aa:

	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	log.Println(<-ctx.Done())
	fmt.Println(99)
}

func sss() {
	// util.InitUtil()
	// cache.InitCache()
	// key := "ccc"
	// cache.C.Set(key, 123, 10*time.Second)
	// log.Debug(cache.C.Get(key))
	// db.InitDB()
	// for i := 0; i < 1000000000000000000; i++ {
	//     // <-time.After(2 * time.Millisecond)
	//     log.Debug(i)
	//     go func(ii int) {
	//         time.After(100 * time.Millisecond)
	//         log.Info(ii)
	//     }(i)
	//     // go db.TestMysql()
	//     // go db.TestCH()

	// }
	// <-time.After(1 * time.Hour)
	// a := "aaaa"
	// fmt.Println(a.Len())

	// sync.TestCond()
	// pprof.Testranddomstr()

	// fmt.Println(leetcode.Stradd("98", "55"))
	// ss := "klsadjla"
	// fmt.Println(ss[2])
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 3; i++ {
	// fmt.Println(myrand())
	// fmt.Println(myrand())
	// fmt.Println(myrand())
	// fmt.Println("--====--")
	// }

	// for i := 0; i < 5; i++ {
	//     rand.Seed(time.Now().UnixNano())
	//     fmt.Println(rand.Intn(100))
	// }

	// leetcode.Test()
	go func() {
		http.ListenAndServe(":9008", nil)
	}()
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// //启动trace goroutine
	// err = trace.Start(f)
	// if err != nil {
	// 	panic(err)
	// }
	// defer trace.Stop()

	// log.SetFlags(0)
	// log.Println("begin", carbon.Now().ToDateTimeString())
	// a := []int{1}
	// a = append(a, 1, 2, 3)
	// println(a)
	// println(append([]byte("hello "), "world"...))
	// fmt.Println(string(append([]byte("hello "), "world"...)))
	// t1 := time.Now()
	// time.Sleep(time.Second)
	// t2 := time.Since(t1)
	// println(t2)
	// sliceBiggerCapNums()
	// sortprintln(5)
	// SignalHandle()
	// testContext()
	// consume(15)
	// sync2.Comparesyncchan()
	// getslicecap()
	// ctx,_:=context.WithTimeout(context.Background(),time.Second)
	// testclosechan()
	// go gogogo()
	// go gogogo()
	// go gogogo2()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo3()
	// go gogogo()
	// gogogo()
	// println(GameCallScore)
	// println(GameCallScore2)
	// println(GameCallScore3)
	// println(GamePlaying)
	// println(GamePlaying2)
	// println(GamePlaying3)
	// println(GamePlaying3)
	// println(aa)
	// println(GameEnd2)
	// s := []int{1, 2, 3, 4, 5, 6}
	// a := s[:3]
	// a = append(a, 7)
	// fmt.Println(s, a) //[6/6]0x14000157f38 [4/6]0x14000157f38 [1 2 3 7 5 6] [1 2 3 7]
	// b := append(a, []int{8, 8, 9}...)
	// fmt.Println(s, a, b)//[1 2 3 7 5 6] [1 2 3 7] [1 2 3 7 8 8 9]
	// a 1,2,3,7,5,6
	// s =a
	// var x int
	// t := runtime.GOMAXPROCS()
	// fmt.Println(runtime.NumCPU())
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		for {
	// 			x++
	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second)
	// fmt.Println("x =", x)
	// log.Println("end", carbon.Now().ToDateTimeString())

	// for {
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	<-time.After(3 * time.Second)
	// 	cancel()
	// }()

	// fmt.Println(<-ctx.Done())
	// fmt.Println(<-ctx.Done())
	// fmt.Println(<-ctx.Done())
	// fmt.Println(<-ctx.Done())
	// cankao, _ := time.Parse("2006-01-02", "2021-08-24")
	// fmt.Println(time.Since(cankao) > 74*time.Hour)
	// println("aaaa")
	// ch := make(chan int)
	// go func() {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// }()
	// ch <- 1
	// close(ch)

	// b1, err := json.Marshal(s1{
	// 	Ctype: "string",
	// 	Value: Vv{
	// 		A: "AAA",
	// 		B: Cc{
	// 			I: 9999,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b1))
	// var ss1 s1
	// err = json.Unmarshal(b1, &ss1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ss1")
	// fmt.Println(ss1)
	// plog.InfoDump(ss1, "ss1")

	// println("=====")

	// b2, _ := json.Marshal(ss1)
	// var ss2 s2
	// json.Unmarshal(b2, &ss2)
	// plog.InfoDump(ss2, "ss2")

	// b2, err := json.Marshal(s1{
	// 	"int",
	// 	2,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b2))

	// type testSturct struct {
	// 	Name string
	// 	Num  int
	// }

	// testStructs := []testSturct{
	// 	{
	// 		Name: "test2",
	// 		Num:  2,
	// 	},
	// 	{
	// 		Name: "test3",
	// 		Num:  3,
	// 	},
	// 	{
	// 		Name: "test1",
	// 		Num:  1,
	// 	},
	// }

	// sort.Slice(testStructs, func(i, j int) bool {
	// 	return testStructs[j].Num > testStructs[i].Num // 升序
	// 	// return testStructs[i].Num > testStructs[j].Num  // 降序
	// })
	// var aa []testSturct

	// for _, v := range testStructs {
	// 	if v.Num > 2 {
	// 		continue
	// 	}
	// 	aa = append(aa, v)
	// }
	// log.Println(testStructs)
	// log.Println(aa)
	// realCommand := []byte{0, 1, 0, 0, 0, 6, 1, 3, 0, 0, 0, 49, 1, 1, 1}
	// sss := "\x00\x01\x00\x00\x00\x06\x01\x03\x00\x00\x001\x01\x01\x01"
	// fmt.Println(realCommand)
	// plog.InfoDump(string([]byte{59, 59, 59}), " string([]byte{57, 57, 57})")
	// plog.InfoDump(string(realCommand), "string(command)")
	// fmt.Println(string([]byte("\x00\x01\x00\x00\x00\x06\x01\x03\x00\x00\x001\x01\x01\x01")) == string(realCommand))
	// fmt.Println(sss)
	// real := "000100000065010362000000010000000000000000000000000001000000000000000000010000000000000000000000000000000000000000000000000000000500000000000002c800000a1501be00000000000000000000000000000000000000000000000000000000"

	// fmt.Println([]byte(real))

	// ss := "999"
	// buf := []byte(ss)
	// plog.InfoDump(buf)
	// plog.InfoDump(buf[0])
	// plog.InfoDump(string(buf))
	// buf := []byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x65, 0x01, 0x03, 0x62, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xc8, 0x00, 0x00, 0x0a, 0x15, 0x01, 0xbe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	// fmt.Println(buf[5])
	// fmt.Println(string(buf))
	// buf := []byte{0, 1, 0, 0, 0, 11, 1, 16, 0, 43, 0, 3, 6, 0, 1, 0, 0, 0, 0}
	// buf2 := []byte{111, 116, 104, 101, 114, 32, 109, 101, 115, 115, 97, 103, 101}
	// fmt.Println(string(buf2))
	// fmt.Println(int(byte(88)))
	// fmt.Println(int64(byte(88)))
	// fmt.Println(int32(byte(88)))
	// fmt.Println((255))
	// test, _ := hex.DecodeString("9")
	// fmt.Println(test)
	// fmt.Println(bytes.Compare(test, 190)) // 0

	// fmt.Println((Transfer2AsciiCodeByte("ST哈哈哦s")))
	// fmt.Println([]byte(rune(("ST哈哈哦s"))))
	// fmt.Println(ByteToCharArray(byte(245))) //70 53

	// fmt.Println(Transfer2AsciiCodeByte(strconv.FormatInt(int64(245), 16)))
	// fmt.Println(len(PadRight(" ", 5, "a")))
	// fmt.Println((PadRight(" ", 5, "a")))
	// str := "中国"
	// fmt.Println(Transfer2AsciiCodeByte("abcabc"))
	// fmt.Println([]byte("中国"))
	// // res, _ := UTF82GB2312([]byte("中国"))
	// // fmt.Println(res)
	// fmt.Println([]rune(str))

	// sText := "中国"
	// textQuoted := strconv.QuoteToASCII(sText)
	// textUnquoted := textQuoted[1 : len(textQuoted)-1]
	// fmt.Println(textUnquoted)

	// sUnicodev := strings.Split(textUnquoted, "\\u")
	// var context string
	// for _, v := range sUnicodev {
	// 	if len(v) < 1 {
	// 		continue
	// 	}
	// 	temp, err := strconv.ParseInt(v, 16, 32)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	context += fmt.Sprintf("%c", temp)
	// }
	// fmt.Println(context)
	// fmt.Println([]rune(sText))
	// fmt.Println([]rune(context))

	// enc := mahonia.NewDecoder("gbk")
	// fmt.Println(([]byte(enc.ConvertString(context))))
	// fmt.Println([]byte(enc.ConvertString(sText)))
	// s := []int{1, 2, 3}
	// f := []int{1, 2, 3}
	// a := append(f, s[1:2]...)
	// fmt.Println(a)
	// fmt.Println(DoubleByteToCharArray(uint16(18)))
	// fmt.Println(buildSend2DeviceDataForCraftCommandCraftNames("i文华懂啊哇"))
	// fmt.Println(len(buildSend2DeviceDataForCraftCommandCraftNames("i文华懂啊哇")))
	// fmt.Println(GetAsciiStr("haskdjh7897", 11, 0))

	// fmt.Println(PadLeft("a33", 7, "xxxsa"))
	// fmt.Println(Transfer2AsciiCodeByte("345.9"))
	// fmt.Println(uint16(0))
	// fmt.Println(uint16(1))
	// fmt.Println(GetAsciiStr("011.5", 5, 1))
	// fmt.Println(PadLeft(Transfer2AsciiCodeByte([]byte("100.5")), 4, "0"))
	// DebugHex(ByteToCharArray(99))
	// DebugHex(Transfer2AsciiCodeByte("DL"))
	// DebugHex(buildSend2DeviceDataForCraftCommandCraftNames("1107-"))
	// DebugHex(buildSend2DeviceDataForCraftCommandCraftNames("1107-黑"))
	// DebugHex(Utf8ToGbk([]byte("1107-黑")))
	// s := []byte{0x30}
	// var i1 int64 = 100 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]

	// s1 := make([]byte, 4)
	// buf := bytes.NewBuffer(s1)

	// // 数字转 []byte, 网络字节序为大端字节序
	// binary.Write(buf, binary.BigEndian, i1)
	// fmt.Println(buf.Bytes())

	// buf.Reset()
	// binary.Write(buf, binary.LittleEndian, i1)
	// fmt.Println(buf.Bytes())

	// fmt.Println(byte(999))
	// fmt.Println(DoubleByteToCharArray(0))
	// fmt.Println(IntToBytesLittle(998))
	// fmt.Println(TransferBytes2AsciiCodeString([]byte{85, 80}))
	// res := [9]int{}
	// a := []int{1, 2, 3, 4, 5}
	// copy(res[:], a)
	// fmt.Println(res)
	// ioutil.WriteFile(`保存文件的文件名.txt`, []byte("888"), 0666)
	// v := sst{
	// 	s: [3]int{
	// 		1, 2, 3,
	// 	},
	// }
	// a := v.s[1]
	// println(a)
	// (&v).ss()
	// b := v.s[1]
	// println(a)
	// println(b)
	// ch := make(chan int, 10)
	// go func() {
	// 	for {
	// 		num := <-ch
	// 		if num < 20 {
	// 			<-time.After(1 * time.Second)
	// 		} else {
	// 			<-time.After(4 * time.Second)
	// 		}
	// 		fmt.Println(num, "执行完了!")
	// 	}
	// }()

	// for i := 0; i < 100; i++ {
	// 	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println(i, "添加时超时!!!")
	// 	case ch <- i:
	// 		fmt.Println(i, "任务添加了!!!")
	// 	}
	// }

	// for {
	// }
	// fmt.Println(ModelTypeInout)

	// ch := make(chan int)
	// ch <- 10
	// fmt.Println("发送成功")
	// fmt.Println(time.Now().Zone())
	// fmt.Println(time.Now().Local().Zone())
	// fmt.Println(time.Now().UTC().Zone())

	// fmt.Println(time.Local)
	// fmt.Println(time.ParseInLocation())
	// name, offset := time.Now().Zone()
	// fmt.Println(name, offset)

	// parsed, _ := time.Parse("2006-01-02 15:04", "2021-09-13 07:05")

	// fmt.Println(parsed)
	// fmt.Println(time.Unix(parsed.Unix(), 0))

	// fmt.Println("=====")

	// parsed2, _ := time.ParseInLocation("2006-01-02 15:04", "2021-09-13 07:05", time.Local)
	// fmt.Println(parsed2)
	// fmt.Println(time.Unix(parsed2.Unix(), 0))
}
