package deduplicate

import (
	"math/rand"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func GenerateWith(n int, duplicates bool) []int {
	var xx = make([]int, n)
	for i := 0; i < n; i++ {
		x := i
		if duplicates {
			x %= 100
		}
		xx[i] = x
	}

	rnd := rand.NewSource(100)
	rand.New(rnd).Shuffle(n, func(i, j int) {
		xx[i], xx[j] = xx[j], xx[i]
	})
	return xx
}

func DeduplicateMap[E comparable](slice []E) []E {
	set := make(map[E]struct{})
	var i int
	for _, x := range slice {
		_, ok := set[x]
		if ok {
			continue
		}
		set[x] = struct{}{}
		slice[i] = x
		i++
	}
	return slice[:i]
}

func DeduplicateSort[E constraints.Ordered](slice []E) []E {
	slices.Sort(slice)
	return slices.Compact(slice)
}
