package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	item := ListItem{value, nil, nil}
	if _, ok := lru.items[key]; ok {
		lru.items[key] = &item
		lru.queue.MoveToFront(&item)
		return true
	} else {
		if len(lru.items) != lru.capacity {
			lru.items[key] = &item 
			lru.queue.PushFront(item)
			lru.capacity++
		} else {
			lru.queue.Remove(&item)
			lru.queue.PushFront(item)
		}
		return false	
	}
}

func (lru *lruCache) Get(key Key) (interface {}, bool) {
	if item, ok := lru.items[key]; ok {
		lru.queue.MoveToFront(item)
		return item, ok
	} else {
		return nil, ok
	}
}

func (lru *lruCache) Clear(){
	lru.queue.ClearList()
	lru.items = make(map[Key]*ListItem, lru.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
