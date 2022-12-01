package deduplicate_test

import (
	"testing"

	"github.com/ninedraft/go-deduplicate-bench"
	"golang.org/x/exp/slices"
)

var _BenchmarkMapNoDuplicates []int

func BenchmarkMapNoDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		_BenchmarkMapNoDuplicates = deduplicate.DeduplicateMap(deduplicate.NoDuplicates)
	}
}

var _BenchmarkMapManyDuplicates []int

func BenchmarkMapManyDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		_BenchmarkMapManyDuplicates = deduplicate.DeduplicateMap(deduplicate.ManyDuplicates)
	}
}

var _BenchmarkSortNoDuplicates []int

func BenchmarkSortNoDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		_BenchmarkSortNoDuplicates = deduplicate.DeduplicateSort(deduplicate.NoDuplicates)
	}
}

var _BenchmarkSortManyDuplicates []int

func BenchmarkSortManyDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		_BenchmarkSortManyDuplicates = deduplicate.DeduplicateSort(deduplicate.ManyDuplicates)
	}
}

func TestDeduplicateMap(t *testing.T) {
	xx := deduplicate.DeduplicateMap([]int{1, 1, 2, 2, 3, 3})
	if !slices.Equal(xx, []int{1, 2, 3}) {
		t.Errorf("[1, 2, 3] is expected, got %v", xx)
	}
}
