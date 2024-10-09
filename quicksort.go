package main

import (
	"fmt"
)

func main() {
	arr := [10]int{5, 10, 8, 6, 7, 2, 9, 4, 1, 3}
	//arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := quicksort(arr[:])
	fmt.Println(res)
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[(len(arr) / 2)]
	left := []int{}
	middle := []int{}
	right := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}

	return append(append(quicksort(left), middle...), quicksort(right)...)

}
