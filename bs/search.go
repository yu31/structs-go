// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"reflect"

	"github.com/yu31/gostructs/stack"
)

// SearchRange to get the data for the specified range.
//
// Range interval: ( start <= x < boundary ).
// We will return data from the beginning if start is nil,
// And return data util the end if boundary is nil.
func SearchRange(root TreeNode, start Key, boundary Key, f func(n TreeNode)) {
	if root == nil {
		return
	}

	s := stack.Default()
	p := root
	for !s.Empty() || (p != nil && !reflect.ValueOf(p).IsNil()) {
		if !reflect.ValueOf(p).IsNil() {
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
			n := s.Pop().(TreeNode)
			p = n.Right()

			f(n)
		}
	}
}

// SearchLastLT search for the last node that less than the key.
func SearchLastLT(root TreeNode, key Key) TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n TreeNode

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

// SearchLastLE search for the last node that less than or equal to the key.
func SearchLastLE(root TreeNode, key Key) TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n TreeNode

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

// SearchFirstGT search for the first node that greater than to the key.
func SearchFirstGT(root TreeNode, key Key) TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n TreeNode

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

// SearchFirstGE search for the first node that greater than or equal to the key.
func SearchFirstGE(root TreeNode, key Key) TreeNode {
	if root == nil || key == nil {
		return nil
	}

	var n TreeNode

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
