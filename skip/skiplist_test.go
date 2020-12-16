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
	_ = element

	element = &listNode{}
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

func TestList_Delete(t *testing.T) {
	sl := New()
	sl.Insert(container.Int(11), 1021)

	element := sl.Delete(container.Int(11))
	require.NotNil(t, element)
	require.Nil(t, element.(*listNode).next)
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
				if _, ok := sl.Insert(k, int64(k*2+1)); ok {
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
			_, ok := sl.Insert(k, k)
			require.True(t, ok)
			_, ok = sl.Insert(k, k)
			require.False(t, ok)
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
