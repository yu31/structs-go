package container

type Searcher interface {
	// Range calls f sequentially each TreeNode present in the Tree.
	// If f returns false, range stops the iteration.
	//
	// The elements range is start <= x < boundary.
	// The elements will return from the beginning if start is nil,
	// And return until the end if the boundary is nil.
	Range(start Key, boundary Key, f func(ele Element) bool)

	// LastLT searches for the last node that less than the key.
	LastLT(k Key) Element

	// LastLE search for the last node that less than or equal to the key.
	LastLE(k Key) Element

	// FirstGT search for the first node that greater than to the key.
	FirstGT(k Key) Element

	// FirstGE search for the first node that greater than or equal to the key.
	FirstGE(k Key) Element
}
