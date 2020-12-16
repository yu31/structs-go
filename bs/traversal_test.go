package bs

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

func buildBSTree() (tr *Tree) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 25
	maxKey := length * 100

	tr = New()

	for i := 0; i < length; i++ {
		for {
			k := container.Int64(r.Intn(maxKey) + 1)
			if _, ok := tr.Insert(k, k*2+1); ok {
				break
			}
		}
	}
	return
}

func TestLDR(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	LDR(tr.root, func(n TreeNode) bool {
		r1 = append(r1, n)
		return true
	})

	var f func(node *treeNode)
	var r2 []Element
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		f(node.left)
		r2 = append(r2, node)
		f(node.right)
	}
	f(tr.root)

	require.Equal(t, r1, r2)
}

func TestDLR(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	DLR(tr.root, func(n TreeNode) bool {
		r1 = append(r1, n)
		return true
	})

	var f func(node *treeNode)
	var r2 []Element
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		r2 = append(r2, node)
		f(node.left)
		f(node.right)
	}
	f(tr.root)

	require.Equal(t, r1, r2)
}

func TestLRD(t *testing.T) {
	tr := buildBSTree()

	var r1 []Element
	LRD(tr.root, func(n TreeNode) bool {
		r1 = append(r1, n)
		return true
	})

	var f func(node *treeNode)
	var r2 []Element
	f = func(node *treeNode) {
		if node == nil {
			return
		}
		f(node.left)
		f(node.right)
		r2 = append(r2, node)
	}
	f(tr.root)

	require.Equal(t, r1, r2)
}
