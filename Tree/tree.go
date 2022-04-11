package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

func ArrayToTree(nums []int) *TreeNode {
	var resRoot *TreeNode

	queue := []*TreeNode{}

	for i := 0; i < len(nums); i++ {
		c := &TreeNode{Val: nums[i]}

		if len(queue) == 0 {
			resRoot = c
			queue = append(queue, c)
		} else {
			el := queue[0]

			if el.Left == nil {
				el.Left = c
			} else {
				el.Right = c
				queue = queue[1:]
			}
		}

	}

	return resRoot
}
