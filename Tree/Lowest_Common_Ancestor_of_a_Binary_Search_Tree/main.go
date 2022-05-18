package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var min, max *TreeNode

	if p.Val < q.Val {
		min = p
		max = q
	} else {
		min = q
		max = p
	}
	return findNode(root, min, max)
}

func findNode(root, p, q *TreeNode) *TreeNode {
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}

	if p.Val <= root.Val && q.Val <= root.Val {
		return findNode(root.Left, p, q)
	} else {
		return findNode(root.Right, p, q)
	}
}
