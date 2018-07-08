package main

import "fmt"

func removeDuplicates(nums []int) int {
	current, next, size := 0, 1, len(nums)
	for ; next < size; next++ {
		if nums[current] == nums[next] {
			continue
		}
		current++
		nums[current], nums[next] = nums[next], nums[current]
	}
	return current + 1
}
