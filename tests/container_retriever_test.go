package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

// positive order in container: [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150]
// reverse order in container:  [150, 145, 133, 130, 97, 91, 87, 84, 76, 67, 64, 61, 35, 24, 22]
var retrieverSeeds = []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

func TestContainerRetriever_Range(t *testing.T) {
	seeds := retrieverSeeds

	process := func(t *testing.T, ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var r1, r2 []container.Element

		// Test case start == nil && boundary == nil
		t.Run("case1", func(t *testing.T) {
			r1 = searchRange(ctr, nil, nil)
			r2 = searchRangeByIter(ctr, nil, nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, nil)
			r2 = searchRangeByIter(ctr, nil, nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, r1, r2)
		})

		// Test case start != nil && boundary == nil
		t.Run("case2", func(t *testing.T) {
			r1 = searchRange(ctr, container.Int64(21), nil)
			r2 = searchRangeByIter(ctr, container.Int64(21), nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(22), nil)
			r2 = searchRangeByIter(ctr, container.Int64(22), nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(27), nil)
			r2 = searchRangeByIter(ctr, container.Int64(27), nil)
			require.Equal(t, len(r1), len(searchSeeds)-2)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(62), nil)
			r2 = searchRangeByIter(ctr, container.Int64(62), nil)
			require.Equal(t, len(r1), len(searchSeeds)-4)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(132), nil)
			r2 = searchRangeByIter(ctr, container.Int64(132), nil)
			require.Equal(t, len(r1), 3)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(144), nil)
			r2 = searchRangeByIter(ctr, container.Int64(144), nil)
			require.Equal(t, len(r1), 2)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(150), nil)
			r2 = searchRangeByIter(ctr, container.Int64(150), nil)
			require.Equal(t, len(r1), 1)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(156), nil)
			r2 = searchRangeByIter(ctr, container.Int64(156), nil)
			require.Equal(t, len(r1), 0)
			require.Equal(t, r1, r2)
		})

		// Test case start == nil && boundary != nil
		t.Run("case3", func(t *testing.T) {
			r1 = searchRange(ctr, nil, container.Int64(21))
			r2 = searchRangeByIter(ctr, nil, container.Int64(21))
			require.Equal(t, len(r1), 0)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, container.Int64(22))
			r2 = searchRangeByIter(ctr, nil, container.Int64(22))
			require.Equal(t, len(r1), 0)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, container.Int64(77))
			r2 = searchRangeByIter(ctr, nil, container.Int64(77))
			require.Equal(t, len(r1), 7)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, container.Int64(147))
			r2 = searchRangeByIter(ctr, nil, container.Int64(147))
			require.Equal(t, len(r1), 14)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, container.Int64(150))
			r2 = searchRangeByIter(ctr, nil, container.Int64(150))
			require.Equal(t, len(r1), 14)
			require.Equal(t, r1[len(r1)-1].Key(), container.Int64(145))
			require.Equal(t, r1[len(r1)-1].Value(), int64(145*2+1))
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, nil, container.Int64(156))
			r2 = searchRangeByIter(ctr, nil, container.Int64(156))
			require.Equal(t, len(r1), 15)
			require.Equal(t, r1[len(r1)-1].Key(), container.Int64(150))
			require.Equal(t, r1[len(r1)-1].Value(), int64(150*2+1))
			require.Equal(t, r1, r2)
		})

		// Test case start != nil && boundary == nil
		t.Run("case4", func(t *testing.T) {
			r1 = searchRange(ctr, container.Int64(21), container.Int64(13))
			r2 = searchRangeByIter(ctr, container.Int64(21), container.Int64(13))
			require.Equal(t, len(r1), 0)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(65), container.Int64(27))
			r2 = searchRangeByIter(ctr, container.Int64(65), container.Int64(27))
			require.Equal(t, len(r1), 0)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(68), container.Int64(132))
			r2 = searchRangeByIter(ctr, container.Int64(68), container.Int64(132))
			require.Equal(t, len(r1), 6)
			require.Equal(t, r1, r2)

			r1 = searchRange(ctr, container.Int64(21), container.Int64(156))
			r2 = searchRangeByIter(ctr, container.Int64(21), container.Int64(156))
			require.Equal(t, len(r1), 15)
			require.Equal(t, r1, r2)
		})
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(t, f())
		})
	}
}

