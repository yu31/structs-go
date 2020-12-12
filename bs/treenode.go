// Copyright (c) 2019, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package bs

import (
	"github.com/yu31/gostructs/container"
)

// Type aliases for simplifying use in this package.
type Key = container.Key
type Value = container.Value
type Element = container.Element

// TreeNode is a universal Binary Search Tree Node type.
type TreeNode interface {
	Element
	// Left returns the left child of the TreeNode.
	Left() TreeNode
	// Right returns the right child of the TreeNode.
	Right() TreeNode
}
