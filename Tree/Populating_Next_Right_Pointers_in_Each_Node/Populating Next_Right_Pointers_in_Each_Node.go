package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	connectTwoNode(root.Left, root.Right)

	return root
}

func connectTwoNode(left *Node, right *Node) {
	if left == nil {
		return
	}
	left.Next = right
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(left.Right, right.Left)
	connectTwoNode(right.Left, right.Right)
}

type tmp struct {
	node  *Node
	level int
}

func connectIterativly(root *Node) *Node {

	queue := []*tmp{}
	queue = append(queue, &tmp{node: root, level: 0})

	var prev *tmp
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node == nil {
			continue
		}

		if prev != nil && prev.node != nil && prev.level == cur.level {
			prev.node.Next = cur.node
		}
		prev = cur

		queue = append(queue, &tmp{node: cur.node.Left, level: cur.level + 1})
		queue = append(queue, &tmp{node: cur.node.Right, level: cur.level + 1})
	}

	return root
}
