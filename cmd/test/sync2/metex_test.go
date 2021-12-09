// 用channel实现sync包的
// mutex
// Rmutex
package sync2

import (
	"fmt"
	"testing"
	"time"
)

type Mutex struct {
	Sem chan struct{}
}

func (m *Mutex) Lock() {
	// 判断是否初始化
	if m.Sem == nil {
		m.Sem = make(chan struct{}, 1)
	}
	m.Sem <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.Sem
}

func (m *Mutex) Close() {
	close(m.Sem)
}

// 并发写map测试锁
func TestMutex(*testing.T) {
	lock := Mutex{}

	ch := map[string]int{
		"concur": 2,
	}

	for ii := 0; ii < 30; ii++ {
		tmp := ii
		// 开十个线程
		go func(ch *map[string]int, tmp int) {
			for i := 0; i < 3; i++ {
				lock.Lock()
				(*ch)["concur"] = i
				println("线程", tmp, "读取到", i)
				lock.Unlock()
			}
		}(&ch, tmp)
	}
	fmt.Println(ch)
	time.Sleep(3 * time.Second)
}

// 信号量的方式

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	// 只有在s还有空间的时候才能发送成功
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	// 为其他信号量腾出空间
	<-s
}

func TestSemaphore(*testing.T) {
	// lock := Mutex{}
	lock := NewSemaphore(1)
	ch := map[string]int{
		"concur": 1,
	}

	for ii := 0; ii < 30; ii++ {
		tmp := ii
		// 开十个线程
		go func(ch *map[string]int, tmp int) {
			for i := 0; i < 3; i++ {
				lock.Lock()
				(*ch)["concur"] = i
				println("线程", tmp, "读取到", i)
				lock.Unlock()
			}
		}(&ch, tmp)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(ch)

}

// 用纯chan的方式
func TestJustChan(*testing.T) {
	lock := make(chan int, 1)
	ch := map[string]int{
		"concur": 1,
	}
	for ii := 0; ii < 50; ii++ {
		tmp := ii
		// 开十个线程
		go func(ch *map[string]int, tmp int) {
			for i := 0; i < 3; i++ {
				lock <- i
				(*ch)["concur"] = i
				println("线程", tmp, "读取到", i)
				// time.Sleep(time.Second)
				<-lock
			}
		}(&ch, tmp)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(ch)

}
