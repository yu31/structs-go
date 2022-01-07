// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package container

import (
	"bytes"
	"time"
)

var (
	_ Comparator = (*String)(nil)
	_ Comparator = (*Byte)(nil)
	_ Comparator = (*Rune)(nil)
	_ Comparator = (*Int)(nil)
	_ Comparator = (*Int8)(nil)
	_ Comparator = (*Int16)(nil)
	_ Comparator = (*Int32)(nil)
	_ Comparator = (*Int64)(nil)
	_ Comparator = (*Uint)(nil)
	_ Comparator = (*Uint8)(nil)
	_ Comparator = (*Uint16)(nil)
	_ Comparator = (*Uint32)(nil)
	_ Comparator = (*Uint64)(nil)
	_ Comparator = (*Bytes)(nil)
	_ Comparator = (*Duration)(nil)
	_ Comparator = (*Time)(nil)
)

// Comparator defines an interface of wraps any data-type of comparison two elements.
type Comparator interface {
	// Compare returns the comparison result with another Comparator.
	// The result will be 0 if source == target, -1 if source < target, and 1 if source > target.
	Compare(target Comparator) int
}

// String Wrapper for builtin type string.
type String string

func (k1 String) Compare(target Comparator) int {
	k2 := target.(String)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Byte Wrapper for builtin type byte.
type Byte byte

func (k1 Byte) Compare(target Comparator) int {
	k2 := target.(Byte)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Rune Wrapper for builtin type rune.
type Rune rune

func (k1 Rune) Compare(target Comparator) int {
	k2 := target.(Rune)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Int Wrapper for builtin type int.
type Int int

func (k1 Int) Compare(target Comparator) int {
	k2 := target.(Int)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Int8 Wrapper for builtin type int8.
type Int8 int8

func (k1 Int8) Compare(target Comparator) int {
	k2 := target.(Int8)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Int16 Wrapper for builtin type int16.
type Int16 int16

func (k1 Int16) Compare(target Comparator) int {
	k2 := target.(Int16)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Int32 Wrapper for builtin type int32.
type Int32 int32

func (k1 Int32) Compare(target Comparator) int {
	k2 := target.(Int32)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Int64 Wrapper for builtin type int64.
type Int64 int64

func (k1 Int64) Compare(target Comparator) int {
	k2 := target.(Int64)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Uint Wrapper for builtin type uint.
type Uint uint

func (k1 Uint) Compare(target Comparator) int {
	k2 := target.(Uint)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Uint8 Wrapper for builtin type uint8.
type Uint8 uint8

func (k1 Uint8) Compare(target Comparator) int {
	k2 := target.(Uint8)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Uint16 Wrapper for builtin type uint16.
type Uint16 uint16

func (k1 Uint16) Compare(target Comparator) int {
	k2 := target.(Uint16)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Uint32 Wrapper for builtin type uint32.
type Uint32 uint32

func (k1 Uint32) Compare(target Comparator) int {
	k2 := target.(Uint32)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Uint64 Wrapper for builtin type uint33.
type Uint64 uint64

func (k1 Uint64) Compare(target Comparator) int {
	k2 := target.(Uint64)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Bytes Wrapper for type byte slice.
type Bytes []byte

func (k1 Bytes) Compare(target Comparator) int {
	k2 := target.(Bytes)
	return bytes.Compare(k1, k2)
}

// Duration Wrapper for type time.Duration.
type Duration time.Duration

func (k1 Duration) Compare(target Comparator) int {
	k2 := target.(Duration)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

// Time Wrapper for type time.Time.
type Time time.Time

func (k1 Time) Compare(target Comparator) int {
	k2 := target.(Time)

	t1 := time.Time(k1)
	t2 := time.Time(k2)

	if t1.Before(t2) {
		return -1
	}
	if t1.After(t2) {
		return 1
	}
	return 0
}
