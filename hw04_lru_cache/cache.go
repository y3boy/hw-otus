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

// func (lru *lruCache) Set(key Key, value interface{}) bool {
// 	item, ok := lru.items[key]
// 	if ok {
// 		item.Value = value
// 		return true
// 	}

// 	if len(lru.items) == lru.capacity {
// 		lru.queue.Remove(lru.queue.Back())
// 	}
// 	ItemForSet := lru.queue.PushFront(value)
// 	lru.items[key] = ItemForSet
// 	lru.capacity++
// 	return false
// }

func (lru *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := lru.items[key]; ok {
		lru.queue.MoveToFront(item)
		item.Value = value
		return true
	} else if len(lru.items) == lru.capacity {
		lru.queue.Remove(lru.queue.Back())
	}
	lru.queue.PushFront(value)
	lru.items[key] = lru.queue.Front()
	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lru.items[key]
	if ok {
		lru.queue.MoveToFront(item)
		return lru.items[key].Value, ok
	}
	return nil, false
}

func (lru *lruCache) Clear() {
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
