// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"reflect"

	"github.com/yu31/gostructs/stack"
)

// LDR return node by in-order traversal.
// The order: Left -> Middle -> Right.
func LDR(root TreeNode, f func(n TreeNode)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if !reflect.ValueOf(p).IsNil() {
			s.Push(p)
			p = p.Left()
		} else {
			n := s.Pop().(TreeNode)
			p = n.Right()

			f(n)
		}
	}
}

// DLR return node by pre-order traversal.
// The order: Middle -> Left -> Right.
func DLR(root TreeNode, f func(n TreeNode)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if !reflect.ValueOf(p).IsNil() {
			f(p)

			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(TreeNode)
			p = p.Right()
		}
	}
}

// LRD return node by post-order traversal.
// The order: Left -> Right -> Middle.
func LRD(root TreeNode, f func(n TreeNode)) {
	if root == nil {
		return
	}

	var lastVisit TreeNode

	s := stack.Default()
	p := root

	for p != nil && !reflect.ValueOf(p).IsNil() {
		s.Push(p)
		p = p.Left()
	}

	for !s.Empty() {
		p = s.Pop().(TreeNode)
		if reflect.ValueOf(p.Right()).IsNil() || p.Right() == lastVisit {
			f(p)

			lastVisit = p
		} else {
			s.Push(p)
			p = p.Right()
			for !reflect.ValueOf(p).IsNil() {
				s.Push(p)
				p = p.Left()
			}
		}
	}
}
