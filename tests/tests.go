package tests

import (
	"math/rand"
	"time"

	"github.com/yu31/structs-go/avl"
	"github.com/yu31/structs-go/bs"
	"github.com/yu31/structs-go/container"
	"github.com/yu31/structs-go/rb"
	"github.com/yu31/structs-go/skip"
)

var containers = map[string]func() container.Container{
	"bstree": func() container.Container {
		return bs.New()
	},
	"avltree": func() container.Container {
		return avl.New()
	},
	"rbtree": func() container.Container {
		return rb.New()
	},
	"skiplist": func() container.Container {
		return skip.New()
	},
}

var trees = map[string]func() container.Tree{
	"bstree": func() container.Tree {
		return bs.New()
	},
	"avltree": func() container.Tree {
		return avl.New()
	},
	"rbtree": func() container.Tree {
		return rb.New()
	},
}

func shuffleSeeds(s1 []container.Int64) []container.Int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s2 := make([]container.Int64, len(s1))
	for i := 0; i < len(s1); i++ {
		s2[i] = s1[i]
	}
	for i := len(s2) - 1; i > 0; i-- {
		num := r.Intn(i + 1)
		s2[i], s2[num] = s2[num], s2[i]
	}
	return s2
}

func searchRange(ctr container.Container, start container.Key, boundary container.Key) []container.Element {
	var result []container.Element

	ctr.Range(start, boundary, func(ele container.Element) bool {
		result = append(result, ele)
		return true
	})
	return result
}

func searchRangeByIter(ctr container.Container, start container.Key, boundary container.Key) []container.Element {
	var result []container.Element

	it := ctr.Iter(start, boundary)
	for it.Valid() {
		n := it.Next()
		result = append(result, n)
	}
	return result
}

func searchReceive(ctr container.Container, start container.Key, boundary container.Key) []container.Element {
	var result []container.Element

	ctr.Reverse(start, boundary, func(ele container.Element) bool {
		result = append(result, ele)
		return true
	})
	return result
}

func reverseElementSlice(elements []container.Element) []container.Element {
	if len(elements) == 0 {
		return nil
	}
	result := make([]container.Element, len(elements))
	copy(result, elements)

	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	return result
}

func reverseSeedsSlice(seeds []container.Int64) []container.Int64 {
	if len(seeds) == 0 {
		return nil
	}
	result := make([]container.Int64, len(seeds))
	copy(result, seeds)

	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	return result
}
