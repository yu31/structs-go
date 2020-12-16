package tests

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
)

func TestContainerTree_Root(t *testing.T) {
	process := func(tr container.Tree) {
		require.Nil(t, tr.Root())
		tr.Insert(container.Int64(3), "1024")
		require.NotNil(t, tr.Root())
		require.Equal(t, tr.Root().Key(), container.Int64(3))
		require.Equal(t, tr.Root().Value(), "1024")
	}

	// Test for all tree implementation.
	for name, f := range trees {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerTree_LDR(t *testing.T) {
	process := func(tr container.Tree) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(searchSeeds) {
			tr.Insert(k, int64(k*2+1))
		}

		var r1 []container.Element
		bs.LDR(tr.Root(), func(n container.TreeNode) bool {
			r1 = append(r1, n)
			return true
		})

		var f func(node container.TreeNode)
		var r2 []container.Element
		f = func(node container.TreeNode) {
			if node == nil || reflect.ValueOf(node).IsNil() {
				return
			}
			f(node.Left())
			r2 = append(r2, node)
			f(node.Right())
		}
		f(tr.Root())

		require.Equal(t, r1, r2)
	}

	// Test for all tree implementation.
	for name, f := range trees {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerTree_DLR(t *testing.T) {
	process := func(tr container.Tree) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(searchSeeds) {
			tr.Insert(k, int64(k*2+1))
		}

		var r1 []container.Element
		bs.DLR(tr.Root(), func(n container.TreeNode) bool {
			r1 = append(r1, n)
			return true
		})

		var f func(node container.TreeNode)
		var r2 []container.Element
		f = func(node container.TreeNode) {
			if node == nil || reflect.ValueOf(node).IsNil() {
				return
			}
			r2 = append(r2, node)
			f(node.Left())
			f(node.Right())
		}
		f(tr.Root())

		require.Equal(t, r1, r2)
	}

	// Test for all tree implementation.
	for name, f := range trees {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerTree_LRD(t *testing.T) {
	process := func(tr container.Tree) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(searchSeeds) {
			tr.Insert(k, int64(k*2+1))
		}

		var r1 []container.Element
		bs.LRD(tr.Root(), func(n container.TreeNode) bool {
			r1 = append(r1, n)
			return true
		})

		var f func(node container.TreeNode)
		var r2 []container.Element
		f = func(node container.TreeNode) {
			if node == nil || reflect.ValueOf(node).IsNil() {
				return
			}
			f(node.Left())
			f(node.Right())
			r2 = append(r2, node)
		}
		f(tr.Root())

		require.Equal(t, r1, r2)
	}

	// Test for all tree implementation.
	for name, f := range trees {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}
