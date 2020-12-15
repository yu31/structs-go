package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/avl"
	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/rb"
	"github.com/yu31/gostructs/skip"
)

func TestContainer_Interface(t *testing.T) {
	// Ensure the bs/avl/rb/skip are implements the container.Container.
	t.Run("container", func(t *testing.T) {
		var box container.Container
		_ = box

		box = bs.New()
		box = avl.New()
		box = rb.New()
		box = skip.New()
	})

	t.Run("iterator", func(t *testing.T) {
		var it container.Iterator
		_ = it

		it = bs.NewIterator(nil, nil, nil)
	})
}

var seeds = map[container.Int]string{
	11: "1024",
	22: "1025",
	33: "1026",
	44: "1027",
	55: "1028",
	66: "1029",
}

func TestContainer_Insert(t *testing.T) {
	process := func(box container.Container) {
		for k, v := range seeds {
			ele := box.Insert(k, v)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
			// Test inserts a exists key.
			require.Nil(t, box.Insert(k, v))
		}
	}

	t.Run("bstree", func(t *testing.T) {
		process(bs.New())
	})
	t.Run("avtree", func(t *testing.T) {
		process(avl.New())
	})
	t.Run("rbtree", func(t *testing.T) {
		process(rb.New())
	})
	t.Run("skiplist", func(t *testing.T) {
		process(skip.New())
	})
}

func TestContainer_Delete(t *testing.T) {
	process := func(box container.Container) {
		// Try to delete key not exists.
		require.Nil(t, box.Delete(container.Int(10240)))

		for k, v := range seeds {
			box.Insert(k, v)
		}

		for k, v := range seeds {
			ele := box.Delete(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
			if n, ok := ele.(bs.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
			require.Nil(t, box.Delete(k))
		}
	}

	t.Run("bstree", func(t *testing.T) {
		process(bs.New())
	})
	t.Run("avtree", func(t *testing.T) {
		process(avl.New())
	})
	t.Run("rbtree", func(t *testing.T) {
		process(rb.New())
	})
	t.Run("skiplist", func(t *testing.T) {
		process(skip.New())
	})
}

func TestContainer_Search(t *testing.T) {
	process := func(box container.Container) {
		// Try to search key not exists.
		require.Nil(t, box.Search(container.Int(10240)))

		for k, v := range seeds {
			box.Insert(k, v)
		}

		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}
	}

	t.Run("bstree", func(t *testing.T) {
		process(bs.New())
	})
	t.Run("avtree", func(t *testing.T) {
		process(avl.New())
	})
	t.Run("rbtree", func(t *testing.T) {
		process(rb.New())
	})
	t.Run("skiplist", func(t *testing.T) {
		process(skip.New())
	})
}

// TODO: test with Update and Replace.
func TestContainer_Len(t *testing.T) {
	process := func(box container.Container) {
		// Try to search key not exists.
		require.Equal(t, box.Len(), 0)

		i := 1
		for k, v := range seeds {
			box.Insert(k, v)
			require.Equal(t, box.Len(), i)
			// Insert duplicates.
			box.Insert(k, v)
			require.Equal(t, box.Len(), i)
			i++
		}
		require.Equal(t, box.Len(), len(seeds))

		// Insert duplicates.
		for k, v := range seeds {
			box.Insert(k, v)
		}
		require.Equal(t, box.Len(), len(seeds))

		// Delete a not exists key.
		require.Nil(t, box.Delete(container.Int(10240)))
		require.Equal(t, box.Len(), len(seeds))

		// Delete
		i = len(seeds) - 1
		for k := range seeds {
			box.Delete(k)
			require.Equal(t, box.Len(), i)
			box.Delete(k)
			require.Equal(t, box.Len(), i)
			i--
		}
		require.Equal(t, box.Len(), 0)
	}

	t.Run("bstree", func(t *testing.T) {
		process(bs.New())
	})
	t.Run("avtree", func(t *testing.T) {
		process(avl.New())
	})
	t.Run("rbtree", func(t *testing.T) {
		process(rb.New())
	})
	t.Run("skiplist", func(t *testing.T) {
		process(skip.New())
	})
}
