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

	// Insert inserts and returns an Element with the given key and value.
	// Returns nil if key already exists.
	//
	// NOTICE: This method should not allow inserts with an exists key.
	// You can use the Replace method if the action you expect is "insert or update".
	Insert(k Key, v Value) Element

	// Delete removes and returns the Element of a given key.
	// Returns nil if key not found.
	Delete(k Key) Element

	// Update updates and returns an Element with the given key and value.
	// Returns nil if key not found.
	//
	// NOTICE: This method should not allow updates with an not exist key.
	// You can use the Replace method if the action you expect is "insert or update".
	Update(k Key, v Value) Element

	// Replace inserts or updates an Element by giving key and value.
	//
	// The action are same as an Insert method if key not found,
	// And are same as an Update method if found the key.
	Replace(k Key, v Value) Element

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
