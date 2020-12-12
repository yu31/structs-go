package queue

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestQueueNew(t *testing.T) {
	capacity := 17
	q := New(capacity)

	require.NotNil(t, q)
	require.True(t, q.Empty())
	require.Equal(t, q.Len(), 0)
	require.Equal(t, q.Cap(), capacity)

	require.NotNil(t, q.items)
	require.Equal(t, len(q.items), capacity+1)
	require.Equal(t, q.cap, capacity+1)
	require.Equal(t, q.front, 0)
	require.Equal(t, q.behind, 0)

}

func TestQueue1(t *testing.T) {
	capacity := 1025
	q := Default()

	// test enqueue and make length of queue equal of capacity
	for i := 0; i < capacity; i++ {
		q.Push(i)
	}

	// queue already full
	require.Equal(t, q.front, 0)
	require.Equal(t, q.behind, capacity)

	require.Equal(t, q.Len(), capacity)

	// test dequeue and make queue empty
	for i := 0; i < capacity; i++ {
		item := q.Pop()
		require.NotNil(t, item)
		require.Equal(t, item, i)

		require.Nil(t, q.items[(q.front-1)%q.cap])
	}

	// queue is empty
	require.Equal(t, q.front, capacity)
	require.Equal(t, q.behind, capacity)

	require.True(t, q.Empty())
	require.Equal(t, q.Len(), 0)
	require.Nil(t, q.Pop())
}

func TestQueue2(t *testing.T) {
	capacity := 2
	q := New(capacity)

	p1 := unsafe.Pointer(&q.items[0])

	q.Push(1)
	q.Push(2)

	p2 := unsafe.Pointer(&q.items[0])
	require.Equal(t, p1, p2)

	q.Push(3)

	p3 := unsafe.Pointer(&q.items[0])
	require.NotEqual(t, p2, p3)

	require.Equal(t, q.Cap(), capacity*2)

	require.Equal(t, q.Pop(), 1)
	require.Equal(t, q.Pop(), 2)
	require.Equal(t, q.Pop(), 3)
}

func TestQueue3(t *testing.T) {
	capacity := 2
	q := New(capacity)

	q.Push(1)
	q.Push(2)
	require.Equal(t, q.Pop(), 1)
	q.Push(3)

	require.Greater(t, q.front, q.behind)

	q.Push(4)
	q.Push(5)
	q.Push(6)

	require.Equal(t, q.Pop(), 2)
	require.Equal(t, q.Pop(), 3)
	require.Equal(t, q.Pop(), 4)
	require.Equal(t, q.Pop(), 5)
	require.Equal(t, q.Pop(), 6)
}
