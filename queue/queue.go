// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package queue

const (
	defaultCapacity = 64
)

// Queue implements a queue by dynamic array.
// Is not thread safe.
type Queue struct {
	items  []interface{}
	cap    int
	front  int
	behind int
}

// Default creates a Queue with default parameters.
func Default() *Queue {
	return New(defaultCapacity)
}

// New creates a Queue with given initialization capacity.
func New(c int) *Queue {
	c += 1
	q := &Queue{
		items:  make([]interface{}, c),
		cap:    c,
		front:  0,
		behind: 0,
	}
	return q
}

// Len return the number of elements in the queue.
func (q *Queue) Len() int {
	return (q.behind - q.front + q.cap) % q.cap
}

// Cap return the current capacity of the queue.
func (q *Queue) Cap() int {
	return q.cap - 1
}

// Empty represents whether the queue is empty.
func (q *Queue) Empty() bool {
	return q.Len() == 0
}

// Push adds an element to the end of the queue.
func (q *Queue) Push(item interface{}) {
	q.autoGrow()
	q.items[q.behind] = item
	q.behind = (q.behind + 1) % q.cap
}

// Pop returns and removes an element that at the head.
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	item := q.items[q.front]
	q.items[q.front] = nil // Prevent memory leaks.
	q.front = (q.front + 1) % q.cap
	return item
}

// Peek returns the element that at the head.
func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return nil
	}
	return q.items[q.front]
}

func (q *Queue) autoGrow() {
	if q.Len() == q.Cap() {
		newCap := q.cap - 1
		if q.cap < 1024 {
			newCap += newCap
		} else {
			newCap += newCap / 2
		}
		q.grow(newCap)
	}
}

func (q *Queue) grow(c int) {
	if c > q.cap-1 {
		oldCap := q.cap
		oldLen := q.cap - 1
		items := q.items

		q.cap = c + 1
		q.items = make([]interface{}, q.cap)
		for i := 0; i < oldLen; i++ {
			q.items[i] = items[(i+q.front)%oldCap]
		}
		q.front = 0
		q.behind = oldLen
	}
}
