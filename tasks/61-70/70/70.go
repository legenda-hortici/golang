package main

import (
	"container/list"
	"fmt"
	"sync"
)

//===========================================================
//Задача 6
//1. Реализовать кеш. Для простоты считаем, что у нас бесконечная память и нам не нужно задумываться об удалении ключей из него.
//1. Почему использовал RWMutex, а не Mutex?
//2. Теперь представим что память не бесконечная. С какими проблемами столкнемся и как их решить?
//1. Какие есть алгоритмы выселения?
//3. Реализуй LRU
//===========================================================

// In-memory cache
// Нужно написать простую библиотеку in-memory cache.
// Реализация должна удовлетворять интерфейсу:

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
	Del(k string)
}

type entry struct {
	key, value string
}

type LRUCache struct {
	capacity int
	mu       sync.RWMutex
	items    map[string]*list.Element
	evict    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		evict:    list.New(),
	}
}

func (c *LRUCache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[k]; ok {
		// обновляем значение и двигаем в начало
		c.evict.MoveToFront(elem)
		elem.Value.(*entry).value = v
		return
	}

	// если места нет — удаляем последний
	if c.evict.Len() >= c.capacity {
		last := c.evict.Back()
		if last != nil {
			ent := last.Value.(*entry)
			delete(c.items, ent.key)
			c.evict.Remove(last)
		}
	}

	// вставляем новый
	ent := &entry{key: k, value: v}
	elem := c.evict.PushFront(ent)
	c.items[k] = elem
}

func (c *LRUCache) Get(k string) (string, bool) {
	c.mu.RLock()
	elem, ok := c.items[k]
	c.mu.RUnlock()
	if !ok {
		return "", false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.evict.MoveToFront(elem)

	return elem.Value.(*entry).value, true
}

func (c *LRUCache) Del(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[k]; ok {
		delete(c.items, k)
		c.evict.Remove(elem)
	}
}

func main() {
	cache := NewLRUCache(3)

	cache.Set("a", "1")
	cache.Set("b", "2")
	cache.Set("c", "3")

	cache.Get("a")      // a становится "свежим"
	cache.Set("d", "4") // вытеснит "b"

	if val, ok := cache.Get("b"); ok {
		fmt.Println("b =", val)
	} else {
		fmt.Println("b not found") // ожидаем это
	}

	if val, ok := cache.Get("a"); ok {
		fmt.Println("a =", val) // должно быть найдено
	}
}
