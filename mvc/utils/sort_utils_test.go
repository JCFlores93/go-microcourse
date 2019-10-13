package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	els = BubbleSort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 9, len(els))

	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
	assert.EqualValues(t, 5, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	els = BubbleSort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 9, len(els))

	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
	assert.EqualValues(t, 5, els[4])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])

}
func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func BenchmarkBubbleSort(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSortNative(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort1000000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSortNative1000000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
