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
	"github.com/yu31/gostructs/internal/tree"
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
		var ctr container.Container
		_ = ctr

		ctr = bs.New()
		ctr = avl.New()
		ctr = rb.New()
		ctr = skip.New()
	})

	t.Run("iterator", func(t *testing.T) {
		var it container.Iterator
		_ = it
		it = tree.NewIterator(nil, nil, nil)
	})

	t.Run("tree", func(t *testing.T) {
		var ctr container.Tree
		_ = ctr

		ctr = bs.New()
		ctr = avl.New()
		ctr = rb.New()
	})
}

func TestContainer_Insert(t *testing.T) {
	process := func(ctr container.Container) {
		for k, v := range seeds {
			// The key not exists before, Insert was creates an new Element.
			ele, ok := ctr.Insert(k, v)
			require.True(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)

			// The key already exists, Insert was return the present Element.
			ele, ok = ctr.Insert(k, v+v)
			require.False(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		// Insert a exists key again.
		for k, v := range seeds {
			ele, ok := ctr.Insert(k, v+v)
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
	process := func(ctr container.Container) {
		// Try to delete key not exists.
		require.Nil(t, ctr.Delete(container.Int(0)))

		for k, v := range seeds {
			ctr.Insert(k, v)
		}

		for k, v := range seeds {
			ele := ctr.Delete(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
			require.Nil(t, ctr.Delete(k))
		}

		for k := range seeds {
			require.Nil(t, ctr.Delete(k))
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
	process := func(ctr container.Container) {
		// Try to search key not exists.
		require.Nil(t, ctr.Search(container.Int(0)))

		for k, v := range seeds {
			require.Nil(t, ctr.Search(k))
			ctr.Insert(k, v)
		}

		for k, v := range seeds {
			require.NotNil(t, ctr.Search(k))
			ctr.Insert(k, v+v)
		}

		for k, v := range seeds {
			ele := ctr.Search(k)
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
	process := func(ctr container.Container) {
		// The updated key not exists.
		for k, v := range seeds {
			require.Nil(t, ctr.Update(k, v+v))
			require.Nil(t, ctr.Search(k))
		}

		// Insert data of seeds.
		for k, v := range seeds {
			_, ok := ctr.Insert(k, v)
			require.True(t, ok)
		}
		for k, v := range seeds {
			ele := ctr.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		// Updated the value of key.
		for k, v := range seeds {
			ele := ctr.Update(k, v+v)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)

			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
			}
		}
		for k, v := range seeds {
			ele := ctr.Search(k)
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
	process := func(ctr container.Container) {
		// The key not exists, Replace same as the Insert
		for k, v := range seeds {
			require.Nil(t, ctr.Search(k))
			ele, ok := ctr.Replace(k, v)
			require.True(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		for k, v := range seeds {
			_, ok := ctr.Insert(k, v+v)
			require.False(t, ok)
		}

		for k, v := range seeds {
			ele := ctr.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)
		}

		require.Equal(t, ctr.Len(), len(seeds))

		// The key already exists, Replace same as the Update.
		for k, v := range seeds {
			ele, ok := ctr.Replace(k, v+v)
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
			ele := ctr.Search(k)
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
	process := func(ctr container.Container) {
		// Try to search key not exists.
		require.Equal(t, ctr.Len(), 0)
		var i int
		i = 1
		for k, v := range seeds {
			ctr.Insert(k, v)
			require.Equal(t, ctr.Len(), i)
			// Insert duplicates.
			ctr.Insert(k, v)
			require.Equal(t, ctr.Len(), i)
			i++
		}
		require.Equal(t, ctr.Len(), len(seeds))

		// Insert duplicates. No changed for length.
		for k, v := range seeds {
			ctr.Insert(k, v)
		}
		require.Equal(t, ctr.Len(), len(seeds))

		// Delete a not exists key.
		require.Nil(t, ctr.Delete(container.Int(10240)))
		require.Equal(t, ctr.Len(), len(seeds))

		// Update, No changed for length.
		for k, v := range seeds {
			ctr.Update(k, v+v)
			require.Equal(t, ctr.Len(), len(seeds))
		}
		require.Equal(t, ctr.Len(), len(seeds))

		// Delete and empty.
		i = len(seeds) - 1
		for k := range seeds {
			ctr.Delete(k)
			require.Equal(t, ctr.Len(), i)
			ctr.Delete(k)
			require.Equal(t, ctr.Len(), i)
			i--
		}
		require.Equal(t, ctr.Len(), 0)

		// Replace as Insert,
		i = 1
		for k, v := range seeds {
			ctr.Insert(k, v)
			require.Equal(t, ctr.Len(), i)
			i++
		}
		require.Equal(t, ctr.Len(), len(seeds))

		// Replace as Update, no changed.
		for k, v := range seeds {
			ctr.Replace(k, v+v)
			require.Equal(t, ctr.Len(), len(seeds))
		}
		require.Equal(t, ctr.Len(), len(seeds))

		//Delete and empty.
		i = len(seeds) - 1
		for k := range seeds {
			ctr.Delete(k)
			require.Equal(t, ctr.Len(), i)
			ctr.Delete(k)
			require.Equal(t, ctr.Len(), i)
			i--
		}
		require.Equal(t, ctr.Len(), 0)
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}
