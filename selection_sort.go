package main

import "fmt"

func main() {
	arr := [10]int{5, 10, 8, 6, 7, 2, 9, 4, 1, 3}
	res := selection_sort(arr[:])
	fmt.Println(res)
}

func find_smallest(slice []int) int {
	smallest := slice[0]
	smallest_index := 0

	for i, v := range slice {
		if v < smallest {
			smallest = v
			smallest_index = i
		}
	}
	return smallest_index
}

func selection_sort(slice []int) []int {
	new_slice := []int{}
	slice_len := len(slice)

	for i := 0; i < slice_len; i++ {
		index := find_smallest(slice)
		new_slice = append(new_slice, slice[index])
		slice = append(slice[:index], slice[index+1:]...)
	}
	return new_slice
}
