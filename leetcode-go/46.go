package main

import "fmt"

// My implementation of Heap's algorithm
func permutations(arr []int) [][]int {
	var findPermutations func([]int, int)
	res := [][]int{}

	findPermutations = func(arr []int, n int) {
		if n == 1 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			res = append(res, temp)
		} else {
			for i := 0; i < n; i++ {
				findPermutations(arr, n-1)
				if n%2 == 1 {
					temp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = temp
				} else {
					temp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = temp
				}
			}
		}
	}
	findPermutations(arr, len(arr))
	return res
}
