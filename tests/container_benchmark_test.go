package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yu31/gostructs/avl"
	"github.com/yu31/gostructs/bs"
	"github.com/yu31/gostructs/container"
	"github.com/yu31/gostructs/rb"
	"github.com/yu31/gostructs/skip"
)

func BenchmarkContainer_Insert(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Int63()), nil)
		}
	}

	b.Run("bstree", func(b *testing.B) {
		process(b, bs.New())
	})
	b.Run("avtree", func(b *testing.B) {
		process(b, avl.New())
	})
	b.Run("rbtree", func(b *testing.B) {
		process(b, rb.New())
	})
	b.Run("skiplist", func(b *testing.B) {
		process(b, skip.New())
	})
}

func BenchmarkContainer_Search(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Int63()), nil)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Search(container.Int64(r.Int63()))
		}
	}

	b.Run("bstree", func(b *testing.B) {
		process(b, bs.New())
	})
	b.Run("avtree", func(b *testing.B) {
		process(b, avl.New())
	})
	b.Run("rbtree", func(b *testing.B) {
		process(b, rb.New())
	})
	b.Run("skiplist", func(b *testing.B) {
		process(b, skip.New())
	})
}

func BenchmarkContainer_Delete(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Int63()), nil)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Delete(container.Int64(r.Int63()))
		}
	}

	b.Run("bstree", func(b *testing.B) {
		process(b, bs.New())
	})
	b.Run("avtree", func(b *testing.B) {
		process(b, avl.New())
	})
	b.Run("rbtree", func(b *testing.B) {
		process(b, rb.New())
	})
	b.Run("skiplist", func(b *testing.B) {
		process(b, skip.New())
	})
}
