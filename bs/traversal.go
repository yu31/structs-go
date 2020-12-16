// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"reflect"

	"github.com/yu31/gostructs/stack"
)

// LDR calls f sequentially each TreeNode by in-order traversal.
// If f returns false, range stops the iteration.
// The order: Left -> Middle -> Right.
func LDR(root TreeNode, f func(node TreeNode) bool) {
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

			if !f(n) {
				return
			}
		}
	}
}

// DLR calls f sequentially each TreeNode by pre-order traversal.
// If f returns false, range stops the iteration.
// The order: Middle -> Left -> Right.
func DLR(root TreeNode, f func(node TreeNode) bool) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if !reflect.ValueOf(p).IsNil() {
			if !f(p) {
				return
			}

			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(TreeNode)
			p = p.Right()
		}
	}
}

// LRD calls f sequentially each TreeNode by post-order traversal.
// If f returns false, range stops the iteration.
// The order: Left -> Right -> Middle.
func LRD(root TreeNode, f func(node TreeNode) bool) {
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
			if !f(p) {
				return
			}

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
