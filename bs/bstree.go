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
	var n *treeNode
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			if p.left == nil {
				n = tr.createNode(k, v)
				p.left = n
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				n = tr.createNode(k, v)
				p.right = n
				break
			}
			p = p.right
		} else {
			// The key already exists. Not allowed duplicates.
			return nil
		}
	}

	if p == nil {
		n = tr.createNode(k, v)
		tr.root = n
	}

	tr.len++
	return n
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Element {
	var dd *treeNode
	d := tr.root

	for d != nil {
		flag := k.Compare(d.key)
		// Found the deletion key
		if flag == 0 {
			break
		}

		dd = d
		if flag == -1 {
			d = d.left
		} else {
			d = d.right
		}
	}

	// Not found.
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

	d.left = nil
	d.right = nil

	tr.len--
	return d
}

// Update updates and returns an Element with the given key and value.
// Returns nil if key not found.
func (tr *Tree) Update(k Key, v Value) Element {
	panic("not implemented")
}

// Replace inserts or updates an Element by giving key and value.
//
// The action are same as an Insert method if key not found,
// And are same as an Update method if found the key.
func (tr *Tree) Replace(k Key, v Value) Element {
	panic("not implemented")
}

// Search searches the Element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k Key) Element {
	p := tr.root
	for p != nil {
		flag := k.Compare(p.key)
		if flag == -1 {
			p = p.left
		} else if flag == 1 {
			p = p.right
		} else {
			return p
		}
	}
	return nil
}

// Iter return an Iterator, it's a wrap for bs.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	return NewIterator(tr.root, start, boundary)
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
