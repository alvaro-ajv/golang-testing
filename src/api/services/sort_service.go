package services

import "github.com/alvaro259818/golang-testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) <= 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
