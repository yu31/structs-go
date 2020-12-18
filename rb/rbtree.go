// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package rb

import (
	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/internal/tree"
)

var (
	_ container.Container = (*Tree)(nil)
	_ container.Element   = (*treeNode)(nil)
	_ container.Tree      = (*Tree)(nil)
	_ container.TreeNode  = (*treeNode)(nil)
)

const (
	red int8 = iota
	black
)

// treeNode is used for Red-Black Tree.
type treeNode struct {
	key    container.Key
	value  container.Value
	left   *treeNode
	right  *treeNode
	parent *treeNode
	color  int8
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

// Tree implements the Red-Black Tree.
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

// Root returns the root node of the tree.
func (tr *Tree) Root() container.TreeNode {
	if tr.root == nil {
		return nil
	}
	return tr.root
}

// Len return number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts a new element if the key doesn't exist, or returns the existing element for the key if present.
// The bool result is true if an element was inserted, false if searched.
func (tr *Tree) Insert(k container.Key, v container.Value) (container.Element, bool) {
	node, ok := tr.insertOrSearch(k, v)
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
	node := tr.searchNode(k)
	if node == nil {
		return nil
	}
	tr.replaceNode(node, tr.createNode(k, v, nil))
	return node
}

// Upsert inserts or updates an element by giving key and value.
// The bool result is true if an element was inserted, false if an element was updated.
func (tr *Tree) Upsert(k container.Key, v container.Value) (container.Element, bool) {
	node, ok := tr.insertOrSearch(k, v)
	if !ok {
		tr.replaceNode(node, tr.createNode(k, v, nil))
	}
	return node, ok
}

// Search searches the element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k container.Key) container.Element {
	node := tr.searchNode(k)
	if node == nil {
		return nil
	}
	return node
}

// Iter return an Iterator, it's a wrap for bs.Iterator.
func (tr *Tree) Iter(start container.Key, boundary container.Key) container.Iterator {
	return tree.NewIterator(tr.root, start, boundary)
}

// Range calls f sequentially each TreeNode present in the Tree.
// If f returns false, range stops the iteration.
func (tr *Tree) Range(start container.Key, boundary container.Key, f func(ele container.Element) bool) {
	tree.Range(tr.root, start, boundary, func(node container.TreeNode) bool {
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

// The insertOrSearch inserts and returns a new node with the given key and value if key doesn't exists.
// Or else, returns the exists node for the key if present.
// The ok result is true if the node was inserted, false if searched.
func (tr *Tree) insertOrSearch(k container.Key, v container.Value) (node *treeNode, ok bool) {
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

	tr.insertReBalance(node)
	tr.len++
	ok = true
	return
}

// Searches and deletes a node of a given key.
func (tr *Tree) deleteAndSearch(k container.Key) *treeNode {
	node := tr.searchNode(k)
	tr.deleteNode(node)
	return node
}

// Creates a new node with the giving key and value.
func (tr *Tree) createNode(k container.Key, v container.Value, p *treeNode) *treeNode {
	return &treeNode{
		key:    k,
		value:  v,
		left:   nil,
		right:  nil,
		parent: p,
		color:  red,
	}
}

// Deletes a node.
func (tr *Tree) deleteNode(d *treeNode) {
	if d == nil {
		return
	}
	if d.left != nil && d.right != nil {
		// Replace the location of the deleted node with its successor
		x := d.left
		for x.right != nil {
			x = x.right
		}
		// Removes the node x.
		tr.deleteNode(x)
		// Replaced deleted node with x.
		tr.replaceNode(d, x)
		return
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
		tr.deleteReBalance(c, d.parent)
	}

	//reset the unused field.
	d.left = nil
	d.right = nil
	d.parent = nil
	d.color = -1

	tr.len--
}

// Replace old with n0.
func (tr *Tree) replaceNode(old, n0 *treeNode) {
	n0.left = old.left
	n0.right = old.right
	n0.color = old.color
	n0.parent = old.parent

	if old.left != nil {
		old.left.parent = n0
	}
	if old.right != nil {
		old.right.parent = n0
	}

	if old.parent == nil {
		tr.root = n0
	} else if old.parent.left == old {
		old.parent.left = n0
	} else {
		old.parent.right = n0
	}

	// reset the unused field.
	old.left = nil
	old.right = nil
	old.parent = nil
	old.color = -1
}

// Search the node of a given key.
func (tr *Tree) searchNode(k container.Key) (node *treeNode) {
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

// Re-Balance after inserts a new node.
// n is the newly inserted node.
func (tr *Tree) insertReBalance(n *treeNode) {
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
		tr.insertReBalance(g)
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

// Re-Balance after delete a node.
// n is the replaces node of deleted node, and p is the parent node of deleted node.
func (tr *Tree) deleteReBalance(n *treeNode, p *treeNode) {
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
			tr.deleteReBalance(p, p.parent)
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
			tr.deleteReBalance(p, p.parent)
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
