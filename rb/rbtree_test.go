package rb

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/structs-go/container"
)

func recurseCalculateNodeHeight(n *treeNode) int {
	if n == nil {
		return 0
	}
	lh := recurseCalculateNodeHeight(n.left)
	rh := recurseCalculateNodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func checkBalance(t *testing.T, n *treeNode) {
	if n == nil {
		return
	}

	checkBalance(t, n.left)
	checkBalance(t, n.right)

	// Check the node color.
	if n.color == red {
		//if n.parent != nil {
		//	require.Equal(t, n.parent.color, black)
		//}
		require.NotNil(t, n.parent)
		require.Equal(t, n.parent.color, black)
		if n.left != nil && n.right != nil {
			require.Equal(t, n.left.color, black)
			require.Equal(t, n.right.color, black)
		} else {
			require.Nil(t, n.left)
			require.Nil(t, n.right)
		}
	} else {
		require.Equal(t, n.color, black)
	}

	if n.left != nil {
		require.Equal(t, n.key.Compare(n.left.key), 1)
	}
	if n.right != nil {
		require.Equal(t, n.key.Compare(n.right.key), -1)
	}

	// The height difference must not exceed one time.
	lh := recurseCalculateNodeHeight(n.left)
	rh := recurseCalculateNodeHeight(n.right)
	if lh > rh {
		require.LessOrEqual(t, lh-rh, lh)
	} else {
		require.LessOrEqual(t, rh-lh, rh)
	}
}

func TestNew(t *testing.T) {
	tr := New()
	require.NotNil(t, tr)
	require.Nil(t, tr.root)
	require.Equal(t, tr.len, 0)
}

func TestTree_createNode(t *testing.T) {
	tr := New()

	k := container.Int64(0xf)
	v := 1024

	n := tr.createNode(k, v, nil)
	require.NotNil(t, n)
	require.Equal(t, n.key.Compare(k), 0)
	require.Equal(t, n.value, v)
	require.Nil(t, n.left)
	require.Nil(t, n.right)
	require.Nil(t, n.parent)
	require.Equal(t, n.color, red)
}

func TestTree(t *testing.T) {
	tr := New()

	length := 257
	maxKey := length * 100
	keys := make([]container.Int64, length)

	for x := 0; x < 2; x++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int64(r.Intn(maxKey) + 1)
				if _, ok := tr.Insert(k, int64(k*2+1)); ok {
					keys[i] = k
					break
				}
			}
			require.Equal(t, tr.root.color, black)
			//checkBalance(t, tr.root)
			require.Equal(t, tr.Len(), i+1)
		}

		require.Equal(t, tr.Len(), length)

		// boundary
		for _, k := range []container.Int64{0, 0xfffffff} {
			_, ok := tr.Insert(k, k)
			require.True(t, ok)
			_, ok = tr.Insert(k, k)
			require.False(t, ok)
			require.NotNil(t, tr.Search(k))
			require.Equal(t, tr.Search(k).Value(), k)
			require.NotNil(t, tr.Delete(k))
			require.Nil(t, tr.Delete(k))
		}

		// search
		for i := 0; i < length; i++ {
			element := tr.Search(keys[i])
			require.NotNil(t, element)
			require.Equal(t, element.Key().Compare(keys[i]), 0)
			require.Equal(t, element.Value(), int64(keys[i]*2+1))
		}

		// delete
		for i := 0; i < length; i++ {
			ele := tr.Delete(keys[i])
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), keys[i])
			require.Equal(t, ele.Value(), int64(keys[i]*2+1))
			require.Nil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Search(keys[i]))

			if tr.root != nil {
				require.Equal(t, tr.root.color, black)
			}

			checkBalance(t, tr.root)
			require.Equal(t, tr.Len(), length-i-1)
		}

		require.Nil(t, tr.root)
		require.Equal(t, tr.Len(), 0)
	}
}
