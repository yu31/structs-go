package bs

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

func searchRange(root TreeNode, start Key, boundary Key) []Element {
	var result []Element

	Range(root, start, boundary, func(n TreeNode) bool {
		result = append(result, n)
		return true
	})

	return result
}

func searchRangeByRecursion(root *treeNode, start Key, boundary Key) []Element {
	if root == nil {
		return nil
	}

	var result []Element

	var recursion func(n *treeNode, start Key, boundary Key)

	recursion = func(n *treeNode, start Key, boundary Key) {
		if n == nil {
			return
		}
		if start != nil && n.key.Compare(start) == -1 {
			recursion(n.right, start, boundary)
		} else if boundary != nil && n.key.Compare(boundary) != -1 {
			recursion(n.left, start, boundary)
		} else {
			// start <= node <= boundary
			recursion(n.left, start, boundary)
			result = append(result, n)
			recursion(n.right, start, boundary)
		}
	}

	recursion(root, start, boundary)

	return result
}

func searchRangeByIter(root *treeNode, start Key, boundary Key) []Element {
	if root == nil {
		return nil
	}

	var result []Element

	it := NewIterator(root, start, boundary)
	for it.Valid() {
		n := it.Next()
		result = append(result, n)
	}
	return result
}

func TestSearchRange(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	var r1, r2, r3 []Element

	/* ------ test start == nil && boundary == nil */

	r1 = searchRange(nil, nil, nil)
	r2 = searchRangeByRecursion(nil, nil, nil)
	r3 = searchRangeByIter(nil, nil, nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, nil)
	r2 = searchRangeByRecursion(tr.root, nil, nil)
	r3 = searchRangeByIter(tr.root, nil, nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start != nil && boundary == nil --- */

	r1 = searchRange(tr.root, container.Int64(21), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(21), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(22), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(22), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(22), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(27), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(27), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(27), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(62), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(62), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(62), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(132), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(132), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(132), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(144), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(144), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(144), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(150), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(150), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(150), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(156), nil)
	r2 = searchRangeByRecursion(tr.root, container.Int64(156), nil)
	r3 = searchRangeByIter(tr.root, container.Int64(156), nil)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start == nil && boundary != nil --- */

	r1 = searchRange(tr.root, nil, container.Int64(21))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(21))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(21))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, container.Int64(22))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(22))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(22))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, container.Int64(77))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(77))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(77))
	require.Equal(t, len(r1), 7)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, container.Int64(147))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(147))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(147))
	require.Equal(t, len(r1), 14)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, container.Int64(150))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(150))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(150))
	require.Equal(t, len(r1), 14)
	require.Equal(t, r1[len(r1)-1].Key(), container.Int64(145))
	require.Equal(t, r1[len(r1)-1].Value(), int64(145*2+1))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, nil, container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, nil, container.Int64(156))
	r3 = searchRangeByIter(tr.root, nil, container.Int64(156))
	require.Equal(t, len(r1), 15)
	require.Equal(t, r1[len(r1)-1].Key(), container.Int64(150))
	require.Equal(t, r1[len(r1)-1].Value(), int64(150*2+1))
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	/* ---  test start != nil && boundary == nil --- */

	r1 = searchRange(tr.root, container.Int64(21), container.Int64(13))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(13))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(13))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(65), container.Int64(27))
	r2 = searchRangeByRecursion(tr.root, container.Int64(65), container.Int64(27))
	r3 = searchRangeByIter(tr.root, container.Int64(65), container.Int64(27))
	require.Equal(t, len(r1), 0)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(68), container.Int64(132))
	r2 = searchRangeByRecursion(tr.root, container.Int64(68), container.Int64(132))
	r3 = searchRangeByIter(tr.root, container.Int64(68), container.Int64(132))
	require.Equal(t, len(r1), 6)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)

	r1 = searchRange(tr.root, container.Int64(21), container.Int64(156))
	r2 = searchRangeByRecursion(tr.root, container.Int64(21), container.Int64(156))
	r3 = searchRangeByIter(tr.root, container.Int64(21), container.Int64(156))
	require.Equal(t, len(r1), 15)
	require.Equal(t, r1, r2)
	require.Equal(t, r2, r3)
}

