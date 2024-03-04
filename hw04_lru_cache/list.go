package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	RemoveAll()
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	tail   *ListItem
	head   *ListItem
}

func NewList() *list {
	l := &list{}
	l.head = nil
	l.tail = nil
	l.length = 0
	return l
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	value := v.(int)
	newItem := &ListItem{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.Next = l.head
		l.head.Prev = newItem
		l.head = newItem
	}
	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	value := v.(int)
	newItem := &ListItem{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		currentItem := l.head
		for currentItem.Next != nil {
			currentItem = currentItem.Next
		}
		newItem.Prev = currentItem
		currentItem.Next = newItem
		l.tail = newItem
	}
	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if l.head == i {
		l.head = l.head.Next
		if l.head != nil {
			l.head.Prev = nil
		}
		l.length--
		return
	}
	if l.tail == i {
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		}
		l.length--
		return
	}
	t := l.head.Next
	for t != nil && t != i {
		t = t.Next
	}
	if t == nil {
		return
	}
	t.Prev.Next = t.Next
	if t.Next != nil {
		t.Next.Prev = t.Prev
	}
	l.length--
}

func (l *list) RemoveAll() {
	tt := l.head.Next
	var t *ListItem
	for tt != l.tail && tt != nil {
		t = tt.Next
		l.Remove(tt)
		tt = t
	}
	l.Remove(l.head)
	l.Remove(l.tail)
}

func (l *list) MoveToFront(i *ListItem) {
	if l.tail == i {
		l.PushFront(i.Value)
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		}
		l.length--
		return
	}
	if l.head == i {
		return
	}
	t := l.head.Next
	for t != nil && t != i {
		t = t.Next
	}
	if t == nil {
		return
	}
	if t.Prev == nil {
		t.Prev = t.Next
	} else {
		t.Prev.Next = t.Next
	}
	if t.Next != nil {
		t.Next.Prev = t.Prev
	}
}
