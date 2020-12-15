package tests

import (
	"testing"

	"github.com/yu31/gostructs/avl"
	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/rb"
	"github.com/yu31/gostructs/skip"
)

func TestContainer_Interface(t *testing.T) {
	// Ensure the bs/avl/rb/skip are implements the container.Container.
	t.Run("container", func(t *testing.T) {
		var box container.Container
		_ = box

		box = bs.New()
		box = avl.New()
		box = rb.New()
		box = skip.New()
	})

	t.Run("iterator", func(t *testing.T) {
		var it container.Iterator
		_ = it

		it = bs.NewIterator(nil, nil, nil)
	})
}
