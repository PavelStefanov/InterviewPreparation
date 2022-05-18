package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	in  *TreeNode
	out *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var out *TreeNode

	queue := []*pair{{root, out}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		cur.out.Val = cur.in.Val

		if cur.in.Left != nil {
			cur.out.Right = &TreeNode{}
			queue = append(queue, &pair{cur.in.Left, cur.out.Right})
		}

		if cur.in.Right != nil {
			cur.out.Left = &TreeNode{}
			queue = append(queue, &pair{cur.in.Right, cur.out.Left})
		}
	}

	return out
}

func invertTreeRecursivly(root *TreeNode) *TreeNode {
	return invertTreeNested(root)
}

func invertTreeNested(in *TreeNode) *TreeNode {
	if in == nil {
		return nil
	}

	return &TreeNode{
		in.Val,
		invertTreeNested(in.Right),
		invertTreeNested(in.Left),
	}
}
