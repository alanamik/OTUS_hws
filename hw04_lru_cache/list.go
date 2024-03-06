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
	//List // Remove me after realization.
	// Place your code here.
	lenght int
	tail   *ListItem
	head   *ListItem
}

func NewList() List {
	l := &list{}
	l.head = nil
	l.tail = nil
	l.lenght = 0
	return l
}

func (l list) Len() int {
	return l.lenght
}

// - Len() int                           // длина списка
// - Front() *ListItem                   // первый элемент списка
// - Back() *ListItem                    // последний элемент списка
// - PushFront(v interface{}) *ListItem  // добавить значение в начало
// - PushBack(v interface{}) *ListItem   // добавить значение в конец
// - Remove(i *ListItem)                 // удалить элемент
// - MoveToFront(i *ListItem)            // переместить элемент в начало

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l list) PushFront(v interface{}) *ListItem {
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
	l.lenght++
	return newItem
}

func (l list) PushBack(v interface{}) *ListItem {
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
	l.lenght++
	return newItem
}

func (l list) Remove(i *ListItem) {
	//if l.head == nil {
	// When linked list is empty
	//	fmt.Print("Empty linked list")
	//}
	if l.head == i {
		// When remove head
		l.head = l.head.Next
		if l.head != nil {
			l.head.Next = nil
		} else {
			// When linked list empty after delete
			l.tail = nil
		}
	} else if l.tail == i {
		// When remove last node
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		} else {
			// Remove all nodes
			l.head = nil
		}
	} else {
		// When need to find deleted node
		var t *ListItem = l.head
		// Get remove node
		for t != nil && t != i {
			t = t.Next
		}
		if t == nil {
			// Node key not exist
			fmt.Println("Deleted node are not found")
		} else {
			// Separating deleted node
			// And combine next and previous node
			t.Prev.Next = t.Next
			if t.Next != nil {
				// When deleted intermediate nodes
				t.Next.Prev = t.Prev
			}
		}
	}
}

func (l list) MoveToFront(i *ListItem) {
	if l.tail == i {
		// When remove last node
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		} else {
			// Remove all nodes
			l.head = nil
		}
	} else {
		// When need to find deleted node
		var t *ListItem = l.head
		// Get remove node
		for t != nil && t != i {
			t = t.Next
		}
		if t == nil {
			// Node key not exist
			fmt.Println("Deleted node are not found")
		} else {
			// Separating deleted node
			// And combine next and previous node
			t.Prev.Next = t.Next
			if t.Next != nil {
				// When deleted intermediate nodes
				t.Next.Prev = t.Prev
			}
		}
	}
}
