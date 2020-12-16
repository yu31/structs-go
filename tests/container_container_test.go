package tests

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/avl"
	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/rb"
	"github.com/yu31/gostructs/skip"
)

var seeds map[container.Int]string

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 1024
	maxKey := length * 1000
	seeds = make(map[container.Int]string, length)

	for i := 0; i < length; i++ {
		for {
			k := container.Int(r.Intn(maxKey) + 1)
			if _, ok := seeds[k]; !ok {
				seeds[k] = strconv.Itoa(int(k*2 + 1))
				break
			}
		}
	}
}

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

	t.Run("tree", func(t *testing.T) {
		var box container.Tree
		_ = box

		box = bs.New()
		box = avl.New()
		box = rb.New()
	})
}

func TestContainer_Insert(t *testing.T) {
	process := func(box container.Container) {
		for k, v := range seeds {
			// The key not exists before, Insert was creates an new Element.
			ele, ok := box.Insert(k, v)
			require.True(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)

			// The key already exists, Insert was return the present Element.
			ele, ok = box.Insert(k, v+v)
			require.False(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		// Insert a exists key again.
		for k, v := range seeds {
			ele, ok := box.Insert(k, v+v)
			require.False(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainer_Delete(t *testing.T) {
	process := func(box container.Container) {
		// Try to delete key not exists.
		require.Nil(t, box.Delete(container.Int(0)))

		for k, v := range seeds {
			box.Insert(k, v)
		}

		for k, v := range seeds {
			ele := box.Delete(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
			require.Nil(t, box.Delete(k))
		}

		for k := range seeds {
			require.Nil(t, box.Delete(k))
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainer_Search(t *testing.T) {
	process := func(box container.Container) {
		// Try to search key not exists.
		require.Nil(t, box.Search(container.Int(0)))

		for k, v := range seeds {
			require.Nil(t, box.Search(k))
			box.Insert(k, v)
		}

		for k, v := range seeds {
			require.NotNil(t, box.Search(k))
			box.Insert(k, v+v)
		}

		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainer_Update(t *testing.T) {
	process := func(box container.Container) {
		// The updated key not exists.
		for k, v := range seeds {
			require.Nil(t, box.Update(k, v+v))
			require.Nil(t, box.Search(k))
		}

		// Insert data of seeds.
		for k, v := range seeds {
			_, ok := box.Insert(k, v)
			require.True(t, ok)
		}
		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		// Updated the value of key.
		for k, v := range seeds {
			ele := box.Update(k, v+v)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)

			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
		}
		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v+v)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainer_Replace(t *testing.T) {
	process := func(box container.Container) {
		// The key not exists, Replace same as the Insert
		for k, v := range seeds {
			require.Nil(t, box.Search(k))
			ele, ok := box.Replace(k, v)
			require.True(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		for k, v := range seeds {
			_, ok := box.Insert(k, v+v)
			require.False(t, ok)
		}

		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		require.Equal(t, box.Len(), len(seeds))

		// The key already exists, Replace same as the Update.
		for k, v := range seeds {
			ele, ok := box.Replace(k, v+v)
			require.False(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v, "key: %v", k)

			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
		}
		for k, v := range seeds {
			ele := box.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v+v)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainer_Len(t *testing.T) {
	process := func(box container.Container) {
		// Try to search key not exists.
		require.Equal(t, box.Len(), 0)
		var i int

		i = 1
		for k, v := range seeds {
			box.Insert(k, v)
			require.Equal(t, box.Len(), i)
			// Insert duplicates.
			box.Insert(k, v)
			require.Equal(t, box.Len(), i)
			i++
		}
		require.Equal(t, box.Len(), len(seeds))

		// Insert duplicates. No changed for length.
		for k, v := range seeds {
			box.Insert(k, v)
		}
		require.Equal(t, box.Len(), len(seeds))

		// Delete a not exists key.
		require.Nil(t, box.Delete(container.Int(10240)))
		require.Equal(t, box.Len(), len(seeds))

		// Update, No changed for length.
		for k, v := range seeds {
			box.Update(k, v+v)
			require.Equal(t, box.Len(), len(seeds))
		}
		require.Equal(t, box.Len(), len(seeds))

		// Delete and empty.
		i = len(seeds) - 1
		for k := range seeds {
			box.Delete(k)
			require.Equal(t, box.Len(), i)
			box.Delete(k)
			require.Equal(t, box.Len(), i)
			i--
		}
		require.Equal(t, box.Len(), 0)

		// Replace as Insert,
		i = 1
		for k, v := range seeds {
			box.Insert(k, v)
			require.Equal(t, box.Len(), i)
			i++
		}
		require.Equal(t, box.Len(), len(seeds))

		// Replace as Update, no changed.
		for k, v := range seeds {
			box.Replace(k, v+v)
			require.Equal(t, box.Len(), len(seeds))
		}
		require.Equal(t, box.Len(), len(seeds))

		//Delete and empty.
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

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}
