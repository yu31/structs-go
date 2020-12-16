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
	tr := &Tree{
		root: nil,
		len:  0,
	}
	return tr
}

// Len returns the number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts and returns an Element with the given key and value.
// Returns nil if key already exists.
func (tr *Tree) Insert(k Key, v Value) Element {
	_, n, ok := tr.insert(k, v)
	if !ok {
		return nil
	}
	return n
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
	p, n := tr.search(k)
	if n != nil {
		tr.update(n, p, k, v)
	}
	return n
}

// Replace inserts or updates an Element by giving key and value.
//
// The action are same as the Insert method if key not found,
// And are same as the Update method if key exists.
func (tr *Tree) Replace(k Key, v Value) Element {
	p, n, ok := tr.insert(k, v)
	if !ok {
		tr.update(n, p, k, v)
	}
	return n
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

// Try to creates and inserts a node with the key and value.
//
// If the key not exists, it will creates and returns a newly node n, and ok is true.
// If the key already exists, n is the node where key is, and ok is false.
func (tr *Tree) insert(k Key, v Value) (p *treeNode, n *treeNode, ok bool) {
	n = tr.root
	for n != nil {
		cmp := k.Compare(n.key)
		if cmp == 0 {
			// The key already exists, returns it.
			return
		}

		p = n // The parent node of n.

		if cmp == -1 {
			if n.left == nil {
				n.left = tr.createNode(k, v)
				n = n.left
				break
			}
			n = n.left
		} else {
			if n.right == nil {
				n.right = tr.createNode(k, v)
				n = n.right
				break
			}
			n = n.right
		}
	}

	if n == nil {
		n = tr.createNode(k, v)
		tr.root = n
	}

	ok = true
	tr.len++
	return
}

// Help ot creates a newly node and instead of the node n.
func (tr *Tree) update(n *treeNode, p *treeNode, k Key, v Value) {
	n0 := tr.createNode(k, v)
	n0.left = n.left
	n0.right = n.right

	if p == nil {
		tr.root = n0
	} else if p.left == n {
		p.left = n0
	} else {
		p.right = n0
	}

	// reset the unused field.
	n.left = nil
	n.right = nil
}

// Searches the node and its parent node of a given key.
func (tr *Tree) search(k Key) (p *treeNode, n *treeNode) {
	n = tr.root
	for n != nil {
		cmp := k.Compare(n.key)
		if cmp == 0 {
			// Found the node of key.
			return
		}

		p = n // The parent node of n.

		if cmp == -1 {
			n = n.left
		} else {
			n = n.right
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
