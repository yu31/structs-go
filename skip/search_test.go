package skip

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/gostructs/container"
)

func TestList_Search_NotEQ(t *testing.T) {
	sl := New()

	// [24, 61, 67, 84, 91, 130, 133, 145, 150]
	seeds := []container.Int64{24, 61, 67, 84, 91, 130, 133, 145, 150}

	for _, k := range seeds {
		sl.Insert(k, int64(k*2+1))
	}

	var node *listNode

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchLastLT(container.Int64(21))
	require.Nil(t, node)

	node = sl.searchLastLT(container.Int64(24))
	require.Nil(t, node)

	node = sl.searchLastLT(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(24))

	node = sl.searchLastLT(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(67))

	node = sl.searchLastLT(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(130))

	node = sl.searchLastLT(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(130))

	node = sl.searchLastLT(container.Int64(146))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(145))

	node = sl.searchLastLT(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(145))

	node = sl.searchLastLT(container.Int64(156))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(150))

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchLastLE(container.Int64(21))
	require.Nil(t, node)

	node = sl.searchLastLE(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(24))

	node = sl.searchLastLE(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(67))

	node = sl.searchLastLE(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(130))

	node = sl.searchLastLE(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(133))

	node = sl.searchLastLE(container.Int64(137))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(133))

	node = sl.searchLastLE(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(150))

	node = sl.searchLastLE(container.Int64(156))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(150))

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchFirstGT(container.Int64(21))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(24))

	node = sl.searchFirstGT(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(61))

	node = sl.searchFirstGT(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(61))

	node = sl.searchFirstGT(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(84))

	node = sl.searchFirstGT(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(133))

	node = sl.searchFirstGT(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(145))

	node = sl.searchFirstGT(container.Int64(150))
	require.Nil(t, node)
	node = sl.searchFirstGT(container.Int64(151))
	require.Nil(t, node)

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchFirstGE(container.Int64(21))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(24))

	node = sl.searchFirstGE(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(24))

	node = sl.searchFirstGE(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(61))

	node = sl.searchFirstGE(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(84))

	node = sl.searchFirstGE(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(133))

	node = sl.searchFirstGE(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(133))

	node = sl.searchFirstGE(container.Int64(146))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(150))

	node = sl.searchFirstGE(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.key, container.Int64(150))

	node = sl.searchFirstGE(container.Int64(151))
	require.Nil(t, node)

}
