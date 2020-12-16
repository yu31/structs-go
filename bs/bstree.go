// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"github.com/yu31/gostructs/container"
)

// treeNode is used for Binary Search Tree.
//
// And it is also the implementation of interface container.Element and bs.TreeNode
type treeNode struct {
	key   Key
	value Value
	left  *treeNode
	right *treeNode
}

// Key returns the key.
func (n *treeNode) Key() Key {
	return n.key
}

// Value returns the value.
func (n *treeNode) Value() Value {
	return n.value
}

// Left returns the left child of the TreeNode.
func (n *treeNode) Left() TreeNode {
	return n.left
}

// Right returns the right child of the TreeNode.
func (n *treeNode) Right() TreeNode {
	return n.right
}

// Tree implements the Binary Search Tree.
//
// And it is also the implementation of interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates an Binary Search Tree.
func New() *Tree {
	return &Tree{
		root: nil,
		len:  0,
	}
}

// Len returns the number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts and returns an Element with the given key and value.
// Returns nil if key already exists.
func (tr *Tree) Insert(k Key, v Value) Element {
	_, node, ok := tr.insertOrSearch(k, v)
	if !ok {
		return nil
	}
	return node
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Element {
	dd, d := tr.search(k)
	if d == nil {
		return nil
	}

	if d.left != nil && d.right != nil {
		xx := d
		x := d.left
		for x.right != nil {
			xx = x
			x = x.right
		}

		tr.swap(d, x)
		dd = xx
		d = x
	}

	var c *treeNode
	if d.left != nil {
		c = d.left
	} else {
		c = d.right
	}

	if dd == nil {
		tr.root = c
	} else if dd.left == d {
		dd.left = c
	} else {
		dd.right = c
	}

	// reset the unused field.
	d.left = nil
	d.right = nil

	tr.len--
	return d
}

// Update updates an Element with the given key and value, And returns the old element.
// Returns nil if the key not be found.
func (tr *Tree) Update(k Key, v Value) Element {
	parent, node := tr.search(k)
	if node != nil {
		tr.update(node, parent, k, v)
	}
	return node
}

// Replace inserts or updates an Element by giving key and value.
//
// The action are same as the Insert method if key not found,
// And are same as the Update method if key exists.
func (tr *Tree) Replace(k Key, v Value) Element {
	parent, node, ok := tr.insertOrSearch(k, v)
	if !ok {
		tr.update(node, parent, k, v)
	}
	return node
}

// Search searches the Element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k Key) Element {
	_, n := tr.search(k)
	return n
}

// Iter return an Iterator, it's a wrap for bs.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	return NewIterator(tr.root, start, boundary)
}

// The insertOrSearch inserts and returns a new node with the given key and value.
// Or else, returns the exists node and its parent node for the key if present.
// The ok result is true if the node was inserted, false if searched.
func (tr *Tree) insertOrSearch(k Key, v Value) (parent *treeNode, node *treeNode, ok bool) {
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

// Help ot creates a newly node and instead of the node.
func (tr *Tree) update(node *treeNode, parent *treeNode, k Key, v Value) {
	n0 := tr.createNode(k, v)
	n0.left = node.left
	n0.right = node.right

	if parent == nil {
		tr.root = n0
	} else if parent.left == node {
		parent.left = n0
	} else {
		parent.right = n0
	}

	// reset the unused field.
	node.left = nil
	node.right = nil
}

// Searches the node and its parent node of a given key.
func (tr *Tree) search(k Key) (parent *treeNode, node *treeNode) {
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

func (tr *Tree) createNode(k Key, v Value) *treeNode {
	return &treeNode{
		key:   k,
		value: v,
		left:  nil,
		right: nil,
	}
}

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}
