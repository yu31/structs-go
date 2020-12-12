package skip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

func TestList_Iter(t *testing.T) {
	sl := New()

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	seeds := []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

	for _, k := range seeds {
		sl.Insert(k, int64(k*2+1))
	}

	var iter container.Iterator
	var element Element

	/* ------ test start == nil && boundary == nil */

	iter = sl.Iter(nil, nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary == nil --- */
	//
	// start < first node
	iter = sl.Iter(container.Int64(21), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start == first node
	iter = sl.Iter(container.Int64(22), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(27), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 2; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(62), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 4; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start > root node && start < last node
	iter = sl.Iter(container.Int64(132), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 12; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start == last node
	iter = sl.Iter(container.Int64(150), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	element = iter.Next()
	require.Equal(t, element.Key(), container.Int64(150))
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start > last node
	iter = sl.Iter(container.Int64(156), nil)
	require.NotNil(t, iter)
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	/* ---  test start == nil && boundary != nil --- */
	//
	// boundary < first node
	iter = sl.Iter(nil, container.Int64(21))
	require.NotNil(t, iter)
	require.False(t, iter.Valid())
	element = iter.Next()
	require.Nil(t, element)

	// boundary == first node
	iter = sl.Iter(nil, container.Int64(22))
	require.NotNil(t, iter)
	require.False(t, iter.Valid())
	element = iter.Next()
	require.Nil(t, element)

	// boundary > first node
	iter = sl.Iter(nil, container.Int64(24))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	element = iter.Next()
	require.Equal(t, element.Key(), container.Int64(22))
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// boundary < last node && bound > first node
	iter = sl.Iter(nil, container.Int64(147))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds)-1; i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// boundary == last node
	iter = sl.Iter(nil, container.Int64(150))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds)-1; i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// boundary > last node
	iter = sl.Iter(nil, container.Int64(156))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := range seeds {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary != nil --- */
	//

	// start < boundary && start > first node && bound < last node
	iter = sl.Iter(container.Int64(68), container.Int64(132))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 6; i < len(seeds)-3; i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound > last node
	iter = sl.Iter(container.Int64(21), container.Int64(153))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		element := iter.Next()
		require.NotNil(t, element, fmt.Sprintf("key %v not found", seeds[i]))
		require.Equal(t, element.Key(), seeds[i])
		require.Equal(t, element.Value(), int64(seeds[i]*2+1))
	}
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start == boundary, start and boundary exists.
	iter = sl.Iter(container.Int64(24), container.Int64(24))
	require.NotNil(t, iter)
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start == boundary, start and boundary not exists.
	iter = sl.Iter(container.Int64(25), container.Int64(25))
	require.NotNil(t, iter)
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound < first node
	iter = sl.Iter(container.Int64(21), container.Int64(13))
	require.NotNil(t, iter)
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())

	// start > boundary && start > first node
	iter = sl.Iter(container.Int64(65), container.Int64(27))
	require.NotNil(t, iter)
	element = iter.Next()
	require.Nil(t, element)
	require.False(t, iter.Valid())
}
