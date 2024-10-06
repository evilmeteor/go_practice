package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	item := 3

	index, err := binary_search(arr[:], 0, len(arr)-1, item)

	if err != nil {
		fmt.Printf("Error: Item %d not found.\n", item)
	} else {
		fmt.Printf("Item %d found at index %d.\n", item, index)
	}
}

func binary_search(list []int, low int, high int, item int) (int, error) {
	if high >= low {
		mid := (low + high) / 2
		if list[mid] == item {
			return mid, nil
		}

		if list[mid] < item {
			return binary_search(list, mid+1, high, item)
		}

		if list[mid] > item {
			return binary_search(list, low, mid-1, item)
		}
	}

	return -1, errors.New("Not Found")
}
