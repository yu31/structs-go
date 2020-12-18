package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

// --------- order in container: [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
var searchSeeds = []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

func TestContainerSearcher_Range(t *testing.T) {
	seeds := searchSeeds
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

func TestContainerSearcher_Reverse(t *testing.T) {
	seeds := searchSeeds
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

func TestContainerSearcher_LastLT(t *testing.T) {
	seeds := searchSeeds
	process := func(ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var element container.Element

		element = ctr.LastLT(container.Int64(21))
		require.Nil(t, element)

		element = ctr.LastLT(container.Int64(22))
		require.Nil(t, element)

		element = ctr.LastLT(container.Int64(25))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(24))
		require.Equal(t, element.Value(), int64(24*2+1))

		element = ctr.LastLT(container.Int64(63))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(61))
		require.Equal(t, element.Value(), int64(61*2+1))

		element = ctr.LastLT(container.Int64(77))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(76))
		require.Equal(t, element.Value(), int64(76*2+1))

		element = ctr.LastLT(container.Int64(84))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(76))
		require.Equal(t, element.Value(), int64(76*2+1))

		element = ctr.LastLT(container.Int64(99))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(97))
		require.Equal(t, element.Value(), int64(97*2+1))

		element = ctr.LastLT(container.Int64(132))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(130))
		require.Equal(t, element.Value(), int64(130*2+1))

		element = ctr.LastLT(container.Int64(133))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(130))
		require.Equal(t, element.Value(), int64(130*2+1))

		element = ctr.LastLT(container.Int64(146))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(145))
		require.Equal(t, element.Value(), int64(145*2+1))

		element = ctr.LastLT(container.Int64(150))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(145))
		require.Equal(t, element.Value(), int64(145*2+1))

		element = ctr.LastLT(container.Int64(156))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerSearcher_LastLE(t *testing.T) {
	seeds := searchSeeds
	process := func(ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var element container.Element

		element = ctr.LastLE(container.Int64(21))
		require.Nil(t, element)

		element = ctr.LastLE(container.Int64(22))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(22))
		require.Equal(t, element.Value(), int64(22*2+1))

		element = ctr.LastLE(container.Int64(25))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(24))
		require.Equal(t, element.Value(), int64(24*2+1))

		element = ctr.LastLE(container.Int64(63))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(61))
		require.Equal(t, element.Value(), int64(61*2+1))

		element = ctr.LastLE(container.Int64(77))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(76))
		require.Equal(t, element.Value(), int64(76*2+1))

		element = ctr.LastLE(container.Int64(76))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(76))
		require.Equal(t, element.Value(), int64(76*2+1))

		element = ctr.LastLE(container.Int64(99))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(97))
		require.Equal(t, element.Value(), int64(97*2+1))

		element = ctr.LastLE(container.Int64(132))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(130))
		require.Equal(t, element.Value(), int64(130*2+1))

		element = ctr.LastLE(container.Int64(133))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(133))
		require.Equal(t, element.Value(), int64(133*2+1))

		element = ctr.LastLE(container.Int64(146))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(145))
		require.Equal(t, element.Value(), int64(145*2+1))

		element = ctr.LastLE(container.Int64(150))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))

		element = ctr.LastLE(container.Int64(156))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerSearcher_FirstGT(t *testing.T) {
	seeds := searchSeeds
	process := func(ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var element container.Element

		element = ctr.FirstGT(container.Int64(21))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(22))
		require.Equal(t, element.Value(), int64(22*2+1))

		element = ctr.FirstGT(container.Int64(24))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(35))
		require.Equal(t, element.Value(), int64(35*2+1))

		element = ctr.FirstGT(container.Int64(25))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(35))
		require.Equal(t, element.Value(), int64(35*2+1))

		element = ctr.FirstGT(container.Int64(63))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(64))
		require.Equal(t, element.Value(), int64(64*2+1))

		element = ctr.FirstGT(container.Int64(77))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(84))
		require.Equal(t, element.Value(), int64(84*2+1))

		element = ctr.FirstGT(container.Int64(99))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(130))
		require.Equal(t, element.Value(), int64(130*2+1))

		element = ctr.FirstGT(container.Int64(132))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(133))
		require.Equal(t, element.Value(), int64(133*2+1))

		element = ctr.FirstGT(container.Int64(133))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(145))
		require.Equal(t, element.Value(), int64(145*2+1))

		element = ctr.FirstGT(container.Int64(147))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))

		element = ctr.FirstGT(container.Int64(150))
		require.Nil(t, element)
		element = ctr.FirstGT(container.Int64(151))
		require.Nil(t, element)
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}

func TestContainerSearcher_FirstGE(t *testing.T) {
	seeds := searchSeeds
	process := func(ctr container.Container) {
		// Insert seeds in random order
		for _, k := range shuffleSeeds(seeds) {
			ctr.Insert(k, int64(k*2+1))
		}

		var element container.Element

		element = ctr.FirstGE(container.Int64(21))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(22))
		require.Equal(t, element.Value(), int64(22*2+1))

		element = ctr.FirstGE(container.Int64(24))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(24))
		require.Equal(t, element.Value(), int64(24*2+1))

		element = ctr.FirstGE(container.Int64(25))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(35))
		require.Equal(t, element.Value(), int64(35*2+1))

		element = ctr.FirstGE(container.Int64(63))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(64))
		require.Equal(t, element.Value(), int64(64*2+1))

		element = ctr.FirstGE(container.Int64(77))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(84))
		require.Equal(t, element.Value(), int64(84*2+1))

		element = ctr.FirstGE(container.Int64(99))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(130))
		require.Equal(t, element.Value(), int64(130*2+1))

		element = ctr.FirstGE(container.Int64(132))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(133))
		require.Equal(t, element.Value(), int64(133*2+1))

		element = ctr.FirstGE(container.Int64(133))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(133))
		require.Equal(t, element.Value(), int64(133*2+1))

		element = ctr.FirstGE(container.Int64(146))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))

		element = ctr.FirstGE(container.Int64(150))
		require.NotNil(t, element)
		require.Equal(t, element.Key(), container.Int64(150))
		require.Equal(t, element.Value(), int64(150*2+1))

		element = ctr.FirstGE(container.Int64(151))
		require.Nil(t, element)
	}

	// Test for all container implementation.
	for name, f := range containers {
		t.Run(name, func(t *testing.T) {
			process(f())
		})
	}
}
