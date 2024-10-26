package hw04lrucache

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
	size int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.head != nil {
		item.Next = l.head
		l.head.Prev = item
		l.head = item
	} else {
		l.tail = item
		l.head = item
	}
	l.size++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.tail != nil {
		item.Prev = l.tail
		l.tail.Next = item
		l.tail = item
	} else {
		l.head = item
		l.tail = item
	}
	l.size++
	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.tail {
		l.tail = i.Prev
	}
	i.Next = l.head
	i.Prev = nil
	if l.head != nil {
		l.head.Prev = i
	}
	l.head = i
	if l.tail == nil {
		l.tail = i
	}
}

func NewList() List {
	return new(list)
}
