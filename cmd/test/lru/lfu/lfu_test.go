package main

import (
	"fmt"
	"testing"

	"github.com/pieterclaerhout/go-log"
)

func main() {
	c := Constructor(2)
	c.Put(1, 1) // 1: 1
	c.Put(2, 2)
	c.Get(1) // 1: 2->1
	c.Get(1) // 1: 2->1
	c.Get(2) // 1: 2->1
	c.Get(2) // 1: 2->1
	c.Get(2) // 1: 2->1
	fmt.Println("GetMinFreq:", c.GetMinFreq())
	c.ShowFreqs()

	return
	c.Put(3, 3) // 2 失效

	c.ShowFreqs()
	fmt.Println(c.Get(1)) // -1
	fmt.Println(c.Get(2)) // -1
	fmt.Println(c.Get(3)) // 1: 3->1
	c.Put(4, 4)           // 1 失效 // 1: 4->3
	fmt.Println(c.Get(1)) // -1
	fmt.Println(c.Get(3)) // 1: 3->4
	fmt.Println(c.Get(4)) // 1: 4->3
}

func TestMain(m *testing.M) {
	c := Constructor(2)
	c.Put(1, 1) // 1: 1
	c.Put(2, 2)
	c.Get(1) // 1: 2->1
	c.Get(1) // 1: 2->1
	// c.Get(2) // 1: 2->1
	// c.Get(2) // 1: 2->1
	fmt.Println("GetMinFreq:", c.GetMinFreq())
	c.ShowFreqs()

	return
}

type LFUCache struct {
	Capacity int
	MinFreq  int
	Items    map[int]*Node
	Freqs    map[int]*List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		MinFreq:  0,
		Items:    make(map[int]*Node),
		Freqs:    make(map[int]*List),
	}
}

func (this *LFUCache) GetMinFreq() int {
	return this.MinFreq
}

func (this *LFUCache) ShowFreqs() {
	log.InfoDump(this.Freqs, "freqs")
	log.InfoDump(this.Freqs[this.MinFreq], "minFreq")
	log.InfoSeparator(this.Freqs[this.MinFreq])
}

func (this *LFUCache) Get(k int) int {
	// 未命中
	node, ok := this.Items[k]
	if !ok {
		return -1
	}

	// 命中
	this.Freqs[node.Freq].Remove(node) // 挪动到下一频率梯队
	node.Freq++
	if _, ok := this.Freqs[node.Freq]; !ok {
		this.Freqs[node.Freq] = NewList()
	}
	newNode := this.Freqs[node.Freq].Prepend(node)
	this.Items[k] = newNode
	// 注意判断最小梯队
	if this.Freqs[this.MinFreq].Size() == 0 {
		this.MinFreq++
	}
	return node.Val
}

func (this *LFUCache) Put(key, value int) {
	if this.Capacity == 0 {
		return
	}
	// 命中
	if val := this.Get(key); val != -1 {
		this.Items[key].Val = value
		return
	}

	// 缓存已满
	if this.Capacity == len(this.Items) {
		oldest := this.Freqs[this.MinFreq].Tail()
		this.Freqs[this.MinFreq].Remove(oldest)
		delete(this.Items, oldest.Key)
	}

	node := &Node{Key: key, Val: value, Freq: 1}
	this.Items[key] = node
	if _, ok := this.Freqs[1]; !ok {
		this.Freqs[1] = NewList()
	}
	this.Freqs[1].Prepend(node)
	this.MinFreq = 1
}

type Node struct {
	Key        int
	Val        int
	Freq       int
	Prev, Next *Node
}

type List struct {
	Head, tail *Node
	size       int
}

func NewList() *List {
	return new(List)
}

func (l *List) Prepend(node *Node) *Node {
	if l.Head == nil {
		l.Head = node
		l.tail = node
	} else {
		node.Prev = nil
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.size++
	return l.Head
}

func (l *List) Remove(node *Node) *Node {
	if node == nil {
		return nil
	}

	prev, next := node.Prev, node.Next
	if prev == nil {
		l.Head = next
	} else {
		prev.Next = next
	}

	if next == nil {
		l.tail = prev
	} else {
		next.Prev = prev
	}

	node.Prev, node.Next = nil, nil
	l.size--
	return node
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Tail() *Node {
	return l.tail
}
