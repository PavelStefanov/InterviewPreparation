package main

func main() {
}

func increasingBST(root *TreeNode) *TreeNode {
	var resRoot *TreeNode
	var res *TreeNode

	var stack []*TreeNode
	curr := root

	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	for len(stack) > 0 {
		n := len(stack) - 1

		curr = stack[n]
		stack = stack[:n]

		if res == nil {
			res = &TreeNode{Val: curr.Val}
			resRoot = res
		} else {
			res.Right = &TreeNode{Val: curr.Val}
			res = res.Right
		}

		curr = curr.Right
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
	}

	return resRoot
}
