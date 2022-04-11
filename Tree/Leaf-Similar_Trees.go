package main

import "fmt"

func main() {
	r := leafSimilar(nil, nil)
	fmt.Print(r)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafes1, leafes2 := getLefes(root1), getLefes(root2)

	if len(leafes1) != len(leafes2) {
		return false
	}

	for i := 0; i < len(leafes1); i++ {
		if leafes1[i] != leafes2[i] {
			return false
		}
	}

	return true
}

func getLefes(root *TreeNode) []int {
	leafes := []int{}

	stack := []*TreeNode{}
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}

	for len(stack) > 0 {
		n := len(stack) - 1
		cur = stack[n]
		stack = stack[:n]

		if cur.Left == nil && cur.Right == nil {
			leafes = append(leafes, cur.Val)
		}

		cur = cur.Right
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
	}

	return leafes
}
