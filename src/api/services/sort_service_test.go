package services

import (
	"github.com/alvaro259818/golang-testing/src/api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	//Init
	elements := sort.GetElements(10001)

	// Execution
	Sort(elements)

	//Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 10000, elements[len(elements)-1])

}

func TestSortMoreThan10000(t *testing.T) {
	//Init
	elements := sort.GetElements(10001)
	// Execution
	Sort(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("Last element should be 10000")
	}
}

func BenchmarkSort10K(b *testing.B) {
	elements := sort.GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100K(b *testing.B) {
	elements := sort.GetElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
