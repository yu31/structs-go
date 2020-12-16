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

// Container declares an data container interface.
type Container interface {
	// Len returns the number of elements.
	Len() int

	// Insert inserts and returns an Element with given key and value if key doesn't exists.
	// Or else, returns the existing Element for the key if present.
	// The bool result is true if an Element was inserted, false if searched.
	Insert(k Key, v Value) (Element, bool)

	// Delete removes and returns the Element of a given key.
	// Returns nil if key not found.
	Delete(k Key) Element

	// Update updates an Element with the given key and value, And returns the old element.
	// Returns nil if the key not be found.
	Update(k Key, v Value) Element

	// Replace inserts or updates an Element by giving key and value.
	// The bool result is true if an Element was inserted, false if an Element was updated.
	//
	// The operation are same as the Insert method if key not found,
	// And are same as the Update method if key exists.
	Replace(k Key, v Value) (Element, bool)

	// Search searches the Element of a given key.
	// Returns nil if key not found.
	Search(k Key) Element

	// Iter creates an iterator to the iteration return element.
	//
	// The elements range is start <= x < boundary.
	// The elements will return from the beginning if start is nil,
	// And return until the end if the boundary is nil.
	Iter(start Key, boundary Key) Iterator
}

// The iterator is an interface for iteration return element.
type Iterator interface {
	// Valid represents whether to have more elements in the Iterator.
	Valid() bool

	// Next returns a Element and moved the iterator to the next Element.
	// Returns nil if no more elements.
	Next() Element
}
