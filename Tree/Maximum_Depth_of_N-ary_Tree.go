package main

func maxDepthIteranivly(root *Node) int {
	if root == nil {
		return 0
	}

	max := 0

	m := make(map[*Node]int)
	m[root] = 1

	var stack []*Node

	stack = append(stack, root)

	for len(stack) > 0 {
		n := len(stack) - 1
		cur := stack[n]
		stack = stack[:n]

		v, _ := m[cur]

		if len(cur.Children) == 0 && max < v+1 {
			max = v
			continue
		}

		for i := 0; i < len(cur.Children); i++ {
			m[cur.Children[i]] = v + 1
			stack = append(stack, cur.Children[i])
		}
	}

	return max
}

func maxDepthRecursivly(root *Node) int {
	if root == nil {
		return 0
	}

	return find(root, 1)
}

func find(root *Node, level int) int {
	if len(root.Children) == 0 {
		return level
	}

	m := 0
	for i := 0; i < len(root.Children); i++ {
		l := find(root.Children[i], level+1)

		if l > m {
			m = l
		}
	}

	return m
}
