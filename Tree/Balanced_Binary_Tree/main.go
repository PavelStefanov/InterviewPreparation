package main

import (
	"math"
)


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	flag := true
	checkBalanced(root, &flag)
	return flag
}

func checkBalanced(node *TreeNode, flag *bool) int {
	if node == nil {
		return 0
	}

	leftResult := checkBalanced(node.Left, flag)
	rightResult := checkBalanced(node.Right, flag)

	diff := math.Abs(float64(leftResult - rightResult))
	if diff > 1 {
		*flag = false
	}

	return max(leftResult, rightResult) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
