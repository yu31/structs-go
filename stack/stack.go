// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package stack

const (
	defaultCapacity = 64
)

// Stack implements a stack (LIFO).
type Stack struct {
	items []interface{}
	cap   int
	len   int
}

// Default creates an Stack with default parameters.
func Default() *Stack {
	return New(defaultCapacity)
}

// New creates an Stack with the given the initialization capacity.
func New(n int) *Stack {
	s := &Stack{
		items: make([]interface{}, n),
		cap:   n,
		len:   0,
	}
	return s
}

// Len return the number of elements in the stack.
func (s *Stack) Len() int {
	return s.len
}

// Cap return the current capacity of the stack.
func (s *Stack) Cap() int {
	return s.cap
}

// Empty represents whether the stack is empty.
func (s *Stack) Empty() bool {
	return s.len == 0
}

// Push adds an element to the end of stack.
func (s *Stack) Push(item interface{}) {
	s.autoGrow()
	s.items[s.len] = item
	s.len++
}

// Pop returns and removes an element that at the end.
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	s.len--
	item := s.items[s.len]
	s.items[s.len] = nil // Prevent memory leaks.
	return item
}

// Peek returns the element that at the end.
func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	item := s.items[s.len-1]
	return item
}

func (s *Stack) autoGrow() {
	if s.len == s.cap {
		newCap := s.cap
		if s.len < 1024 {
			newCap += s.cap
		} else {
			newCap += s.cap / 2
		}
		s.grow(newCap)
	}
}

func (s *Stack) grow(c int) {
	if c > s.cap {
		items := s.items
		s.cap = c
		s.items = make([]interface{}, s.cap)
		copy(s.items, items)
	}
}
