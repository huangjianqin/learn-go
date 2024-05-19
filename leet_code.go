package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

// Create a map to store the values and their indices
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// Iterate through the array
	for i, num := range nums {
		// Check if the value is already in the map
		j, exist := m[num]
		// If it is, return the indices
		if exist {
			return []int{j, i}
		}
		// Otherwise, add the value and its index to the map
		m[target-num] = i
	}

	// If no pair is found, return an empty slice
	return []int{}
}