func TestContainerRetriever_Reverse(t *testing.T) {
	seeds := retrieverSeeds

	process := func(t *testing.T, ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var r1, r2 []container.Element

		// Test case start == nil && boundary == nil
		t.Run("case1", func(t *testing.T) {
			r1 = searchRange(ctr, nil, nil)
			r2 = searchReceive(ctr, nil, nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, nil)
			r2 = searchReceive(ctr, nil, nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, reverseElementSlice(r1), r2)
		})

		// Test case start != nil && boundary == nil
		t.Run("case2", func(t *testing.T) {
			r1 = searchRange(ctr, container.Int64(21), nil)
			r2 = searchReceive(ctr, container.Int64(21), nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(22), nil)
			r2 = searchReceive(ctr, container.Int64(22), nil)
			require.Equal(t, len(r1), len(searchSeeds))
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(27), nil)
			r2 = searchReceive(ctr, container.Int64(27), nil)
			require.Equal(t, len(r1), len(searchSeeds)-2)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(62), nil)
			r2 = searchReceive(ctr, container.Int64(62), nil)
			require.Equal(t, len(r1), len(searchSeeds)-4)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(132), nil)
			r2 = searchReceive(ctr, container.Int64(132), nil)
			require.Equal(t, len(r1), 3)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(144), nil)
			r2 = searchReceive(ctr, container.Int64(144), nil)
			require.Equal(t, len(r1), 2)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(150), nil)
			r2 = searchReceive(ctr, container.Int64(150), nil)
			require.Equal(t, len(r1), 1)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(156), nil)
			r2 = searchReceive(ctr, container.Int64(156), nil)
			require.Equal(t, len(r1), 0)
			require.Equal(t, reverseElementSlice(r1), r2)
		})

		// Test case start == nil && boundary != nil
		t.Run("case3", func(t *testing.T) {
			r1 = searchRange(ctr, nil, container.Int64(21))
			r2 = searchReceive(ctr, nil, container.Int64(21))
			require.Equal(t, len(r1), 0)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, container.Int64(22))
			r2 = searchReceive(ctr, nil, container.Int64(22))
			require.Equal(t, len(r1), 0)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, container.Int64(77))
			r2 = searchReceive(ctr, nil, container.Int64(77))
			require.Equal(t, len(r1), 7)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, container.Int64(147))
			r2 = searchReceive(ctr, nil, container.Int64(147))
			require.Equal(t, len(r1), 14)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, container.Int64(150))
			r2 = searchReceive(ctr, nil, container.Int64(150))
			require.Equal(t, len(r1), 14)
			require.Equal(t, r1[len(r1)-1].Key(), container.Int64(145))
			require.Equal(t, r1[len(r1)-1].Value(), int64(145*2+1))
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, nil, container.Int64(156))
			r2 = searchReceive(ctr, nil, container.Int64(156))
			require.Equal(t, len(r1), 15)
			require.Equal(t, r1[len(r1)-1].Key(), container.Int64(150))
			require.Equal(t, r1[len(r1)-1].Value(), int64(150*2+1))
			require.Equal(t, reverseElementSlice(r1), r2)
		})

		// Test case start != nil && boundary == nil
		t.Run("case4", func(t *testing.T) {
			r1 = searchRange(ctr, container.Int64(21), container.Int64(13))
			r2 = searchReceive(ctr, container.Int64(21), container.Int64(13))
			require.Equal(t, len(r1), 0)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(65), container.Int64(27))
			r2 = searchReceive(ctr, container.Int64(65), container.Int64(27))
			require.Equal(t, len(r1), 0)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(68), container.Int64(132))
			r2 = searchReceive(ctr, container.Int64(68), container.Int64(132))
			require.Equal(t, len(r1), 6)
			require.Equal(t, reverseElementSlice(r1), r2)

			r1 = searchRange(ctr, container.Int64(21), container.Int64(156))
			r2 = searchReceive(ctr, container.Int64(21), container.Int64(156))
			require.Equal(t, len(r1), 15)
			require.Equal(t, reverseElementSlice(r1), r2)
		})
	}

	// Test for all container implementation.
	for name, f := range containers {
		// TODO: skiplist not implemented on now.
		if name == "skiplist" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			process(t, f())
		})
	}
}

