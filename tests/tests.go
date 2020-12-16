package tests

import (
	"math/rand"
	"time"

	"github.com/yu31/gostructs/container"
)

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

func searchRange(box container.Container, start container.Key, boundary container.Key) []container.Element {
	var result []container.Element

	box.Range(start, boundary, func(ele container.Element) bool {
		result = append(result, ele)
		return true
	})
	return result
}

func searchRangeByIter(box container.Container, start container.Key, boundary container.Key) []container.Element {
	var result []container.Element

	it := box.Iter(start, boundary)
	for it.Valid() {
		n := it.Next()
		result = append(result, n)
	}
	return result
}
