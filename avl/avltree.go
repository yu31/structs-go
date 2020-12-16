// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package avl

import (
	"fmt"

	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/internal/tree"
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element
type TreeNode = container.TreeNode

// treeNode is used for avl tree.
//
// And it is also the implementation of interface container.Element and container.TreeNode
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
	return &Tree{
		root: nil,
		len:  0,
	}
}

// Root returns the root node of the tree.
func (tr *Tree) Root() TreeNode {
	return tr.root
}

// Len returns the number of elements.
func (tr *Tree) Len() int {
	return tr.len
}

// Insert inserts and returns an Element with given key and value if key doesn't exists.
// Or else, returns the existing Element for the key if present.
// The bool result is true if an Element was inserted, false if searched.
func (tr *Tree) Insert(k Key, v Value) (Element, bool) {
	_, node, ok := tr.insertOrSearch(k, v)
	return node, ok
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (tr *Tree) Delete(k Key) Element {
	var d *treeNode
	tr.root, d = tr.deleteBalance(tr.root, k)
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
	node, parent := tr.searchNode(k)
	if node != nil {
		tr.updateNode(node, parent, k, v)
	}
	return node
}

// Replace inserts or updates an Element by giving key and value.
// The bool result is true if an Element was inserted, false if an Element was updated.
//
// The operation are same as the Insert method if key not found,
// And are same as the Update method if key exists.
func (tr *Tree) Replace(k Key, v Value) (Element, bool) {
	parent, node, ok := tr.insertOrSearch(k, v)
	if !ok {
		tr.updateNode(node, parent, k, v)
	}
	return node, ok
}

// Search searches the Element of a given key.
// Returns nil if key not found.
func (tr *Tree) Search(k Key) Element {
	node, _ := tr.searchNode(k)
	return node
}

// Iter return an Iterator, it's a wrap for bs.Iterator.
func (tr *Tree) Iter(start Key, boundary Key) container.Iterator {
	return tree.NewIterator(tr.root, start, boundary)
}

// Range calls f sequentially each TreeNode present in the Tree.
// If f returns false, range stops the iteration.
func (tr *Tree) Range(start Key, boundary Key, f func(ele Element) bool) {
	tree.Range(tr.root, start, boundary, func(node TreeNode) bool {
		return f(node)
	})
}

// LastLT searches for the last node that less than the key.
func (tr *Tree) LastLT(k Key) Element {
	return tree.LastLT(tr.root, k)
}

// LastLE search for the last node that less than or equal to the key.
func (tr *Tree) LastLE(k Key) Element {
	return tree.LastLE(tr.root, k)
}

// FirstGT search for the first node that greater than to the key.
func (tr *Tree) FirstGT(k Key) Element {
	return tree.FirstGT(tr.root, k)
}

// FirstGE search for the first node that greater than or equal to the key.
func (tr *Tree) FirstGE(k Key) Element {
	return tree.FirstGE(tr.root, k)
}

// The insertOrSearch inserts and returns a new node with given key and value if key not exists.
// Or else, returns the exists node and its parent node for the key if present.
// The ok result is true if the node was inserted, false if searched.
func (tr *Tree) insertOrSearch(k Key, v Value) (parent *treeNode, node *treeNode, ok bool) {
	tr.root, node, parent, ok = tr.insertBalance(tr.root, k, v)
	if !ok {
		return
	}
	tr.len++
	return
}

// Helps to creates an tree node with given key and value.
func (tr *Tree) createNode(k Key, v Value) *treeNode {
	return &treeNode{
		key:    k,
		value:  v,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

// Help to creates a new tree node and instead of the node.
func (tr *Tree) updateNode(node *treeNode, parent *treeNode, k Key, v Value) {
	n0 := tr.createNode(k, v)
	n0.left = node.left
	n0.right = node.right
	n0.height = node.height

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
	node.height = -1
}

// Searches the node and its parent node of a given key.
func (tr *Tree) searchNode(k Key) (node *treeNode, parent *treeNode) {
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

func (tr *Tree) swap(n1, n2 *treeNode) {
	n1.key, n2.key = n2.key, n1.key
	n1.value, n2.value = n2.value, n1.value
}

// ok is false means no new node created.
func (tr *Tree) insertBalance(r0 *treeNode, k Key, v Value) (root *treeNode, node *treeNode, parent *treeNode, ok bool) {
	if r0 == nil {
		node = tr.createNode(k, v)
		root = node
		parent = nil
		ok = true
		return
	}

	root = r0

	cmp := k.Compare(root.key)
	if cmp == 0 {
		node = root
		return
	}

	// search the parent node
	parent = root
	if parent.left != nil && k.Compare(parent.left.key) == 0 {
		// Found the key
		node = parent.left
		return
	}
	if parent.right != nil && k.Compare(parent.right.key) == 0 {
		// Found the key
		node = parent.right
		return
	}

	if cmp == -1 {
		// Insert into the left subtree.
		root.left, node, parent, ok = tr.insertBalance(root.left, k, v)
	} else {
		// Insert into the right subtree
		root.right, node, parent, ok = tr.insertBalance(root.right, k, v)
	}

	if ok {
		root = tr.reBalance(root)
	}
	return
}

// return (root root, delete node).
func (tr *Tree) deleteBalance(r0 *treeNode, k Key) (root *treeNode, d *treeNode) {
	root = r0
	if root == nil {
		// The key not exists.
		return
	}

	cmp := k.Compare(root.key)
	if cmp == -1 {
		// delete from the left subtree.
		root.left, d = tr.deleteBalance(root.left, k)
	} else if cmp == 1 {
		// delete from the right subtree.
		root.right, d = tr.deleteBalance(root.right, k)
	} else {
		if root.left != nil && root.right != nil {
			if tr.nodeHeight(root.left) > tr.nodeHeight(root.right) {
				x := root.left
				for x.right != nil {
					x = x.right
				}
				tr.swap(root, x)
				root.left, d = tr.deleteBalance(root.left, k)
			} else {
				x := root.right
				for x.left != nil {
					x = x.left
				}
				tr.swap(root, x)
				root.right, d = tr.deleteBalance(root.right, k)
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

	root = tr.reBalance(root)
	return
}

func (tr *Tree) reBalance(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}

	factor := tr.nodeHeight(node.left) - tr.nodeHeight(node.right)
	switch factor {
	case -1, 0, 1:
		node.height = tr.calculateHeight(node)
	case 2:
		// Left subtree higher than right subtree.
		if tr.nodeHeight(node.left.right) > tr.nodeHeight(node.left.left) {
			node.left = tr.leftRotate(node.left)
		}
		node = tr.rightRotate(node)
	case -2:
		// Left subtree lower than right subtree.
		if tr.nodeHeight(node.right.left) > tr.nodeHeight(node.right.right) {
			node.right = tr.rightRotate(node.right)
		}
		node = tr.leftRotate(node)
	default:
		panic(fmt.Errorf("avl: unexpected cases with invalid factor <%d>", factor))
	}
	return node
}

func (tr *Tree) nodeHeight(node *treeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tr *Tree) calculateHeight(node *treeNode) int {
	lh := tr.nodeHeight(node.left)
	rh := tr.nodeHeight(node.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func (tr *Tree) leftRotate(node *treeNode) *treeNode {
	r := node.right

	node.right = r.left
	r.left = node

	node.height = tr.calculateHeight(node)
	r.height = tr.calculateHeight(r)
	return r
}

func (tr *Tree) rightRotate(node *treeNode) *treeNode {
	l := node.left

	node.left = l.right
	l.right = node

	node.height = tr.calculateHeight(node)
	l.height = tr.calculateHeight(l)
	return l
}
