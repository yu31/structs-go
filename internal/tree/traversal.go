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

// LDR calls f sequentially each TreeNode by in-order traversal.
// If f returns false, range stops the iteration.
// The order: Left -> Middle -> Right.
func LDR(root container.TreeNode, f func(node container.TreeNode) bool) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if p != nil && !reflect.ValueOf(p).IsNil() {
			s.Push(p)
			p = p.Left()
		} else {
			n := s.Pop().(container.TreeNode)
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
func DLR(root container.TreeNode, f func(node container.TreeNode) bool) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root

	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if p != nil && !reflect.ValueOf(p).IsNil() {
			if !f(p) {
				return
			}

			s.Push(p)
			p = p.Left()
		} else {
			p = s.Pop().(container.TreeNode)
			p = p.Right()
		}
	}
}

// LRD calls f sequentially each TreeNode by post-order traversal.
// If f returns false, range stops the iteration.
// The order: Left -> Right -> Middle.
func LRD(root container.TreeNode, f func(node container.TreeNode) bool) {
	if root == nil {
		return
	}

	var lastVisit container.TreeNode

	s := stack.Default()
	p := root

	for p != nil && !reflect.ValueOf(p).IsNil() {
		s.Push(p)
		p = p.Left()
	}

	for !s.Empty() {
		p = s.Pop().(container.TreeNode)
		if (p.Right() == nil || reflect.ValueOf(p.Right()).IsNil()) || p.Right() == lastVisit {
			if !f(p) {
				return
			}

			lastVisit = p
		} else {
			s.Push(p)
			p = p.Right()
			for p != nil && !reflect.ValueOf(p).IsNil() {
				s.Push(p)
				p = p.Left()
			}
		}
	}
}

// RDL calls f sequentially each TreeNode by reverse-order traversal.
// If f returns false, range stops the iteration.
// The order: Right -> Middle -> Left
func RDL(root container.TreeNode, f func(node container.TreeNode) bool) {
	if root == nil {
		return
	}
	s := stack.Default()
	p := root
	// Right -> Middle -> Left
	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if p != nil && !reflect.ValueOf(p).IsZero() {
			s.Push(p)
			p = p.Right()
		} else {
			n := s.Pop().(container.TreeNode)
			p = n.Left()

			if !f(n) {
				return
			}
		}
	}
}