func TestSearchLastLT(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = LastLT(tr.root, container.Int64(21))
	require.Nil(t, element)

	element = LastLT(tr.root, container.Int64(22))
	require.Nil(t, element)

	element = LastLT(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(24))
	require.Equal(t, element.Value(), int64(24*2+1))

	element = LastLT(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(61))
	require.Equal(t, element.Value(), int64(61*2+1))

	element = LastLT(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(76))
	require.Equal(t, element.Value(), int64(76*2+1))

	element = LastLT(tr.root, container.Int64(84))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(76))
	require.Equal(t, element.Value(), int64(76*2+1))

	element = LastLT(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(97))
	require.Equal(t, element.Value(), int64(97*2+1))

	element = LastLT(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(130))
	require.Equal(t, element.Value(), int64(130*2+1))

	element = LastLT(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(130))
	require.Equal(t, element.Value(), int64(130*2+1))

	element = LastLT(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(145))
	require.Equal(t, element.Value(), int64(145*2+1))

	element = LastLT(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(145))
	require.Equal(t, element.Value(), int64(145*2+1))

	element = LastLT(tr.root, container.Int64(156))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))
}

func TestSearchLastLE(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = LastLE(tr.root, container.Int64(21))
	require.Nil(t, element)

	element = LastLE(tr.root, container.Int64(22))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(22))
	require.Equal(t, element.Value(), int64(22*2+1))

	element = LastLE(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(24))
	require.Equal(t, element.Value(), int64(24*2+1))

	element = LastLE(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(61))
	require.Equal(t, element.Value(), int64(61*2+1))

	element = LastLE(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(76))
	require.Equal(t, element.Value(), int64(76*2+1))

	element = LastLE(tr.root, container.Int64(76))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(76))
	require.Equal(t, element.Value(), int64(76*2+1))

	element = LastLE(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(97))
	require.Equal(t, element.Value(), int64(97*2+1))

	element = LastLE(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(130))
	require.Equal(t, element.Value(), int64(130*2+1))

	element = LastLE(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(133))
	require.Equal(t, element.Value(), int64(133*2+1))

	element = LastLE(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(145))
	require.Equal(t, element.Value(), int64(145*2+1))

	element = LastLE(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))

	element = LastLE(tr.root, container.Int64(156))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))
}

func TestSearchFirstGT(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = FirstGT(tr.root, container.Int64(21))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(22))
	require.Equal(t, element.Value(), int64(22*2+1))

	element = FirstGT(tr.root, container.Int64(24))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(35))
	require.Equal(t, element.Value(), int64(35*2+1))

	element = FirstGT(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(35))
	require.Equal(t, element.Value(), int64(35*2+1))

	element = FirstGT(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(64))
	require.Equal(t, element.Value(), int64(64*2+1))

	element = FirstGT(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(84))
	require.Equal(t, element.Value(), int64(84*2+1))

	element = FirstGT(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(130))
	require.Equal(t, element.Value(), int64(130*2+1))

	element = FirstGT(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(133))
	require.Equal(t, element.Value(), int64(133*2+1))

	element = FirstGT(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(145))
	require.Equal(t, element.Value(), int64(145*2+1))

	element = FirstGT(tr.root, container.Int64(147))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))

	element = FirstGT(tr.root, container.Int64(150))
	require.Nil(t, element)
	element = FirstGT(tr.root, container.Int64(151))
	require.Nil(t, element)
}

func TestSearchFirstGE(t *testing.T) {
	tr := New()

	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150, 87, 97, 22, 35, 64, 76}

	// seeds sequence in tree by in order: 22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	for _, k := range seeds {
		tr.Insert(k, int64(k*2+1))
	}

	var element Element

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	/*
		                                 84
		              61 			                        130
			 24		            67	               91	  	           145
		22        35       64        76       87        97       133         150
	*/

	element = FirstGE(tr.root, container.Int64(21))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(22))
	require.Equal(t, element.Value(), int64(22*2+1))

	element = FirstGE(tr.root, container.Int64(24))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(24))
	require.Equal(t, element.Value(), int64(24*2+1))

	element = FirstGE(tr.root, container.Int64(25))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(35))
	require.Equal(t, element.Value(), int64(35*2+1))

	element = FirstGE(tr.root, container.Int64(63))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(64))
	require.Equal(t, element.Value(), int64(64*2+1))

	element = FirstGE(tr.root, container.Int64(77))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(84))
	require.Equal(t, element.Value(), int64(84*2+1))

	element = FirstGE(tr.root, container.Int64(99))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(130))
	require.Equal(t, element.Value(), int64(130*2+1))

	element = FirstGE(tr.root, container.Int64(132))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(133))
	require.Equal(t, element.Value(), int64(133*2+1))

	element = FirstGE(tr.root, container.Int64(133))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(133))
	require.Equal(t, element.Value(), int64(133*2+1))

	element = FirstGE(tr.root, container.Int64(146))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))

	element = FirstGE(tr.root, container.Int64(150))
	require.NotNil(t, element)
	require.Equal(t, element.Key(), container.Int64(150))
	require.Equal(t, element.Value(), int64(150*2+1))

	element = FirstGE(tr.root, container.Int64(151))
	require.Nil(t, element)
}
