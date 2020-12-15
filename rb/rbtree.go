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
type TreeNode = bs.TreeNode

// treeNode is used for Red-Black Tree.
//
// And it is also the implementation of interface container.Element and bs.TreeNode
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
	tr := &Tree{
		root: nil,
		len:  0,
	}
	return tr
}

// Len return number of elements.
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
				n = tr.createNode(k, v, p)
				p.left = n
				break
			}
			p = p.left
		} else if flag == 1 {
			if p.right == nil {
				n = tr.createNode(k, v, p)
				p.right = n
				break
			}
			p = p.right
		} else {
			// The key already exists. Not allowed duplicates.
			return nil
		}
	}
	if n == nil {
		n = tr.createNode(k, v, p)
	}

	tr.insertBalance(n)
	tr.len++

	return n
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Element {
	d := tr.root
	for d != nil {
		flag := k.Compare(d.key)
		if flag == -1 {
			d = d.left
		} else if flag == 1 {
			d = d.right
		} else {
			break
		}
	}
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

	d.left = nil
	d.right = nil
	d.parent = nil
	d.color = -1

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
	return bs.NewIterator(tr.root, start, boundary)
}

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

func (tr *Tree) insertBalance(n *treeNode) {
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

func (tr *Tree) deleteBalance(n *treeNode, p *treeNode) {
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

func (tr *Tree) leftRotate(n *treeNode) {
	r := n.right
	if r.left != nil {
		r.left.parent = n
	}

	n.right = r.left
	r.left = n

	r.parent = n.parent
	n.parent = r

	if r.parent == nil {
		tr.root = r
	} else if r.parent.left == n {
		r.parent.left = r
	} else {
		r.parent.right = r
	}
}

func (tr *Tree) rightRotate(n *treeNode) {
	l := n.left
	if l.right != nil {
		l.right.parent = n
	}

	n.left = l.right
	l.right = n

	l.parent = n.parent
	n.parent = l

	if l.parent == nil {
		tr.root = l
	} else if l.parent.left == n {
		l.parent.left = l
	} else {
		l.parent.right = l
	}
}
