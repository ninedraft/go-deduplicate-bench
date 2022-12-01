package deduplicate

import (
	"math/rand"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

var NoDuplicates []int
var ManyDuplicates []int

const n = 10_000

func init() {
	for i := 0; i < n; i++ {
		NoDuplicates = append(NoDuplicates, i)
		ManyDuplicates = append(ManyDuplicates, i%100)
	}

	rnd := rand.NewSource(100)
	rand.New(rnd).Shuffle(n, func(i, j int) {
		NoDuplicates[i], NoDuplicates[j] = NoDuplicates[j], NoDuplicates[i]
		ManyDuplicates[i], ManyDuplicates[j] = ManyDuplicates[j], ManyDuplicates[i]
	})
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
