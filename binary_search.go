package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	item := 3

	index, err := binary_search(arr[:], item)

	if err != nil {
		fmt.Printf("Error: Item %d not found.\n", item)
	} else {
		fmt.Printf("Item %d found at index %d.\n", item, index)
	}
}

func binary_search(list []int, item int) (int, error) {
	low := 0
	high := len(list) - 1
	count := 0

	for low <= high {
		count += 1
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid, nil
		}

		if guess < item {
			low = mid + 1
		}

		if guess > item {
			high = mid - 1
		}

	}

	return -1, errors.New("Not Found")
}
