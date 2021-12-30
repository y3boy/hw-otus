package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l list) PushFront(v interface{}) *ListItem {
	NewItemList := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.head == nil {
		l.head = NewItemList
		l.tail = NewItemList
	} else {
		NewItemList.Next = l.head
		l.head.Prev = NewItemList
		l.head = NewItemList
	}
	l.len++
	return NewItemList
}

func (l list) PushBack(v interface{}) *ListItem {
	NewItemList := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	if l.head == nil {
		l.head = NewItemList
		l.tail = NewItemList
	} else {
		l.tail.Next = NewItemList
		NewItemList.Prev = l.tail
		l.tail = NewItemList
	}
	l.len++
	return NewItemList
}

func (l list) Remove(i *ListItem) {
	fmt.Print(0)
}

func (l list) MoveToFront(i *ListItem) {
	fmt.Print(0)
}

func NewList() List {
	return new(list)
}
