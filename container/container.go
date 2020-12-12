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

// Container defines an data container interface.
type Container interface {
	// Len returns the number of elements.
	Len() int

	// Insert inserts the giving key and value as an Element and return.
	// Returns nil if key already exists.
	Insert(k Key, v Value) (element Element)

	// Delete removes and returns the Element of a given key.
	// Returns nil if not found.
	Delete(k Key) (element Element)

	// Search returns the Element of a given key.
	// Returns nil if not found.
	Search(k Key) (element Element)

	// Iter creates an iterator to iteration return element.
	//
	// The element range is start <= x < boundary.
	// The element will return from the beginning if start is nil,
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
