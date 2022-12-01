package deduplicate_test

import (
	"testing"

	"github.com/ninedraft/go-deduplicate-bench"
	"golang.org/x/exp/slices"
)

const N = 10_000

var _BenchmarkMapNoDuplicates []int

func BenchmarkMapNoDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		xx := deduplicate.GenerateWith(N, false)
		bench.StartTimer()
		_BenchmarkMapNoDuplicates = deduplicate.DeduplicateMap(xx)
	}
}

var _BenchmarkMapManyDuplicates []int

func BenchmarkMapManyDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		xx := deduplicate.GenerateWith(N, true)
		bench.StartTimer()

		_BenchmarkMapManyDuplicates = deduplicate.DeduplicateMap(xx)
	}
}

var _BenchmarkSortNoDuplicates []int

func BenchmarkSortNoDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		xx := deduplicate.GenerateWith(N, false)
		bench.StartTimer()

		_BenchmarkSortNoDuplicates = deduplicate.DeduplicateSort(xx)
	}
}

var _BenchmarkSortManyDuplicates []int

func BenchmarkSortManyDuplicates(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		xx := deduplicate.GenerateWith(N, true)
		bench.StartTimer()

		_BenchmarkSortManyDuplicates = deduplicate.DeduplicateSort(xx)
	}
}

func TestDeduplicateMap(t *testing.T) {
	xx := deduplicate.DeduplicateMap([]int{1, 1, 2, 2, 3, 3})
	if !slices.Equal(xx, []int{1, 2, 3}) {
		t.Errorf("[1, 2, 3] is expected, got %v", xx)
	}
}