func TestContainerRetriever_Iter(t *testing.T) {
	seeds := retrieverSeeds

	process := func(t *testing.T, ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var iter container.Iterator
		var element container.Element

		// Test case: start == nil and boundary == nil
		t.Run("case1", func(t *testing.T) {
			iter = ctr.Iter(nil, nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start != nil && boundary == nil
		t.Run("case2", func(t *testing.T) {
			// start < first node
			iter = ctr.Iter(container.Int64(21), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == first node
			iter = ctr.Iter(container.Int64(22), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > first node && start < last node
			iter = ctr.Iter(container.Int64(27), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 2; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > first node && start < last node
			iter = ctr.Iter(container.Int64(62), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 4; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > root node && start < last node
			iter = ctr.Iter(container.Int64(132), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 12; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == last node
			iter = ctr.Iter(container.Int64(150), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			element = iter.Next()
			require.Equal(t, element.Key(), container.Int64(150))
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > last node
			iter = ctr.Iter(container.Int64(156), nil)
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start == nil && boundary != nil
		t.Run("case3", func(t *testing.T) {
			// boundary < first node
			iter = ctr.Iter(nil, container.Int64(21))
			require.NotNil(t, iter)
			require.False(t, iter.Valid())
			element = iter.Next()
			require.Nil(t, element)

			// boundary == first node
			iter = ctr.Iter(nil, container.Int64(22))
			require.NotNil(t, iter)
			require.False(t, iter.Valid())
			element = iter.Next()
			require.Nil(t, element)

			// boundary > first node
			iter = ctr.Iter(nil, container.Int64(24))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			element = iter.Next()
			require.Equal(t, element.Key(), container.Int64(22))
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary < last node && bound > first node
			iter = ctr.Iter(nil, container.Int64(147))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds)-1; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary == last node
			iter = ctr.Iter(nil, container.Int64(150))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds)-1; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary > last node
			iter = ctr.Iter(nil, container.Int64(156))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := range seeds {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start != nil && boundary != nil
		t.Run("case4", func(t *testing.T) {
			// start < boundary && start > first node && bound < last node
			iter = ctr.Iter(container.Int64(68), container.Int64(132))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 6; i < len(seeds)-3; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start < boundary && start < first node && bound > last node
			iter = ctr.Iter(container.Int64(21), container.Int64(153))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == boundary, start and boundary exists.
			iter = ctr.Iter(container.Int64(24), container.Int64(24))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == boundary, start and boundary not exists.
			iter = ctr.Iter(container.Int64(25), container.Int64(25))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start < boundary && start < first node && bound < first node
			iter = ctr.Iter(container.Int64(21), container.Int64(13))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > boundary && start > first node
			iter = ctr.Iter(container.Int64(65), container.Int64(27))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(t, f())
		})
	}
}

func TestContainerRetriever_IterReverse(t *testing.T) {
	seeds := reverseSeedsSlice(retrieverSeeds)

	process := func(t *testing.T, ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var iter container.Iterator
		var element container.Element

		// Test case: start == nil and boundary == nil
		t.Run("case1", func(t *testing.T) {
			iter = ctr.IterReverse(nil, nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start != nil && boundary == nil
		t.Run("case2", func(t *testing.T) {
			// start < first node
			iter = ctr.IterReverse(container.Int64(21), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == first node
			iter = ctr.IterReverse(container.Int64(22), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > first node && start < last node
			iter = ctr.IterReverse(container.Int64(27), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds)-2; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > first node && start < last node
			iter = ctr.IterReverse(container.Int64(62), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds)-4; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > root node && start < last node
			iter = ctr.IterReverse(container.Int64(132), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds)-12; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == last node
			iter = ctr.IterReverse(container.Int64(150), nil)
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			element = iter.Next()
			require.Equal(t, element.Key(), container.Int64(150))
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > last node
			iter = ctr.IterReverse(container.Int64(156), nil)
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start == nil && boundary != nil
		t.Run("case3", func(t *testing.T) {
			// boundary < first node
			iter = ctr.IterReverse(nil, container.Int64(21))
			require.NotNil(t, iter)
			require.False(t, iter.Valid())
			element = iter.Next()
			require.Nil(t, element)

			// boundary == first node
			iter = ctr.IterReverse(nil, container.Int64(22))
			require.NotNil(t, iter)
			require.False(t, iter.Valid())
			element = iter.Next()
			require.Nil(t, element)

			// boundary > first node
			iter = ctr.IterReverse(nil, container.Int64(24))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			element = iter.Next()
			require.Equal(t, element.Key(), container.Int64(22))
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary < last node && bound > first node
			iter = ctr.IterReverse(nil, container.Int64(147))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 1; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary == last node
			iter = ctr.IterReverse(nil, container.Int64(150))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 1; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// boundary > last node
			iter = ctr.IterReverse(nil, container.Int64(156))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := range seeds {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})

		// Test case: start != nil && boundary != nil
		t.Run("case4", func(t *testing.T) {
			// start < boundary && start > first node && bound < last node
			iter = ctr.IterReverse(container.Int64(68), container.Int64(132))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 3; i < len(seeds)-6; i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start < boundary && start < first node && bound > last node
			iter = ctr.IterReverse(container.Int64(21), container.Int64(153))
			require.NotNil(t, iter)
			require.True(t, iter.Valid())
			for i := 0; i < len(seeds); i++ {
				element = iter.Next()
				require.NotNil(t, element, "key: %v", seeds[i])
				require.Equal(t, element.Key(), seeds[i])
				require.Equal(t, element.Value(), int64(seeds[i]*2+1))
			}
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == boundary, start and boundary exists.
			iter = ctr.IterReverse(container.Int64(24), container.Int64(24))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start == boundary, start and boundary not exists.
			iter = ctr.IterReverse(container.Int64(25), container.Int64(25))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start < boundary && start < first node && bound < first node
			iter = ctr.IterReverse(container.Int64(21), container.Int64(13))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())

			// start > boundary && start > first node
			iter = ctr.IterReverse(container.Int64(65), container.Int64(27))
			require.NotNil(t, iter)
			element = iter.Next()
			require.Nil(t, element)
			require.False(t, iter.Valid())
		})
	}

	// Test for all container implementation.
	for name, f := range containers {
		// TODO: skiplist not implemented on now.
		if name == "skiplist" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			process(t, f())
		})
	}
}
