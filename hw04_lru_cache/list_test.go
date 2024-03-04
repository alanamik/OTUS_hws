package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
	// added new tests
	t.Run("back elem to front", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)  // [1]
		l.PushBack(2)  // [1, 2]
		l.PushFront(0) // [0, 1, 2]

		del := l.PushBack(1) // [0, 1, 2, 1]
		l.MoveToFront(del)   // [1, 0, 1, 2]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{1, 0, 1, 2}, elems)
		require.Equal(t, 4, l.Len())
	})

	t.Run("front elem to front", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)  // [1]
		l.PushBack(2)  // [1, 2]
		l.PushFront(0) // [0, 1, 2]

		del := l.PushFront(3) // [3, 0, 1, 2]
		l.MoveToFront(del)    // [3, 0, 1, 2]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{3, 0, 1, 2}, elems)
		require.Equal(t, 4, l.Len())
	})

	t.Run("remove elems", func(t *testing.T) {
		l := NewList()
		d3 := l.PushBack(1)  // [1]
		d4 := l.PushBack(2)  // [1, 2]
		d1 := l.PushFront(0) // [0, 1, 2]
		d2 := l.PushFront(0) // [0, 0, 1, 2]
		require.Equal(t, 4, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{0, 0, 1, 2}, elems)

		l.Remove(d1)
		l.Remove(d2)
		l.Remove(d4)
		l.Remove(d3)
		require.Equal(t, 0, l.Len())
	})

	t.Run("remove all elems", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)  // [1]
		l.PushBack(2)  // [1, 2]
		l.PushFront(0) // [0, 1, 2]
		l.PushFront(0) // [0, 0, 1, 2]
		require.Equal(t, 4, l.Len())

		l.RemoveAll()

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{}, elems)
		require.Equal(t, 0, l.Len())
	})
}
