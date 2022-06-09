package lru

import (
	"errors"
	"testing"

	"github.com/pieterclaerhout/go-log"
)

func TestMain(m *testing.M) {
	c := Constructor(2)
	c.Put(1, 1) // 1: 1
	c.Put(2, 2)
	c.Get(1) // 1: 2->1
	c.Get(1) // 1: 2->1
	c.Get(1) // 1: 2->1
	c.Put(3, 3)
	c.Get(3)
	c.Get(3)
	// c.Get(2) // 1: 2->1
	// c.Get(2) // 1: 2->1
	// fmt.Println("GetMinFreq:", c.GetMinFreq())
	c.ShowFreqs()

	return
}

type LFUCache struct {
	Items   map[int]*Node
	MinFreq int
	Freqs   map[int]*List
	Cap     int
}

type List struct {
	HeadNode, TailNode *Node
	Size               int
}

type Node struct {
	Freq       int
	Key        int
	Val        int
	Next, Prev *Node
}

func Constructor(n int) *LFUCache {
	return &LFUCache{
		Cap:   n,
		Items: make(map[int]*Node),
		Freqs: make(map[int]*List),
	}
}

func NewList() *List {
	return new(List)
}

func (this *LFUCache) GetMinFreq() int {
	return this.MinFreq
}

func (this *LFUCache) ShowFreqs() {
	log.InfoDump(this.Freqs, "freqs")
	// log.InfoDump(this.Freqs[this.MinFreq], "minFreq")
	// log.InfoSeparator(this.Freqs[this.MinFreq])
}

func (l *LFUCache) Put(k int, v int) {
	// 存在
	// if _, ok := l.Items[k]; ok {
	// 	l.Items[k].Val = v
	// 	return
	// }
	res := l.Get(k)
	if res != -1 {
		l.Items[k].Val = v
		return
	}
	// 满额
	if l.Sizes() == l.Cap {
		oldest := l.Freqs[l.MinFreq].TailNode
		l.Freqs[l.MinFreq].Del(oldest)
	}
	// 不存在
	nod := &Node{Key: k, Val: v, Freq: 1}
	l.Items[k] = nod
	if _, ok := l.Freqs[1]; !ok {
		l.Freqs[1] = NewList()
	}
	l.Freqs[1].Add(nod)
	l.MinFreq = 1
}

func (l *LFUCache) Get(k int) (v int) {
	nod, ok := l.Items[k]
	if !ok {
		// err = errors.New("not found")
		return -1
	}
	// remove from old freq list
	l.Freqs[nod.Freq].Del(nod)
	// add to new freq list
	nod.Freq++
	if _, ok := l.Freqs[nod.Freq]; !ok {
		l.Freqs[nod.Freq] = NewList()
	}
	nod.Prev = nil
	nod.Next = nil
	l.Freqs[nod.Freq].Add(nod)
	// add freq
	if l.Freqs[l.MinFreq].Sizes() == 0 {
		l.MinFreq++
	}
	v = nod.Val
	return
}

func (l *LFUCache) Del(k int) (err error) {
	nod, ok := l.Items[k]
	if !ok {
		err = errors.New("not found")
		return
	}
	// remove from old freq list
	l.Freqs[nod.Freq].Del(nod)
	delete(l.Items, k)
	return
}

func (l *LFUCache) Sizes() int {
	return len(l.Items)
}

// Prepend \ add one node from head of the list
func (s *List) Add(nod *Node) {
	if s.HeadNode == nil {
		s.HeadNode = nod
		s.TailNode = nod
	} else {
		nod.Prev = nil
		nod.Next = s.HeadNode
		s.HeadNode.Prev = nod
		s.HeadNode = nod
	}
	s.Size++
}

func (s *List) Del(n *Node) {
	if n == nil {
		return
	}
	prev, next := n.Prev, n.Next
	if prev != nil {
		prev.Next = next
	} else {
		s.HeadNode = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		s.TailNode = prev
	}
	s.Size--
}

func (l *List) Sizes() int {
	return l.Size
}
