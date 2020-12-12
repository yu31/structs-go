// Copyright (c) 2020, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"reflect"

	"github.com/yu31/gostructs/stack"
)

// Iterator to iteration return element.
//
// The element range is start <= x < boundary.
// The element will return from the beginning if start is nil,
// And return until the end if the boundary is nil.
//
// The Iterator return element with in-order traversal,
// And it can used with all-type binary search trees.
//
// And it is also the implementation of interface container.Iterator
type Iterator struct {
	s        *stack.Stack
	start    Key
	boundary Key
}

// NewIterator creates an Iterator with given parameters.
func NewIterator(root TreeNode, start Key, boundary Key) *Iterator {
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

// Next returns a Element and moved the iterator to the next Element.
// Returns nil if no more elements.
func (it *Iterator) Next() Element {
	if it.s.Empty() {
		return nil
	}

	p := it.s.Pop().(TreeNode)
	n := p

	fillStack(p.Right(), it.start, it.boundary, it.s)

	return n
}

func fillStack(root TreeNode, start Key, boundary Key, s *stack.Stack) {
	p := root
	for !reflect.ValueOf(p).IsNil() {
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
