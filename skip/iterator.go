// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package skip

import (
	"github.com/yu31/structs-go/container"
)

var _ container.Iterator = (*Iterator)(nil)

// Iterator creates an Iterator positioned on the first element that key >= start key.
// If the start key is nil, it will return from the beginning.
// It yields only keys that < boundary. If boundary is nil, iteration until the end.
//
// Thus, the ranges is: start <= x < boundary.
type Iterator struct {
	node *listNode
	end  *listNode
}

// creates an Iterator.
func newIterator(sl *List, start container.Key, boundary container.Key) *Iterator {
	var node, end *listNode

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

	iter := &Iterator{
		node: node,
		end:  end,
	}
	return iter
}

// Valid represents whether to have more elements in the Iterator.
func (iter *Iterator) Valid() bool {
	if iter.node == nil || iter.node == iter.end {
		return false
	}
	return true
}

// Next returns a element and moved the iterator to the next element.
// Returns nil if no more elements.
func (iter *Iterator) Next() container.Element {
	if !iter.Valid() {
		return nil
	}
	n := iter.node
	iter.node = iter.node.next[0]
	return n
}
