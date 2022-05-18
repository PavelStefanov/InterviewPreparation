package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

type level struct {
	node  *TreeNode
	level int
}

func maxDepth(root *TreeNode) int {
	maxLevel := 0

	stack := []level{}
	stack = append(stack, level{root, 1})

	for len(stack) > 0 {
		n := len(stack) - 1
		cur := stack[n]
		stack = stack[:n]

		if cur.node == nil {
			continue
		}

		if cur.node.Left == nil && cur.node.Right == nil {
			if cur.level > maxLevel {
				maxLevel = cur.level
			}
		}

		if cur.node.Left != nil {
			stack = append(stack, level{cur.node.Left, cur.level + 1})
		}
		if cur.node.Right != nil {
			stack = append(stack, level{cur.node.Right, cur.level + 1})
		}
	}

	return maxLevel
}
