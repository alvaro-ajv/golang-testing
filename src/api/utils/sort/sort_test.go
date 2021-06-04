package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println(elements)

	// Execution
	BubbleSort(elements)

	//Validation
	if elements[0] != 9 {
		t.Error("First element should be 9")
	}

	if elements[len(elements)-1] != 0{
		t.Error("Last element should be 0")
	}
	fmt.Println(elements)
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Init
	elements := []int{7,6,5,4,3,2,1,0}

	// Execution
	BubbleSort(elements)

}
