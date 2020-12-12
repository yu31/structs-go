package skip

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

func output(sl *List) {
	fmt.Println("---------- output ---------")
	for i := 0; i <= sl.level; i++ {
		fmt.Printf("Level <%d, %d> | ", i, sl.lens[i])
		p := sl.head.next[i]
		for p != nil {
			fmt.Printf("%d -> ", p.key)
			p = p.next[i]
		}
		fmt.Printf("\n")
	}
	fmt.Println("---------- output ---------")
}

func checkCorrect(t *testing.T, sl *List) {
	for i := 0; i <= sl.level; i++ {
		p := sl.head.next[i]
		for p != nil && p.next[i] != nil {
			require.Equal(t, p.key.Compare(p.next[i].key), -1)
			p = p.next[i]
		}
	}
}

func Test_Interface(t *testing.T) {
	// Ensure the interface is implemented.
	var element container.Element
	var ct container.Container
	var it container.Iterator

	element = &listNode{}
	_ = element
	ct = New()
	_ = ct
	it = ct.Iter(nil, nil)
	_ = it
}

func TestNew(t *testing.T) {
	sl := New()

	require.NotNil(t, sl)
	require.Equal(t, sl.level, 0)
	require.NotNil(t, sl.r)
	require.NotNil(t, sl.head)
	require.Equal(t, len(sl.lens), maxLevel+1)

	for i := 0; i <= maxLevel; i++ {
		require.Nil(t, sl.head.next[i])
		require.Equal(t, sl.lens[i], 0)
	}
}

func TestList_createNode(t *testing.T) {
	sl := New()

	k := container.Int64(0xf)
	v := 1024
	level := 10

	n := sl.createNode(k, v, level)
	require.NotNil(t, n)
	require.Equal(t, n.key.Compare(k), 0)
	require.Equal(t, n.value, v)
	require.Equal(t, len(n.next), level+1)

	for i := 0; i <= level; i++ {
		require.Nil(t, sl.head.next[i])
		require.Equal(t, sl.lens[i], 0)
	}
}

func TestList_Insert(t *testing.T) {
	sl := New()

	require.NotNil(t, sl.Insert(container.Int(11), 1024))
	require.Nil(t, sl.Insert(container.Int(11), 1023))
	require.NotNil(t, sl.Insert(container.Int(33), nil))
	require.Nil(t, sl.Insert(container.Int(33), nil))
	require.NotNil(t, sl.Insert(container.Int(22), nil))
	require.Nil(t, sl.Insert(container.Int(22), nil))
}

func TestList_Delete(t *testing.T) {
	sl := New()
	require.NotNil(t, sl.Insert(container.Int(11), 1021))
	require.NotNil(t, sl.Insert(container.Int(22), 1022))
	require.NotNil(t, sl.Insert(container.Int(33), 1023))

	element := sl.Delete(container.Int(11))
	require.NotNil(t, element)
	require.Equal(t, element.Key().Compare(container.Int(11)), 0)
	require.Equal(t, element.Value(), 1021)
	require.Nil(t, element.(*listNode).next)
	require.Nil(t, sl.Delete(container.Int(11)))

	require.NotNil(t, sl.Delete(container.Int(22)))
	require.Nil(t, sl.Delete(container.Int(22)))
	require.NotNil(t, sl.Delete(container.Int(33)))
	require.Nil(t, sl.Delete(container.Int(33)))

	// Try to delete key not exists.
	require.Nil(t, sl.Delete(container.Int(1024)))
}

func TestList_Search(t *testing.T) {
	sl := New()
	require.NotNil(t, sl.Insert(container.Int(11), 1021))
	require.NotNil(t, sl.Insert(container.Int(22), 1022))
	require.NotNil(t, sl.Insert(container.Int(33), 1023))

	require.Equal(t, sl.Search(container.Int(11)).Key().Compare(container.Int(11)), 0)
	require.Equal(t, sl.Search(container.Int(11)).Value(), 1021)
	require.Equal(t, sl.Search(container.Int(22)).Key().Compare(container.Int(22)), 0)
	require.Equal(t, sl.Search(container.Int(22)).Value(), 1022)
	require.Equal(t, sl.Search(container.Int(33)).Key().Compare(container.Int(33)), 0)
	require.Equal(t, sl.Search(container.Int(33)).Value(), 1023)

	// Try to search key not exists.
	require.Nil(t, sl.Search(container.Int(1024)))
}

func TestList(t *testing.T) {
	sl := New()

	length := 257
	maxKey := length * 100
	keys := make([]container.Int64, length)

	for x := 0; x < 2; x++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int64(r.Intn(maxKey) + 1)
				if sl.Insert(k, int64(k*2+1)) != nil {
					require.Nil(t, sl.Insert(k, int64(k*2+1)))
					keys[i] = k
					break
				}
			}
			checkCorrect(t, sl)
			require.Equal(t, sl.Len(), i+1)
		}

		require.Equal(t, sl.Len(), length)
		require.LessOrEqual(t, sl.level, maxLevel)

		// boundary
		for _, k := range []container.Int64{0, 0xfffffff} {
			require.NotNil(t, sl.Insert(k, k))
			require.Nil(t, sl.Insert(k, k))
			require.NotNil(t, sl.Search(k))
			require.Equal(t, sl.Search(k).Value(), k)
			require.NotNil(t, sl.Delete(k))
			require.Nil(t, sl.Delete(k))
		}

		// search
		for i := 0; i < length; i++ {
			element := sl.Search(keys[i])
			require.NotNil(t, element)
			require.Equal(t, element.Value(), int64(keys[i]*2+1))
		}

		// delete
		for i := 0; i < length; i++ {
			require.NotNil(t, sl.Delete(keys[i]))
			require.Nil(t, sl.Delete(keys[i]))

			checkCorrect(t, sl)
			require.Equal(t, sl.Len(), length-i-1)
		}

		require.Equal(t, sl.Len(), 0)
		require.Equal(t, sl.level, 0)

		output(sl)

		for i := 0; i <= maxLevel; i++ {
			require.Nil(t, sl.head.next[i])
			require.Equal(t, sl.lens[i], 0)
		}
	}
}
