package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	dummy := &TreeNode{}

	merge(root1, root2, dummy, 0)

	return dummy.Left
}

func merge(root1 *TreeNode, root2 *TreeNode, result *TreeNode, r int) {
	if root1 == nil && root2 == nil {
		return
	}

	if root1 != nil && root2 != nil {
		cur := &TreeNode{Val: root1.Val + root2.Val}
		if r == 0 {
			result.Left = cur
		} else {
			result.Right = cur
		}

		merge(root1.Left, root2.Left, cur, 0)
		merge(root1.Right, root2.Right, cur, 1)
	} else if root1 != nil {
		if r == 0 {
			result.Left = root1
		} else {
			result.Right = root1
		}
	} else {
		if r == 0 {
			result.Left = root2
		} else {
			result.Right = root2
		}
	}
	return
}
