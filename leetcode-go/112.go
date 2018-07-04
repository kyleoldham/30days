package main

import "fmt"

func hasPathSum(root *TreeNode, sum int) bool {
	// Base Case
	if root == nil {
		return false
	}

	// decrease sum until it equals 0 or misses
	sum -= root.Val

	// Values
	if root.Left == nil && root.Right == nil {
		// Lol Syntax. ==
		return sum == 0
		//return true
	}

	// Send back values similar to 104.go
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
