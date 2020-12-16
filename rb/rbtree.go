// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package rb

import (
	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
)

const (
	red int8 = iota
	black
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element
type TreeNode = container.TreeNode

// treeNode is used for Red-Black Tree.
//
// And it is also the implementation of interface container.Element and container.TreeNode
type treeNode struct {
	key    Key
	value  Value
	left   *treeNode
	right  *treeNode
	parent *treeNode
	color  int8
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

// Tree implements the Red-Black Tree.
//
// And it is also the implementation of interface container.Container
type Tree struct {
	root *treeNode
	len  int
}

// New creates an Red-Black Tree.
func New() *Tree {
	return &Tree{
		root: nil,
		len:  0,
	}
}

// Len return number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts and returns an Element with the given key and value.
// Returns nil if key already exists.
func (tr *Tree) Insert(k Key, v Value) Element {
	node, ok := tr.insertOrSearch(k, v)
	if !ok {
		return nil
	}
	return node
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Element {
	d := tr.search(k)
	if d == nil {
		return nil
	}

	if d.left != nil && d.right != nil {
		x := d.left
		for x.right != nil {
			x = x.right
		}
		tr.swap(d, x)
		d = x
	}

	var c *treeNode

	if d.left != nil {
		c = d.left
	} else {
		c = d.right
	}
	if c != nil {
		c.parent = d.parent
	}

	if d.parent == nil {
		tr.root = c
	} else if d.parent.left == d {
		d.parent.left = c
	} else {
		d.parent.right = c
	}

	if d.color == black {
		tr.deleteBalance(c, d.parent)
	}

	// reset the unused field.
	d.left = nil
	d.right = nil
	d.parent = nil
	d.color = -1

	tr.len--
	return d
}

// Update updates an Element with the given key and value, And returns the old element.
// Returns nil if the key not be found.
func (tr *Tree) Update(k Key, v Value) Element {
	node := tr.search(k)
	if node != nil {
		tr.update(node, k, v)
	}
	return node
}

// Replace inserts or updates an Element by giving key and value.
//
// The action are same as the Insert method if key not found,
// And are same as the Update method if key exists.
func (tr *Tree) Replace(k Key, v Value) Element {
	n, ok := tr.insertOrSearch(k, v)
	if !ok {
		tr.update(n, k, v)
	}
	return n
}

// Search searches the Element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k Key) Element {
	return tr.search(k)
}

// Iter return an Iterator, it's a wrap for bs.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	return bs.NewIterator(tr.root, start, boundary)
}

// The insertOrSearch inserts and returns a newly node with the given key and value.
// Or else, returns the exists node for the key if present.
// The ok result is true if the node was inserted, false if searched.
func (tr *Tree) insertOrSearch(k Key, v Value) (node *treeNode, ok bool) {
	node = tr.root
	for node != nil {
		cmp := k.Compare(node.key)
		if cmp == 0 {
			// The key already exists, returns it.
			return
		}

		if cmp == -1 {
			if node.left == nil {
				node.left = tr.createNode(k, v, node)
				node = node.left
				break
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = tr.createNode(k, v, node)
				node = node.right
				break
			}
			node = node.right
		}
	}

	if node == nil {
		node = tr.createNode(k, v, nil)
	}

	tr.insertBalance(node)
	tr.len++
	ok = true
	return
}

// Help ot creates a newly node and instead of the node n.
func (tr *Tree) update(node *treeNode, k Key, v Value) {
	p := node.parent

	n0 := tr.createNode(k, v, p)
	n0.left = node.left
	n0.right = node.right
	n0.color = node.color

	if node.left != nil {
		node.left.parent = n0
	}
	if node.right != nil {
		node.right.parent = n0
	}

	if p == nil {
		tr.root = n0
	} else if p.left == node {
		p.left = n0
	} else {
		p.right = n0
	}

	// reset the unused field.
	node.left = nil
	node.right = nil
	node.parent = nil
	node.color = -1
}

// Search the node of a given key.
func (tr *Tree) search(k Key) (node *treeNode) {
	node = tr.root
	for node != nil {
		cmp := k.Compare(node.key)
		if cmp == -1 {
			node = node.left
		} else if cmp == 1 {
			node = node.right
		} else {
			return
		}
	}
	return
}

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

func (tr *Tree) insertBalance(node *treeNode) {
	n := node
	if n.parent == nil {
		n.color = black
		tr.root = n
		return
	}
	if n.parent.color == black {
		return
	}

	var (
		p, g, u *treeNode
	)

	p = n.parent
	g = n.parent.parent

	if g.left == p {
		u = g.right
	} else {
		u = g.left
	}

	if u != nil && u.color == red {
		g.color = red
		p.color = black
		u.color = black
		tr.insertBalance(g)
		return
	}

	if g.left == p {
		if p.right == n {
			tr.leftRotate(p)
			p = g.left
		}
		g.color = red
		p.color = black
		tr.rightRotate(g)
	} else {
		if p.left == n {
			tr.rightRotate(p)
			p = g.right
		}
		g.color = red
		p.color = black
		tr.leftRotate(g)
	}
}

func (tr *Tree) deleteBalance(node *treeNode, parent *treeNode) {
	n := node
	p := parent
	if n != nil && n.color == red {
		n.color = black
		return
	}
	if p == nil {
		tr.root = n
		return
	}

	var s *treeNode

	if p.left == n {
		s = p.right
		if s.color == red {
			s.color = black
			p.color = red
			tr.leftRotate(p)
			s = p.right
		}
		if (s.left == nil || s.left.color == black) && (s.right == nil || s.right.color == black) {
			s.color = red
			tr.deleteBalance(p, p.parent)
			return
		}
		if (s.left != nil && s.left.color == red) && (s.right == nil || s.right.color == black) {
			s.color = red
			s.left.color = black
			tr.rightRotate(s)
			s = p.right
		}
		if s.right != nil && s.right.color == red {
			s.color = p.color
			p.color = black
			s.right.color = black
			tr.leftRotate(p)
		}
	} else {
		s = p.left
		if s.color == red {
			s.color = black
			p.color = red
			tr.rightRotate(p)
			s = p.left
		}
		if (s.left == nil || s.left.color == black) && (s.right == nil || s.right.color == black) {
			s.color = red
			tr.deleteBalance(p, p.parent)
			return
		}
		if (s.right != nil && s.right.color == red) && (s.left == nil || s.left.color == black) {
			s.color = red
			s.right.color = black
			tr.leftRotate(s)
			s = p.left
		}
		if s.left != nil && s.left.color == red {
			s.color = p.color
			p.color = black
			s.left.color = black
			tr.rightRotate(p)
		}
	}
}

func (tr *Tree) createNode(k Key, v Value, p *treeNode) *treeNode {
	return &treeNode{
		key:    k,
		value:  v,
		left:   nil,
		right:  nil,
		parent: p,
		color:  red,
	}
}

func (tr *Tree) leftRotate(node *treeNode) {
	r := node.right
	if r.left != nil {
		r.left.parent = node
	}

	node.right = r.left
	r.left = node

	r.parent = node.parent
	node.parent = r

	if r.parent == nil {
		tr.root = r
	} else if r.parent.left == node {
		r.parent.left = r
	} else {
		r.parent.right = r
	}
}

func (tr *Tree) rightRotate(node *treeNode) {
	l := node.left
	if l.right != nil {
		l.right.parent = node
	}

	node.left = l.right
	l.right = node

	l.parent = node.parent
	node.parent = l

	if l.parent == nil {
		tr.root = l
	} else if l.parent.left == node {
		l.parent.left = l
	} else {
		l.parent.right = l
	}
}
