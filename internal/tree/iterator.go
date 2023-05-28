// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tree

import (
	"reflect"

	"github.com/yu31/structs-go/container"
	"github.com/yu31/structs-go/stack"
)

var _ container.Iterator = (*Iterator)(nil)

// Iter creates an Iterator positioned on the first element that key >= start key.
// If the start key is nil, it will return from the beginning.
// It yields only keys that < boundary. If boundary is nil, iteration until the end.
//
// Thus, the ranges is: start <= x < boundary.
//
// The Iterator return element with in-order traversal,
// And it can used with all-type binary search trees.
type Iterator struct {
	stack    *stack.Stack
	start    container.Key
	boundary container.Key
}

// NewIterator creates an Iterator with given parameters.
func NewIterator(root container.TreeNode, start container.Key, boundary container.Key) *Iterator {
	it := &Iterator{
		stack:    stack.Default(),
		start:    start,
		boundary: boundary,
	}
	it.fillStack(root)
	return it
}

// Valid represents whether to have more elements in the Iterator.
func (it *Iterator) Valid() bool {
	return !it.stack.Empty()
}

// Next returns a element and moved the iterator to the next Element.
// Returns nil if no more elements.
func (it *Iterator) Next() container.Element {
	if it.stack.Empty() {
		return nil
	}

	p := it.stack.Pop().(container.TreeNode)
	n := p

	it.fillStack(p.Right())
	return n
}

func (it *Iterator) fillStack(root container.TreeNode) {
	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		if it.start != nil && p.Key().Compare(it.start) == -1 {
			p = p.Right()
			continue
		}
		if it.boundary != nil && p.Key().Compare(it.boundary) != -1 {
			p = p.Left()
			continue
		}

		it.stack.Push(p)
		p = p.Left()
	}
}
