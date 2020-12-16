// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package skip

import (
	"math/rand"
	"time"

	"github.com/yu31/gostructs/container"
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element

const (
	maxLevel = 0x1f
)

// listNode is used for Skip List.
//
// And it is also the implementation of interface container.Element.
type listNode struct {
	key   Key
	value Value
	next  []*listNode
}

// Key returns the key.
func (n *listNode) Key() Key {
	return n.key
}

// Value returns the value.
func (n *listNode) Value() Value {
	return n.value
}

// List implements Skip List.
//
// And it is also the implementation of interface container.Container
type List struct {
	head  *listNode
	level int
	lens  []int
	r     *rand.Rand
}

// New creates an Skip List.
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

// Insert inserts and returns an Element with the given key and value.
// Returns nil if key already exists.
func (sl *List) Insert(k Key, v Value) Element {
	level := sl.chooseLevel()
	if level > sl.level {
		sl.level = level
	}

	updates := make([]*listNode, sl.level+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			// The key already exists. Not allowed duplicates.
			return nil
		}
		updates[i] = p
	}

	n := sl.createNode(k, v, level)
	for i := 0; i <= level; i++ {
		n.next[i] = updates[i].next[i]
		updates[i].next[i] = n
		sl.lens[i]++
	}
	return n
}

// Delete removes and returns the Element of a given key.
// Returns nil if not found.
func (sl *List) Delete(k Key) Element {
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

// Update updates an Element with the given key and value, And returns the old element.
// Returns nil if the key not be found.
func (sl *List) Update(k Key, v Value) Element {
	var node *listNode

	updates := make([]*listNode, sl.level+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			node = p.next[i]
		}
		updates[i] = p

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

	for i := 0; i < len(updates); i++ {
		if updates[i] != nil && updates[i].next[i] == node {
			updates[i].next[i] = n0
		}
	}

	// reset the unused field.
	node.next = nil
	return node
}

// Replace inserts or updates an Element by giving key and value.
//
// The action are same as the Insert method if key not found,
// And are same as the Update method if key exists.
func (sl *List) Replace(k Key, v Value) Element {
	var node *listNode

	updates := make([]*listNode, maxLevel+1)
	p := sl.head
	for i := sl.level; i >= 0; i-- {
		for p.next[i] != nil && p.next[i].key.Compare(k) == -1 {
			p = p.next[i]
		}
		if p.next[i] != nil && p.next[i].key.Compare(k) == 0 {
			node = p.next[i]
		}
		updates[i] = p
	}

	if node == nil {
		// The key not found, creates and inserts a newly node.
		level := sl.chooseLevel()
		if level > sl.level {
			for i := level; i > sl.level; i-- {
				updates[i] = sl.head
			}
			sl.level = level
		}

		node = sl.createNode(k, v, level)
		for i := 0; i <= level; i++ {
			node.next[i] = updates[i].next[i]
			updates[i].next[i] = node
			sl.lens[i]++
		}
		return node
	}

	// creates a new node and instead of the old node.
	n0 := &listNode{
		key:   k,
		value: v,
		next:  node.next,
	}

	for i := 0; i < len(updates); i++ {
		if updates[i] != nil && updates[i].next[i] == node {
			updates[i].next[i] = n0
		}
	}

	// reset the unused field.
	node.next = nil
	return node
}

// Search searches the Element of a given key.
// Returns nil if key not found.
func (sl *List) Search(k Key) Element {
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
func (sl *List) Iter(start Key, boundary Key) container.Iterator {
	return newIterator(sl, start, boundary)
}

func (sl *List) createNode(k Key, v Value, level int) *listNode {
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
func (sl *List) searchLastLT(k Key) *listNode {
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
func (sl *List) searchLastLE(k Key) *listNode {
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
func (sl *List) searchFirstGT(k Key) *listNode {
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
func (sl *List) searchFirstGE(k Key) *listNode {
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
