// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tree

import (
	"reflect"

	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/stack"
)

var _ container.Iterator = (*Iterator)(nil)

// Iterator to iteration return element.
//
// The range is start <= x < boundary.
// The elements will return from the beginning if start is nil,
// And return until the end if the boundary is nil.
//
// The Iterator return element with in-order traversal,
// And it can used with all-type binary search trees.
type Iterator struct {
	s        *stack.Stack
	start    container.Key
	boundary container.Key
}

// NewIterator creates an Iterator with given parameters.
func NewIterator(root container.TreeNode, start container.Key, boundary container.Key) *Iterator {
	s := stack.Default()

	fillStack(root, start, boundary, s)

	it := &Iterator{
		s:        s,
		start:    start,
		boundary: boundary,
	}
	return it
}

// Valid represents whether to have more elements in the Iterator.
func (it *Iterator) Valid() bool {
	return !it.s.Empty()
}

// Next returns a element and moved the iterator to the next Element.
// Returns nil if no more elements.
func (it *Iterator) Next() container.Element {
	if it.s.Empty() {
		return nil
	}

	p := it.s.Pop().(container.TreeNode)
	n := p

	fillStack(p.Right(), it.start, it.boundary, it.s)
	return n
}

func fillStack(root container.TreeNode, start container.Key, boundary container.Key, s *stack.Stack) {
	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		if start != nil && p.Key().Compare(start) == -1 {
			p = p.Right()
			continue
		}
		if boundary != nil && p.Key().Compare(boundary) != -1 {
			p = p.Left()
			continue
		}

		s.Push(p)
		p = p.Left()
	}
}
