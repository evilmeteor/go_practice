package main

import (
	"fmt"
)

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res1 := sum(arr[:])
	fmt.Printf("sum: %d\n", res1)
	res2 := count(arr[:])
	fmt.Printf("count: %d\n", res2)
	res3 := max(arr[:])
	fmt.Printf("max: %d\n", res3)
}

func sum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	return slice[0] + sum(slice[1:len(slice)])
}

func count(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	return 1 + count(slice[1:len(slice)])
}

func max(slice []int) int {
	if len(slice) == 2 {
		if slice[0] > slice[1] {
			return slice[0]
		} else {
			return slice[1]
		}
	}

	sub_max := max(slice[1:len(slice)])

	if slice[0] > sub_max {
		return slice[0]
	}

	return sub_max
}
