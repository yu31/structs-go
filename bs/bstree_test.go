package bs

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/structs-go/container"
)

// Iteration with recursion.
func rangeRecursion(root *treeNode, start container.Key, boundary container.Key, f func(node *treeNode) bool) {
	if root == nil {
		return
	}

	if start != nil && root.key.Compare(start) == -1 {
		rangeRecursion(root.right, start, boundary, f)
	} else if boundary != nil && root.key.Compare(boundary) != -1 {
		rangeRecursion(root.left, start, boundary, f)
	} else {
		// start <= node <= boundary
		rangeRecursion(root.left, start, boundary, f)
		if !f(root) {
			return
		}
		rangeRecursion(root.right, start, boundary, f)
	}
}

func checkCorrect(t *testing.T, n *treeNode) {
	if n == nil {
		return
	}
	checkCorrect(t, n.left)
	checkCorrect(t, n.right)

	if n.left != nil {
		require.Equal(t, n.key.Compare(n.left.key), 1)
	}
	if n.right != nil {
		require.Equal(t, n.key.Compare(n.right.key), -1)
	}
}

func TestNew(t *testing.T) {
	tr := New()
	require.NotNil(t, tr)
	require.Nil(t, tr.root)
	require.Equal(t, tr.Len(), 0)
	_ = rangeRecursion
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
			checkCorrect(t, tr.root)
			require.Equal(t, tr.Len(), i+1)
		}

		require.Equal(t, tr.Len(), length)

		// search
		for i := 0; i < length; i++ {
			ele := tr.Search(keys[i])
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), keys[i])
			require.Equal(t, ele.Value(), int64(keys[i]*2+1))
		}

		// delete
		for i := 0; i < length; i++ {
			ele := tr.Delete(keys[i])
			require.NotNil(t, ele)
			require.Equal(t, ele.Key(), keys[i])
			require.Equal(t, ele.Value(), int64(keys[i]*2+1))
			require.Nil(t, tr.Delete(keys[i]))
			require.Nil(t, tr.Search(keys[i]))

			checkCorrect(t, tr.root)
			require.Equal(t, tr.Len(), length-i-1)
		}

		require.Nil(t, tr.root)
		require.Equal(t, tr.Len(), 0)
	}
}
