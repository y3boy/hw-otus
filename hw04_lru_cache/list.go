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
// 	var ForDelete *ListItem = nil
// 	for n := l.Front(); n != nil && !found; n = n.Next{
// 		if n.Value == i.Value {
// 			found = true
// 			ForDelete = n
// 		}
// 	}
// 	return ForDelete, found
// }

func (l *list) Remove(i *ListItem) {
	// if ForDelete, found := l.Find(i); found {
	ForDelete := i
	switch ForDelete {
	case l.tail:
		ForDelete.Prev.Next = nil
		l.tail = ForDelete.Prev
	case l.head:
		ForDelete.Next.Prev = nil
		l.head = ForDelete.Next
	default:
		ForDelete.Prev.Next = ForDelete.Next
		ForDelete.Next.Prev = ForDelete.Prev
		l.len--
	}
	// }
}

func (l *list) MoveToFront(i *ListItem) {
	// if ForMove, found := l.Find(i); found {
	ForMove := i
	switch ForMove {
	case l.tail:
		l.tail = ForMove.Prev
		ForMove.Prev.Next = nil
		ForMove.Next = l.head
		l.head = ForMove
		if l.tail.Prev == nil {
			l.tail.Prev = l.head
		}
	case l.head:
		break
	default:
		ForMove.Prev.Next = ForMove.Next
		ForMove.Next.Prev = ForMove.Prev
		ForMove.Next = l.head
		l.head = ForMove
	}
	// }
}

func (l *list) ClearList() {
	l.head = nil
	l.tail = nil
}

func NewList() List {
	return new(list)
}
