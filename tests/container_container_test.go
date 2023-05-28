package tests

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/structs-go/container"
)

var seeds map[container.Int]string

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 2048
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

type ComplexKey struct {
	ID int64
}

func (k1 *ComplexKey) Compare(target container.Comparator) int {
	k2 := target.(*ComplexKey)
	if k1.ID < k2.ID {
		return -1
	}
	if k1.ID > k2.ID {
		return 1
	}
	return 0
}

type ComplexValue struct {
	Name string
}

func TestContainer_Insert(t *testing.T) {
	process := func(ctr container.Container) {
		for k, v := range seeds {
			// The key not exists before, Insert was creates an new element.
			ele, ok := ctr.Insert(k, v)
			require.True(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v)

			// The key already exists, Insert was return the present element.
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

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}

	boundary := func(ctr container.Container) {
		// Test complex type of key and value.
		ck := &ComplexKey{ID: 1024}
		cv := &ComplexValue{Name: "developer"}
		actual, ok := ctr.Insert(ck, cv)
		require.True(t, ok)
		require.Equal(t, actual.Key(), ck)
		require.Equal(t, actual.Value(), cv)

		actual, ok = ctr.Insert(ck, cv)
		require.False(t, ok)
		require.Equal(t, actual.Key(), ck)
		require.Equal(t, actual.Value(), cv)
	}

	// Boundary test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_boundary", func(t *testing.T) {
			boundary(f())
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
				require.True(t, n.Left() == nil)
				require.True(t, n.Right() == nil)
			}
			require.Nil(t, ctr.Delete(k))
			require.True(t, ctr.Delete(k) == nil)
		}

		for k := range seeds {
			require.Nil(t, ctr.Delete(k))
			require.True(t, ctr.Delete(k) == nil)
		}
	}

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}

	boundary := func(ctr container.Container) {
		// Test complex type of key and value.
		ck := &ComplexKey{ID: 1024}
		cv := &ComplexValue{Name: "developer"}
		ctr.Insert(ck, cv)

		ele := ctr.Delete(ck)
		require.NotNil(t, ele)
		require.Equal(t, ele.Key(), ck)
		require.Equal(t, ele.Value(), cv)

		require.Nil(t, ctr.Delete(ck))
		require.True(t, ctr.Delete(ck) == nil)
	}

	// Boundary test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_boundary", func(t *testing.T) {
			boundary(f())
		})
	}
}

func TestContainer_Search(t *testing.T) {
	process := func(ctr container.Container) {
		// Try to search key not exists.
		require.Nil(t, ctr.Search(container.Int(0)))
		require.True(t, ctr.Search(container.Int(0)) == nil)

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

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}

	boundary := func(ctr container.Container) {
		// Test complex type of key and value.
		ck := &ComplexKey{ID: 1024}
		cv := &ComplexValue{Name: "developer"}
		ctr.Insert(ck, cv)

		require.NotNil(t, ctr.Search(ck))
		require.Equal(t, ctr.Search(ck).Key(), ck)
		require.Equal(t, ctr.Search(ck).Value(), cv)

		require.Nil(t, ctr.Search(&ComplexKey{ID: 2048}))
		require.True(t, ctr.Search(&ComplexKey{ID: 2048}) == nil)
	}

	// Boundary test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_boundary", func(t *testing.T) {
			boundary(f())
		})
	}
}

func TestContainer_Update(t *testing.T) {
	process := func(ctr container.Container) {
		// The updated key not exists.
		for k, v := range seeds {
			require.Nil(t, ctr.Update(k, v+v))
			require.Nil(t, ctr.Search(k))
			require.True(t, ctr.Update(k, v+v) == nil)
			require.True(t, ctr.Search(k) == nil)
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
				require.True(t, n.Left() == nil)
				require.True(t, n.Right() == nil)
			}
		}
		for k, v := range seeds {
			ele := ctr.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v+v)
		}
	}

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}

	boundary := func(ctr container.Container) {
		// Test complex type of key and value.
		ck := &ComplexKey{ID: 1024}
		cv := &ComplexValue{Name: "developer"}
		ctr.Insert(ck, cv)

		cv2 := &ComplexValue{Name: "update-v2"}
		ele := ctr.Update(ck, cv2)
		require.NotNil(t, ele)
		require.Equal(t, ele.Key(), ck)
		require.Equal(t, ele.Value(), cv)

		require.NotNil(t, ctr.Search(ck))
		require.Equal(t, ctr.Search(ck).Key(), ck)
		require.Equal(t, ctr.Search(ck).Value(), cv2)

		require.Nil(t, ctr.Update(&ComplexKey{ID: 2048}, nil))
		require.True(t, ctr.Update(&ComplexKey{ID: 2048}, nil) == nil)
	}

	// Boundary test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_boundary", func(t *testing.T) {
			boundary(f())
		})
	}
}

func TestContainer_Upsert(t *testing.T) {
	process := func(ctr container.Container) {
		// The key not exists, Upsert same as the Insert
		for k, v := range seeds {
			require.Nil(t, ctr.Search(k))
			require.True(t, ctr.Search(k) == nil)
			ele, ok := ctr.Upsert(k, v)
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

		// The key already exists, Upsert same as the Update.
		for k, v := range seeds {
			ele, ok := ctr.Upsert(k, v+v)
			require.False(t, ok)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v, "key: %v", k)

			if n, ok := ele.(container.TreeNode); ok {
				require.Nil(t, n.Left())
				require.Nil(t, n.Right())
				require.True(t, n.Left() == nil)
				require.True(t, n.Right() == nil)
			}
		}
		for k, v := range seeds {
			ele := ctr.Search(k)
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), k)
			require.Equal(t, ele.Value(), v+v)
		}
	}

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}

	boundary := func(ctr container.Container) {
		// Test complex type of key and value.
		ck := &ComplexKey{ID: 1024}
		cv := &ComplexValue{Name: "developer"}
		ele, ok := ctr.Upsert(ck, cv)
		require.True(t, ok)
		require.NotNil(t, ele)
		require.Equal(t, ele.Key(), ck)
		require.Equal(t, ele.Value(), cv)

		cv2 := &ComplexValue{Name: "update-v2"}
		ele, ok = ctr.Upsert(ck, cv2)
		require.False(t, ok)
		require.NotNil(t, ele)
		require.Equal(t, ele.Key(), ck)
		require.Equal(t, ele.Value(), cv)

		require.NotNil(t, ctr.Search(ck))
		require.Equal(t, ctr.Search(ck).Key(), ck)
		require.Equal(t, ctr.Search(ck).Value(), cv2)
	}

	// Boundary test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_boundary", func(t *testing.T) {
			boundary(f())
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

		// Upsert as Insert,
		i = 1
		for k, v := range seeds {
			ctr.Insert(k, v)
			require.Equal(t, ctr.Len(), i)
			i++
		}
		require.Equal(t, ctr.Len(), len(seeds))

		// Upsert as Update, no changed.
		for k, v := range seeds {
			ctr.Upsert(k, v+v)
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

	// Base test for all container implementation.
	for name, f := range containers {
		t.Run(name+"_base", func(t *testing.T) {
			process(f())
		})
	}
}
