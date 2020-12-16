package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yu31/gostructs/container"
)

func BenchmarkContainer_Insert(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Intn(b.N*100)), nil)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		b.Run(name, func(b *testing.B) {
			process(b, f())
		})
	}
}

func BenchmarkContainer_Search(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Intn(b.N*100)), nil)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Search(container.Int64(r.Intn(b.N * 100)))
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		b.Run(name, func(b *testing.B) {
			process(b, f())
		})
	}
}

func BenchmarkContainer_Delete(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Intn(b.N*100)), nil)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Delete(container.Int64(r.Intn(b.N * 100)))
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		b.Run(name, func(b *testing.B) {
			process(b, f())
		})
	}
}

func BenchmarkContainer_Update(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		for i := 0; i < b.N; i++ {
			box.Insert(container.Int64(r.Intn(b.N*100)), nil)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			box.Update(container.Int64(r.Intn(b.N*100)), nil)
		}
	}

	// Test for all container implementation.
	for name, f := range containers {
		b.Run(name, func(b *testing.B) {
			process(b, f())
		})
	}
}

func BenchmarkContainer_Replace(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	process := func(b *testing.B, box container.Container) {
		b.Run("same-insert", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				box.Replace(container.Int64(r.Intn(b.N*100)), nil)
			}
		})
		b.Run("same-update", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				box.Replace(container.Int64(r.Intn(b.N*100)), nil)
			}
		})
	}

	// Test for all container implementation.
	for name, f := range containers {
		b.Run(name, func(b *testing.B) {
			process(b, f())
		})
	}
}
