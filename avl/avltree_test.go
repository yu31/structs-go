package avl

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

func checkBalance(t *testing.T, tr *Tree, n *treeNode) {
	if n == nil {
		return
	}

	checkBalance(t, tr, n.left)
	checkBalance(t, tr, n.right)

	// Check the node height calculate.
	require.Equal(t, tr.nodeHeight(n), recurseCalculateNodeHeight(n))

	if n.left != nil {
		require.Equal(t, n.key.Compare(n.left.key), 1)
	}
	if n.right != nil {
		require.Equal(t, n.key.Compare(n.right.key), -1)
	}

	// The height difference cannot exceed 1 in AVL Tree.
	lh := tr.nodeHeight(n.left)
	rh := tr.nodeHeight(n.right)
	if lh > rh {
		require.Equal(t, lh-rh, 1)
	} else if lh < rh {
		require.Equal(t, rh-lh, 1)
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

	n := tr.createNode(k, v)
	require.NotNil(t, n)
	require.Equal(t, n.key.Compare(k), 0)
	require.Equal(t, n.value, v)
	require.Nil(t, n.left)
	require.Nil(t, n.right)
	require.Equal(t, n.height, 1)
}

func TestTree(t *testing.T) {
	tr := New()

	length := 257
	maxKey := length * 100
	keys := make([]container.Int, length)

	for x := 0; x < 2; x++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int(r.Intn(maxKey) + 1)
				if _, ok := tr.Insert(k, int64(k*2+1)); ok {
					keys[i] = k
					break
				}
			}
			checkBalance(t, tr, tr.root)
			require.Equal(t, tr.Len(), i+1)
		}

		require.Equal(t, tr.Len(), length)

		// boundary
		for _, k := range []container.Int{0, 0xfffffff} {
			_, ok := tr.Insert(k, k)
			require.True(t, ok)
			_, ok = tr.Insert(k, k)
			require.False(t, ok)
			require.NotNil(t, tr.Search(k))
			require.Equal(t, tr.Search(k).Value(), k)
			require.NotNil(t, tr.Delete(k))
			require.Nil(t, tr.Delete(k))
		}

		// get
		for i := 0; i < length; i++ {
			element := tr.Search(keys[i])
			require.NotNil(t, element)
			require.Equal(t, element.Key().Compare(keys[i]), 0)
			require.Equal(t, element.Value(), int64(keys[i]*2+1))
		}

		// delete
		for i := 0; i < length; i++ {
			require.NotNil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Search(keys[i]))

			checkBalance(t, tr, tr.root)
			require.Equal(t, tr.Len(), length-i-1)
		}

		require.Nil(t, tr.root)
		require.Equal(t, tr.Len(), 0)
	}
}
