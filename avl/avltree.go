// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package avl

import (
	"fmt"

	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element
type TreeNode = bs.TreeNode

// treeNode is used for avl tree.
//
// And it is also the implementation of interface container.Element and bs.TreeNode
type treeNode struct {
	key    Key
	value  Value
	left   *treeNode
	right  *treeNode
	height int
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

// Tree implements the AVL Tree.
//
// And it is also the implementation of interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates an AVL Tree.
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
	var d *treeNode
	tr.root, d = tr.delete(tr.root, k)
	if d == nil {
		return nil
	}

	// reset the unused field.
	d.left = nil
	d.right = nil
	d.height = -1

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
	return bs.NewIterator(tr.root, start, boundary)
}

// Try to creates and inserts a node with the key and value.
//
// If the key not exists, it will creates and returns a newly node n, and ok is true.
// If the key already exists, n is the node where key is, and ok is false.
func (tr *Tree) insert(k Key, v Value) (p *treeNode, n *treeNode, ok bool) {
	tr.root, p, n, ok = tr.insertBalance(tr.root, k, v)
	if !ok {
		return
	}
	tr.len++
	return
}

// Help ot creates a newly node and instead of the node n.
func (tr *Tree) update(n *treeNode, p *treeNode, k Key, v Value) {
	n0 := tr.createNode(k, v)
	n0.left = n.left
	n0.right = n.right
	n0.height = n.height

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
	n.height = -1
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

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

// return (root root, new node)
// ok is false means no newly node created.
func (tr *Tree) insertBalance(r0 *treeNode, k Key, v Value) (root *treeNode, p *treeNode, n *treeNode, ok bool) {
	root = r0
	if r0 == nil {
		n = tr.createNode(k, v)
		root = n
		p = nil
		ok = true
		return
	}

	cmp := k.Compare(r0.key)
	if cmp == 0 {
		n = r0
		return
	}

	// search the parent node
	p = r0
	if p.left != nil && p.left.key.Compare(k) == 0 {
		// Found the key
		n = p.left
		return
	}
	if p.right != nil && p.right.key.Compare(k) == 0 {
		// Found the key
		n = p.right
		return
	}

	if cmp == -1 {
		// Insert into the left subtree.
		root.left, p, n, ok = tr.insertBalance(root.left, k, v)
	} else {
		// Insert into the right subtree
		root.right, p, n, ok = tr.insertBalance(root.right, k, v)
	}

	if ok {
		root = tr.reBalance(root)
	}

	return
}

// return (root root, delete node).
func (tr *Tree) delete(root *treeNode, k Key) (*treeNode, *treeNode) {
	var d *treeNode
	if root == nil {
		// not found
		return nil, nil
	} else {
		cmp := k.Compare(root.key)
		if cmp == -1 {
			// delete from the left subtree.
			root.left, d = tr.delete(root.left, k)
		} else if cmp == 1 {
			// delete from the right subtree.
			root.right, d = tr.delete(root.right, k)
		} else {
			if root.left != nil && root.right != nil {
				if tr.nodeHeight(root.left) > tr.nodeHeight(root.right) {
					x := root.left
					for x.right != nil {
						x = x.right
					}
					tr.swap(root, x)
					root.left, d = tr.delete(root.left, k)
				} else {
					x := root.right
					for x.left != nil {
						x = x.left
					}
					tr.swap(root, x)
					root.right, d = tr.delete(root.right, k)
				}
			} else {
				d = root
				if d.left != nil {
					root = d.left
				} else {
					root = d.right
				}
			}
		}
		if root != nil {
			root = tr.reBalance(root)
		}
	}

	return root, d
}

func (tr *Tree) reBalance(n *treeNode) *treeNode {
	if n == nil {
		return nil
	}

	factor := tr.nodeHeight(n.left) - tr.nodeHeight(n.right)

	switch factor {
	case -1, 0, 1:
		n.height = tr.calculateHeight(n)
	case 2:
		// Left subtree higher than right subtree.
		if tr.nodeHeight(n.left.right) > tr.nodeHeight(n.left.left) {
			n.left = tr.leftRotate(n.left)
		}
		n = tr.rightRotate(n)
	case -2:
		// Left subtree lower than right subtree.
		if tr.nodeHeight(n.right.left) > tr.nodeHeight(n.right.right) {
			n.right = tr.rightRotate(n.right)
		}
		n = tr.leftRotate(n)
	default:
		panic(fmt.Errorf("avl: unexpected cases with invalid factor <%d>", factor))
	}

	return n
}

func (tr *Tree) createNode(k Key, v Value) *treeNode {
	return &treeNode{
		key:    k,
		value:  v,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (tr *Tree) nodeHeight(n *treeNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (tr *Tree) calculateHeight(n *treeNode) int {
	lh := tr.nodeHeight(n.left)
	rh := tr.nodeHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func (tr *Tree) leftRotate(n *treeNode) *treeNode {
	r := n.right

	n.right = r.left
	r.left = n

	n.height = tr.calculateHeight(n)
	r.height = tr.calculateHeight(r)

	return r
}

func (tr *Tree) rightRotate(n *treeNode) *treeNode {
	l := n.left

	n.left = l.right
	l.right = n

	n.height = tr.calculateHeight(n)
	l.height = tr.calculateHeight(l)

	return l
}
