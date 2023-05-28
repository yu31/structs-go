// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package minheap

import (
	"github.com/yu31/structs-go/container"
)

const (
	defaultCapacity = 64
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value

// Item is an item of a MinHeap.
type Item struct {
	key   Key
	value Value
	// The index of the item in the heap.
	index int
}

// Key returns the key in the item.
func (item *Item) Key() Key {
	return item.key
}

// Value returns the value in the item.
func (item *Item) Value() Value {
	return item.value
}

// Value returns the index of the item.
func (item *Item) Index() int {
	return item.index
}

// MinHeap implements min heap and can used as priority queue.
type MinHeap struct {
	items []*Item
	cap   int
	len   int
}

// Default creates an MinHeap with default parameters.
func Default() *MinHeap {
	return New(defaultCapacity)
}

// New creates an MinHeap with given initialization capacity.
func New(c int) *MinHeap {
	h := &MinHeap{
		items: make([]*Item, c),
		cap:   c,
		len:   0,
	}
	return h
}

// Len return the number of elements in the heap.
func (h *MinHeap) Len() int {
	return h.len
}

// Cap return the current capacity of the heap.
func (h *MinHeap) Cap() int {
	return h.cap
}

// Empty represents whether the heap is empty.
func (h *MinHeap) Empty() bool {
	return h.len == 0
}

// Push adds an element to the heap, Return the index number of the location.
func (h *MinHeap) Push(k Key, v Value) *Item {
	h.autoGrow()
	item := &Item{
		key:   k,
		value: v,
		index: h.len,
	}
	h.items[h.len] = item

	h.up(h.len)
	h.len++
	return item
}

// Remove removes and returns the item at index i from the heap.
// The complexity is O(log n) where n = h.Len().
// Return nil if the i >= h.Len().
func (h *MinHeap) Remove(i int) *Item {
	if i >= h.Len() {
		return nil
	}
	item := h.delete(i)
	if i != h.len {
		if !h.down(i, h.len) {
			h.up(i)
		}
	}
	return item
}

// Pop returns and removes an element that at the head.
// Return nil if the heap is empty.
func (h *MinHeap) Pop() *Item {
	if h.Empty() {
		return nil
	}
	item := h.delete(0)
	_ = h.down(0, h.len)
	return item
}

// Peek returns the element that at the head.
// Return nil if the heap is empty.
func (h *MinHeap) Peek() *Item {
	if h.Empty() {
		return nil
	}
	return h.items[0]
}

func (h *MinHeap) delete(i int) *Item {
	item := h.items[i]
	h.len--
	h.swap(i, h.len)

	item.index = -1
	h.items[h.len] = nil // Prevent memory leaks.
	return item
}

// up build heap with bottom-up
func (h *MinHeap) up(i int) {
	var p int
	for {
		p = (i - 1) >> 1 // parent
		if p < 0 || i == p || h.compare(i, p) != -1 {
			break
		}
		h.swap(p, i)
		i = p
	}
}

// down build heap with top-down.
// n is the length of items where valid element.
func (h *MinHeap) down(i0 int, n int) bool {
	i := i0
	for {
		c := (i << 1) + 1 // left child
		if c >= n || c < 0 {
			// after int overflow
			break
		}

		if r := c + 1; r < n && h.compare(r, c) == -1 {
			c = r // right child
		}

		if h.compare(c, i) != -1 {
			break
		}

		h.swap(i, c)
		i = c
	}

	return i > i0
}

func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].index = i
	h.items[j].index = j
}

func (h *MinHeap) compare(i, j int) int {
	return h.items[i].key.Compare(h.items[j].key)
}

func (h *MinHeap) autoGrow() {
	if h.len == h.cap {
		newCap := h.cap
		if h.len < 1024 {
			newCap += h.cap
		} else {
			newCap += h.cap / 2
		}
		h.grow(newCap)
	}
}

func (h *MinHeap) grow(c int) {
	if c > h.cap {
		items := h.items
		h.items = make([]*Item, c)
		h.cap = c
		copy(h.items, items)
	}
}
