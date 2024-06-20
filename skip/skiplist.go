// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package skip

import (
	"math/rand"
	"time"

	"github.com/yu31/structs-go/container"
)

var (
	_ container.Container = (*List)(nil)
	_ container.Element   = (*listNode)(nil)
)

const (
	maxLevel = 0x1f
)

// listNode is used for Skip List.
type listNode struct {
	key   container.Key
	value container.Value
	next  []*listNode
}

// Key returns the key.
func (n *listNode) Key() container.Key {
	return n.key
}

// Value returns the value.
func (n *listNode) Value() container.Value {
	return n.value
}

// List implements Skip List.
type List struct {
	head  *listNode
	level int
	lens  []int
	r     *rand.Rand
}

// New creates a Skip List.
func New() *List {
	sl := new(List)
	sl.head = sl.createNode(nil, nil, maxLevel)
	sl.level = 0
	sl.lens = make([]int, maxLevel+1)
	sl.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return sl
}

// Len returns the number of elements.
func (sl *List) Len() int {
	return sl.lens[0]
}

// Insert inserts a new element if the key doesn't exist, or returns the existing element for the key if present.
// The bool result is true if an element was inserted, false if searched.
func (sl *List) Insert(k container.Key, v container.Value) (container.Element, bool) {
	level := sl.chooseLevel()
	if level > sl.level {
		sl.level = level
	}

	previous := make([]*listNode, sl.level+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			// The key already exists. Not allowed duplicates.
			return p.next[i], false
		}
		previous[i] = p
	}

	n := sl.createNode(k, v, level)
	for i := 0; i <= level; i++ {
		n.next[i] = previous[i].next[i]
		previous[i].next[i] = n
		sl.lens[i]++
	}
	return n, true
}

// Delete removes and returns the element of a given key.
// Returns nil if key not found.
func (sl *List) Delete(k container.Key) container.Element {
	var d *listNode
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			if d == nil {
				d = p.next[i]
			}
			p.next[i] = p.next[i].next[i]
			sl.lens[i]--
		}

		if sl.head.next[i] == nil && i != 0 {
			sl.level--
		}
	}
	if d == nil {
		return nil
	}
	// reset the unused field.
	d.next = nil
	return d
}

// Update updates an element with the given key and value, And returns the old element of key.
// Returns nil if the key not be found.
func (sl *List) Update(k container.Key, v container.Value) container.Element {
	var node *listNode

	previous := make([]*listNode, sl.level+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			node = p.next[i]
		}
		previous[i] = p

	}
	if node == nil {
		return nil
	}

	// creates a new node and instead of the old node.
	n0 := &listNode{
		key:   k,
		value: v,
		next:  node.next,
	}

	for i := 0; i < len(previous); i++ {
		if previous[i] != nil && previous[i].next[i] == node {
			previous[i].next[i] = n0
		}
	}

	// reset the unused field.
	node.next = nil
	return node
}

// Upsert inserts or updates an element by giving key and value.
// The bool result is true if an element was inserted, false if an element was updated.
func (sl *List) Upsert(k container.Key, v container.Value) (container.Element, bool) {
	var node *listNode

	previous := make([]*listNode, maxLevel+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			node = p.next[i]
		}
		previous[i] = p
	}

	if node == nil {
		// The key not found, creates and inserts a new node.
		level := sl.chooseLevel()
		if level > sl.level {
			for i := level; i > sl.level; i-- {
				previous[i] = sl.head
			}
			sl.level = level
		}

		node = sl.createNode(k, v, level)
		for i := 0; i <= level; i++ {
			node.next[i] = previous[i].next[i]
			previous[i].next[i] = node
			sl.lens[i]++
		}
		return node, true
	}

	// creates a new node and instead of the old node.
	n0 := &listNode{
		key:   k,
		value: v,
		next:  node.next,
	}

	for i := 0; i < len(previous); i++ {
		if previous[i] != nil && previous[i].next[i] == node {
			previous[i].next[i] = n0
		}
	}

	// reset the unused field.
	node.next = nil
	return node, false
}

// Search searches the element of a given key.
// Returns nil if key not found.
func (sl *List) Search(k container.Key) container.Element {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			return p.next[i]
		}
	}
	return nil
}

// Iter return an Iterator, it's a wrap for skip.Iterator
func (sl *List) Iter(start container.Key, boundary container.Key) container.Iterator {
	return newIterator(sl, start, boundary)
}

// IterReverse return an Iterator, it's a wrap for tree.IterReverse.
func (sl *List) IterReverse(start container.Key, boundary container.Key) container.Iterator {
	// TODO:
	panic("skiplist: IterReverse method not implemented")
}

// Range calls f sequentially each TreeNode present in the Tree.
// If f returns false, range stops the iteration.
func (sl *List) Range(start container.Key, boundary container.Key, f func(elem container.Element) bool) {
	var node *listNode
	var end *listNode

	// If both the start and boundary are not nil, the start should less than the boundary.
	if !(start != nil && boundary != nil && start.Compare(boundary) != -1) {
		if start == nil {
			node = sl.head.next[0]
		} else {
			node = sl.searchFirstGE(start)
		}
		if boundary != nil {
			end = sl.searchFirstGE(boundary)
		}
	}

	for node != nil && node != end {
		// Stop iteration if return false.
		if !f(node) {
			return
		}
		node = node.next[0]
	}
}

// Reverse is similar to the Range method. But it iteration element in reverse.
// If f returns false, range stops the iteration.
func (sl *List) Reverse(start container.Key, boundary container.Key, f func(ele container.Element) bool) {
	// TODO:
	panic("skiplist: Reverse method not implemented")
}

// LastLT searches for the last node that less than the key.
func (sl *List) LastLT(k container.Key) container.Element {
	return sl.searchLastLT(k)
}

// LastLE search for the last node that less than or equal to the key.
func (sl *List) LastLE(k container.Key) container.Element {
	return sl.searchLastLE(k)
}

// FirstGT search for the first node that greater than to the key.
func (sl *List) FirstGT(k container.Key) container.Element {
	return sl.searchFirstGT(k)
}

// FirstGE search for the first node that greater than or equal to the key.
func (sl *List) FirstGE(k container.Key) container.Element {
	return sl.searchFirstGE(k)
}

// Creates a new node with the giving key and value.
func (sl *List) createNode(k container.Key, v container.Value, level int) *listNode {
	return &listNode{
		key:   k,
		value: v,
		next:  make([]*listNode, level+1),
	}
}

func (sl *List) chooseLevel() int {
	level := 0
	for sl.r.Int63()&1 == 1 && level < maxLevel {
		level++
	}
	return level
}

// Search the last node that less than the key.
func (sl *List) searchLastLT(k container.Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if i == 0 && p.key != nil {
			return p
		}
	}
	return nil
}

// Search the last node that less than or equal to the key.
func (sl *List) searchLastLE(k container.Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			return p.next[i]
		} else if i == 0 && p.key != nil {
			return p
		}

	}
	return nil
}

// Search the first node that greater than to the key.
func (sl *List) searchFirstGT(k container.Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].key.Compare(k) == 0 {
				return p.next[i].next[0]
			}
			if i == 0 {
				return p.next[i]
			}
		}

	}
	return nil
}

// Search the first node that greater than or equal to the key.
func (sl *List) searchFirstGE(k container.Key) *listNode {
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}

		if p.next[i] != nil {
			if p.next[i].key.Compare(k) == 0 || i == 0 {
				return p.next[i]
			}
		}

	}
	return nil
}
