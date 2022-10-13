package lru

import (
	"fmt"
	"testing"

	"github.com/pieterclaerhout/go-log"
)

func TestMain(m *testing.M) {
	c := Constructor(3)
	c.Put(1, 1) // 1: 1
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Put(5, 5)
	c.Put(6, 6)

	// c.Put(3, 3)
	// c.Get(1) // 1: 2->1
	// c.Get(1) // 1: 2->1
	// c.Get(1) // 1: 2->1
	// c.Put(3, 3)
	// c.Get(3)
	// c.Get(3)
	// c.Get(2) // 1: 2->1
	// c.Get(2) // 1: 2->1
	// fmt.Println("GetMinFreq:", c.GetMinFreq())
	c.ShowValues()

	return
}

type LRUCache struct {
	Items map[int]*Node
	Root  *Node
	Cap   int
}

type Node struct {
	Key        int
	Val        int
	Next, Prev *Node
}

func Constructor(n int) *LRUCache {
	return &LRUCache{
		Cap:   n,
		Items: make(map[int]*Node),
		// List:  make([]*Node, 0),

	}
}

// func (this *LRUCache) Tail() *Node {
// return this.List[len(this.List)-1]
// }

// func (this *LRUCache) GetMinFreq() int {
// 	return this.MinFreq
// }

func (this *LRUCache) ShowFreqs() {
	log.InfoDump(this.Root, "root")
	log.InfoDump(this.Root.Prev, "tail")
	log.InfoDump(this.Items, "Items")
	// log.InfoDump(this.Freqs[this.MinFreq], "minFreq")
	// log.InfoSeparator(this.Freqs[this.MinFreq])
}

func (this *LRUCache) ShowValues() {
	for k, v := range this.Items {
		fmt.Println(k, v.Val)
	}
}

func (this *LRUCache) Put(k int, v int) {
	oldnod, ok := this.Items[k]
	if ok {
		oldnod.Val = v
		return
	}
	if len(this.Items) == this.Cap {
		this.DelOld()
	}
	// new
	// nod:=&Node{Key: k, Val: v,Prev:this.Root.Prev,Next:this.Root}
	nod := &Node{Key: k, Val: v}

	if this.Root == nil {
		this.Root = nod
		nod.Prev = nod
		nod.Next = nod
	} else {
		tail := this.Root.Prev

		nod.Prev = tail
		nod.Next = this.Root

		// 修改之前的root的下一个和尾节点的指向
		this.Root = nod
		nod.Next.Prev = nod
		tail.Next = nod

	}
	this.Items[k] = nod
}

func (this *LRUCache) Get(k int) int {
	nod, ok := this.Items[k]
	if !ok {
		return -1
	}
	// 处理前后关系
	prev, next := nod.Prev, nod.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	// 提到root位置
	this.Root.Prev = nod
	nod.Prev = this.Root.Prev
	this.Root = nod
	return nod.Val
}

func (this *LRUCache) DelOld() {
	old := this.Root.Prev
	this.Root.Prev = old.Prev
	delete(this.Items, old.Key)
}
