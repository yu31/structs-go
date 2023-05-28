package minheap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/structs-go/container"
)

// checkCorrect check every item is less than of equal to the left child and right child.
func checkCorrect(t *testing.T, h *MinHeap) {
	// Check the index.
	for i := 0; i < h.len; i++ {
		require.Equal(t, h.items[i].index, i)
	}

	// Unused location should be nil.
	for i := h.len; i < h.cap; i++ {
		require.Nil(t, h.items[i])
	}

	for i := 0; i < (h.len-1)>>1; i++ {
		require.NotEqual(t, h.items[i].key.Compare(h.items[(i<<1)+1].key), 1)
		require.NotEqual(t, h.items[i].key.Compare(h.items[(i<<1)+2].key), 1)
	}
}

func TestNew(t *testing.T) {
	_ = Default()
	max := 17
	h := New(max)

	require.NotNil(t, h)
	require.True(t, h.Empty())
	require.Equal(t, h.len, 0)
	require.Equal(t, h.cap, max)
}

func TestMinHeap_Remove(t *testing.T) {
	h := Default()
	max := 1025
	maxKey := max * 100

	for i := 0; i < 2; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		// enqueue
		for i := 0; i < max; i++ {
			k := container.Int64(r.Intn(maxKey) + 1)

			item := h.Push(k, int(k*2+1))
			require.Equal(t, item.key.(container.Int64), k)
		}

		// remove first
		require.NotNil(t, h.Remove(0))
		// remove end
		require.NotNil(t, h.Remove(h.len-1))

		// Randomly removes part.
		part := 33
		for i := 2; i < 2+part; i++ {
			item := h.Remove(r.Intn(max - i))
			require.NotNil(t, item, fmt.Sprintf("del i: %d, len: %d", i, h.len))
			require.Equal(t, item.index, -1)

			require.Equal(t, h.Len(), max-i-1)
			require.Nil(t, h.items[h.len])

			checkCorrect(t, h)
		}

		// Randomly removes.
		for i := 2 + part; i < max; i++ {
			item := h.Remove(r.Intn(max - i))
			require.NotNil(t, item, fmt.Sprintf("del i: %d, len: %d", i, h.len))
			require.Equal(t, item.index, -1)

			require.Equal(t, h.Len(), max-i-1)
			require.Nil(t, h.items[h.len])
		}
		checkCorrect(t, h)
	}
}

func TestMinHeap(t *testing.T) {
	max := 1025
	maxKey := max * 100

	h := Default()

	for x := 0; x < 2; x++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		// enqueue
		for i := 0; i < max; i++ {
			k := container.Int64(r.Intn(maxKey) + 1)

			item := h.Push(k, int(k*2+1))
			require.Equal(t, item.key.(container.Int64), k)
			require.Equal(t, h.Len(), i+1)
		}

		checkCorrect(t, h)

		require.False(t, h.Empty())
		require.Equal(t, h.Len(), max)

		// dequeue and make queue empty.
		p1 := h.Peek()
		last := h.Pop()
		require.NotNil(t, last)
		require.Equal(t, last, p1)
		for i := 1; i < max; i++ {
			p1 := h.Peek()
			item := h.Pop()
			require.NotNil(t, item)
			require.Equal(t, item, p1)
			require.Equal(t, item.index, -1)

			require.True(t, last.key.Compare(item.key) != 1)
			require.Equal(t, item.value, int(item.key.(container.Int64))*2+1)

			require.Equal(t, h.Len(), max-i-1)
			require.Nil(t, h.items[h.len])

			last = item
		}

		checkCorrect(t, h)

		require.True(t, h.Empty())
		require.Equal(t, h.Len(), 0)
		require.Nil(t, h.Pop())
	}

}
