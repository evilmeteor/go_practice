package main

import (
	"fmt"
)

func main() {
	arr := [10]int{5, 10, 8, 6, 7, 2, 9, 4, 1, 3}
	res := quicksort(arr[:])
	fmt.Println(res)
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	middle := []int{}
	right := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}
	middle = append(middle, pivot)

	fmt.Println(pivot, left, middle, right)

	return append(append(quicksort(left), middle...), quicksort(right)...)

}
