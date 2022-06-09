package lru

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"testing"
	"unsafe"
)

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}

func OnEvictedTest(key string, value Value) {
	log.Println(key, value)
	return
}

type myint struct {
	Num interface{}
}

func (t myint) Len() int {
	// var aa int64 = 10
	n := unsafe.Sizeof(t)
	str := fmt.Sprintf("%d", n)
	res, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	// log.Println(t, res)
	return int(res)
	// str := strconv.FormatInt(int64(t.Num), 10)
	// return len(str)
}

func TestLru(*testing.T) {
	log.SetFlags(9)
	l := New(10, OnEvictedTest)
	l.Add("11", myint{Num: 11})
	l.Add("113", myint{Num: 113})
	l.Add("1134", myint{Num: 1134})
	l.Add("11345", myint{Num: 1134582913227892719})
	v, ok := l.Get("1134")
	if ok {
		log.Println(v)
	}
}
