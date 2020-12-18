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

// Range calls f sequentially each TreeNode present in the Tree.
// If f returns false, range stops the iteration.
//
// The range is start <= x < boundary.
// The elements will return from the beginning if start is nil,
// And return until the end if the boundary is nil.
func Range(root container.TreeNode, start container.Key, boundary container.Key, f func(node container.TreeNode) bool) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root
	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if p != nil && !reflect.ValueOf(p).IsNil() {
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
		} else {
			n := s.Pop().(container.TreeNode)
			p = n.Right()

			// Stop iteration if return false.
			if !f(n) {
				return
			}
		}
	}
}

// LastLT searches for the last node that less than the key.
func LastLT(root container.TreeNode, key container.Key) container.TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n container.TreeNode

	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		cmp := key.Compare(p.Key())
		if cmp == 1 {
			n = p
			p = p.Right()
		} else {
			p = p.Left()
		}
	}

	return n
}

// LastLE search for the last node that less than or equal to the key.
func LastLE(root container.TreeNode, key container.Key) container.TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n container.TreeNode

	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		cmp := key.Compare(p.Key())
		if cmp == 1 {
			n = p
			p = p.Right()
		} else if cmp == -1 {
			p = p.Left()
		} else {
			n = p
			break
		}
	}

	return n
}

// FirstGT search for the first node that greater than to the key.
func FirstGT(root container.TreeNode, key container.Key) container.TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n container.TreeNode

	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		cmp := key.Compare(p.Key())
		if cmp == -1 {
			n = p
			p = p.Left()
		} else {
			p = p.Right()
		}
	}

	return n
}

// FirstGE search for the first node that greater than or equal to the key.
func FirstGE(root container.TreeNode, key container.Key) container.TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n container.TreeNode

	p := root
	for p != nil && !reflect.ValueOf(p).IsNil() {
		cmp := key.Compare(p.Key())
		if cmp == -1 {
			n = p
			p = p.Left()
		} else if cmp == 1 {
			p = p.Right()
		} else {
			n = p
			break
		}
	}

	return n
}
