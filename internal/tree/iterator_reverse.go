// Copyright (c) 2020, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tree

import (
	"reflect"

	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/stack"
)

var _ container.Iterator = (*IteratorReverse)(nil)

// Iter creates an reversed Iterator positioned on the first element that key >= start key.
// If the start key is nil, it will return from the beginning.
// It yields only keys that < boundary. If boundary is nil, iteration until the end.
//
// Thus, the ranges is: start <= x < boundary.
//
// The Iterator return element with in-order traversal,
// And it can used with all-type binary search trees.
type IteratorReverse struct {
	stack    *stack.Stack
	start    container.Key
	boundary container.Key
}

// NewIteratorReverse creates an reversed Iterator with given parameters.
func NewIteratorReverse(root container.TreeNode, start container.Key, boundary container.Key) *IteratorReverse {
	it := &IteratorReverse{
		stack:    stack.Default(),
		start:    start,
		boundary: boundary,
	}
	it.fillStack(root)
	return it
}

// Valid represents whether to have more elements in the Iterator.
func (it *IteratorReverse) Valid() bool {
	return !it.stack.Empty()
}

// Next returns a element and moved the iterator to the next Element.
// Returns nil if no more elements.
func (it *IteratorReverse) Next() container.Element {
	if it.stack.Empty() {
		return nil
	}

	p := it.stack.Pop().(container.TreeNode)
	n := p

	it.fillStack(p.Left())
	return n
}

func (it *IteratorReverse) fillStack(root container.TreeNode) {
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
		p = p.Right()
	}
}
