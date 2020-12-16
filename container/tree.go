// Copyright (c) 2020, Yu Wu <yu.771991@gmail.com> All rights reserved.
//
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package container

// TreeNode declares an interface of Binary Search Tree Node.
type TreeNode interface {
	Element
	// Left returns the left child of the TreeNode.
	Left() TreeNode
	// Right returns the right child of the TreeNode.
	Right() TreeNode
}

// Tree declares an interface of Binary Search Tree.
type Tree interface {
	Container
	// Root returns the root node of the tree.
	Root() TreeNode
}
