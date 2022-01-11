package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	ClearList() // For cleaning
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

func (l *list) PushFront(v interface{}) *ListItem {
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

func (l *list) PushBack(v interface{}) *ListItem {
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

// Создал эту функцию для себя
// func (l list) Find(i *ListItem) (*ListItem, bool) {
// 	found := false
// 	var i *ListItem = nil
// 	for n := l.Front(); n != nil && !found; n = n.Next{
// 		if n.Value == i.Value {
// 			found = true
// 			i = n
// 		}
// 	}
// 	return i, found
// }

func (l *list) Remove(i *ListItem) {
	switch i {
	case l.tail:
		l.tail = i.Prev
		l.tail.Next = nil
		i.Prev.Next = nil
		i.Prev = nil
	case l.head:
		l.head = i.Next
		i.Next.Prev = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) ClearList() {
	l.head = nil
	l.tail = nil
}

func NewList() List {
	return new(list)
}
