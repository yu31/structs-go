package container

// TreeNode declares an interface of Binary Search Tree Node.
type TreeNode interface {
	Element
	// Left returns the left child of the TreeNode.
	Left() TreeNode
	// Right returns the right child of the TreeNode.
	Right() TreeNode
}
