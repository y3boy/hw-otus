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
	ItemForSet := &ListItem{value, nil, nil}
	item, ok := lru.items[key]
	if ok {
		item.Value = value
		lru.queue.MoveToFront(item)
		return true
	}

	lru.items[key] = ItemForSet
	switch len(lru.items) {
	case lru.capacity:
		lru.queue.Remove(lru.queue.Back())
		lru.queue.PushFront(lru.items[key])
		return false
	default:
		lru.queue.PushFront(lru.items[key])
		lru.capacity++
		return false
	}
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := lru.items[key]; ok {
		item := item
		lru.queue.MoveToFront(item)
		return item.Value, ok
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
