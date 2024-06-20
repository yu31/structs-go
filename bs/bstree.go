// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"github.com/yu31/structs-go/container"
	"github.com/yu31/structs-go/internal/tree"
)

var (
	_ container.Container = (*Tree)(nil)
	_ container.Element   = (*treeNode)(nil)
	_ container.Tree      = (*Tree)(nil)
	_ container.TreeNode  = (*treeNode)(nil)
)

// treeNode is used for Binary Search Tree.
type treeNode struct {
	key   container.Key
	value container.Value
	left  *treeNode
	right *treeNode
}

// Key returns the key.
func (n *treeNode) Key() container.Key {
	return n.key
}

// Value returns the value.
func (n *treeNode) Value() container.Value {
	return n.value
}

// Left returns the left child of the TreeNode.
func (n *treeNode) Left() container.TreeNode {
	if n.left == nil {
		return nil
	}
	return n.left
}

// Right returns the right child of the TreeNode.
func (n *treeNode) Right() container.TreeNode {
	if n.right == nil {
		return nil
	}
	return n.right
}

// Tree implements the Binary Search Tree.
type Tree struct {
	root *treeNode
	len  int
}

// New creates a Binary Search Tree.
func New() *Tree {
	return &Tree{
		root: nil,
		len:  0,
	}
}

// Root returns the root node of the tree.
func (tr *Tree) Root() container.TreeNode {
	if tr.root == nil {
		return nil
	}
	return tr.root
}

// Len returns the number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts a new element if the key doesn't exist, or returns the existing element for the key if present.
// The bool result is true if an element was inserted, false if searched.
func (tr *Tree) Insert(k container.Key, v container.Value) (container.Element, bool) {
	node, _, ok := tr.insertOrSearch(k, v)
	return node, ok
}

// Delete removes and returns the element of a given key.
// Returns nil if key not found.
func (tr *Tree) Delete(k container.Key) container.Element {
	node := tr.deleteAndSearch(k)
	if node == nil {
		return nil
	}
	return node
}

// Update updates an element with the given key and value, And returns the old element of key.
// Returns nil if the key not be found.
func (tr *Tree) Update(k container.Key, v container.Value) container.Element {
	node, parent := tr.searchNode(k)
	if node == nil {
		return nil
	}
	tr.replaceNode(node, parent, tr.createNode(k, v))
	return node
}

// Upsert inserts or updates an element by giving key and value.
// The bool result is true if an element was inserted, false if an element was updated.
func (tr *Tree) Upsert(k container.Key, v container.Value) (container.Element, bool) {
	node, parent, ok := tr.insertOrSearch(k, v)
	if !ok {
		tr.replaceNode(node, parent, tr.createNode(k, v))
	}
	return node, ok
}

// Search searches the element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k container.Key) container.Element {
	node, _ := tr.searchNode(k)
	if node == nil {
		return nil
	}
	return node
}

// Iter return an Iterator, it's a wrap for tree.Iterator.
func (tr *Tree) Iter(start container.Key, boundary container.Key) container.Iterator {
	return tree.NewIterator(tr.root, start, boundary)
}

// IterReverse return an Iterator, it's a wrap for tree.IterReverse.
func (tr *Tree) IterReverse(start container.Key, boundary container.Key) container.Iterator {
	return tree.NewIteratorReverse(tr.root, start, boundary)
}

// Range calls f sequentially each TreeNode present in the Tree.
// If f returns false, range stops the iteration.
func (tr *Tree) Range(start container.Key, boundary container.Key, f func(ele container.Element) bool) {
	tree.Range(tr.root, start, boundary, func(node container.TreeNode) bool {
		return f(node)
	})
}

// Reverse is similar to the Range method. But it iteration element in reverse.
// If f returns false, range stops the iteration.
func (tr *Tree) Reverse(start container.Key, boundary container.Key, f func(ele container.Element) bool) {
	tree.Reverse(tr.root, start, boundary, func(node container.TreeNode) bool {
		return f(node)
	})
}

// LastLT searches for the last node that less than the key.
func (tr *Tree) LastLT(k container.Key) container.Element {
	return tree.LastLT(tr.root, k)
}

// LastLE search for the last node that less than or equal to the key.
func (tr *Tree) LastLE(k container.Key) container.Element {
	return tree.LastLE(tr.root, k)
}

// FirstGT search for the first node that greater than to the key.
func (tr *Tree) FirstGT(k container.Key) container.Element {
	return tree.FirstGT(tr.root, k)
}

// FirstGE search for the first node that greater than or equal to the key.
func (tr *Tree) FirstGE(k container.Key) container.Element {
	return tree.FirstGE(tr.root, k)
}

// The insertOrSearch inserts and returns a new node with given key and value if key not exists.
// Or else, returns the exists node and its parent node for the key if present.
// The ok result is true if the node was inserted, false if searched.
func (tr *Tree) insertOrSearch(k container.Key, v container.Value) (node *treeNode, parent *treeNode, ok bool) {
	node = tr.root
	for node != nil {
		cmp := k.Compare(node.key)
		if cmp == 0 {
			// Found the exists key, returns it
			return
		}

		parent = node // The parent node of n.

		if cmp == -1 {
			if node.left == nil {
				node.left = tr.createNode(k, v)
				node = node.left
				break
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = tr.createNode(k, v)
				node = node.right
				break
			}
			node = node.right
		}
	}

	if node == nil {
		node = tr.createNode(k, v)
		tr.root = node
	}

	tr.len++
	ok = true
	return
}

// Searches and deletes a node of a given key.
func (tr *Tree) deleteAndSearch(k container.Key) *treeNode {
	node, parent := tr.searchNode(k)
	if node == nil {
		return nil
	}
	tr.deleteNode(node, parent)
	return node
}

// Creates a new node with the giving key and value.
func (tr *Tree) createNode(k container.Key, v container.Value) *treeNode {
	return &treeNode{
		key:   k,
		value: v,
		left:  nil,
		right: nil,
	}
}

// Deletes a node.
func (tr *Tree) deleteNode(d *treeNode, parent *treeNode) {
	if d.left != nil && d.right != nil {
		// Replace the location of the deleted node with its successor
		xx := d
		x := d.right
		for x.left != nil {
			xx = x
			x = x.left
		}
		// Removes the node x.
		tr.deleteNode(x, xx)
		// Replaced deleted node with x.
		tr.replaceNode(d, parent, x)
		return
	}

	var c *treeNode
	if d.left != nil {
		c = d.left
	} else {
		c = d.right
	}

	if parent == nil {
		tr.root = c
	} else if parent.left == d {
		parent.left = c
	} else {
		parent.right = c
	}
	tr.len--
	// reset the unused field.
	d.left = nil
	d.right = nil
}

// Replace old with n0. parent is old's parent node.
func (tr *Tree) replaceNode(old, parent, n0 *treeNode) {
	n0.left = old.left
	n0.right = old.right

	if parent == nil {
		tr.root = n0
	} else if parent.left == old {
		parent.left = n0
	} else {
		parent.right = n0
	}
	// reset the unused field.
	old.left = nil
	old.right = nil
}

// Searches the node and its parent node of a given key.
func (tr *Tree) searchNode(k container.Key) (node *treeNode, parent *treeNode) {
	node = tr.root
	for node != nil {
		cmp := k.Compare(node.key)
		if cmp == 0 {
			// Found the node of key.
			return
		}

		parent = node // The parent node of n.

		if cmp == -1 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return
}
