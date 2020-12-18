// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package container

// Key represents high-level Key type.
type Key = Comparator

// Value represents high-level Value type.
type Value interface{}

// Element is an element of a Container.
type Element interface {
	// Key returns the key that stored with this element.
	Key() Key

	// Value returns the value that stored with this element.
	Value() Value
}

// Searcher declares an interface to performs query in a Container.
type Searcher interface {
	// Range calls f sequentially each TreeNode present in the Container.
	// If f returns false, range stops the iteration.
	//
	// The range is start <= x < boundary.
	// The elements will return from the beginning if start is nil,
	// And return until the end if the boundary is nil.
	Range(start Key, boundary Key, f func(ele Element) bool)

	// Reverse is similar to the Range method. But it iteration element in reverse.
	// If f returns false, range stops the iteration.
	Reverse(start Key, boundary Key, f func(ele Element) bool)

	// LastLT searches for the last element that less than the key.
	LastLT(k Key) Element

	// LastLE search for the last element that less than or equal to the key.
	LastLE(k Key) Element

	// FirstGT search for the first element that greater than to the key.
	FirstGT(k Key) Element

	// FirstGE search for the first element that greater than or equal to the key.
	FirstGE(k Key) Element
}

// Container declares an data container interface.
type Container interface {
	Searcher

	// Len returns the number of elements.
	Len() int

	// Insert inserts a new element if the key doesn't exist, or returns the existing element for the key if present.
	// The bool result is true if an element was inserted, false if searched.
	Insert(k Key, v Value) (Element, bool)

	// Delete removes and returns the element of a given key.
	// Returns nil if key not found.
	Delete(k Key) Element

	// Update updates an element with the given key and value, And returns the old element of key.
	// Returns nil if the key not be found.
	Update(k Key, v Value) Element

	// Upsert inserts or updates an element by giving key and value.
	// The bool result is true if an element was inserted, false if an element was updated.
	//
	// The operation are same as the Insert method if key not found,
	// And are same as the Update method if key exists.
	Upsert(k Key, v Value) (Element, bool)

	// Search searches the element of a given key.
	// Returns nil if key not found.
	Search(k Key) Element

	// Iter creates an iterator to the iteration return element.
	//
	// The range is start <= x < boundary.
	// The elements will return from the beginning if start is nil,
	// And return until the end if the boundary is nil.
	Iter(start Key, boundary Key) Iterator
}

// The iterator is an interface for iteration return element.
type Iterator interface {
	// Valid represents whether to have more elements in the Iterator.
	// Returns false if no more.
	Valid() bool

	// Next returns a element and moved the iterator to the next element.
	// Returns nil if no more elements.
	Next() Element
}
