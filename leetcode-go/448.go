package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	// swap to order the array of numbers
	for i := 0; i < len(nums); {
		if nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			highestNum := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = highestNum
		} else {
			i += 1
		}
	}
	// if num isn't in sequence, add it to response array
	response := make([]int, 0)
	for i, _ := range nums {
		if nums[i] != i+1 {
			response = append(response, i+1)
		}
	}
	return response
}
