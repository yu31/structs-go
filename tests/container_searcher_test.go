package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/structs-go/container"
)

// --------- order in container: [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
var searchSeeds = []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

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
