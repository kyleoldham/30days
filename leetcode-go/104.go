package main

import "fmt"

func maxDepth(root *TreeNode) int {
	// Base Case
	if root == nil {
		return 0
	}
	// Max of left and right + 1 to set root further
	return (maxCalc(maxDepth(root.Left), maxDepth(root.Right))) + 1
}

// get max
func maxCalc(x, y int) int {
	if x > y {
		return x
	}
	// this is all you need for go; no else
	return y
}
