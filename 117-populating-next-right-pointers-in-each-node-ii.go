package main

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	type Pair struct {
		node  *Node
		level int
	}
	stack := []Pair{{node: root, level: 1}}
	var prev *Node
	for curr := 0; len(stack) > 0; {
		pair := stack[0]
		stack = stack[1:]
		if pair.level != curr {
			curr = pair.level
		} else {
			prev.Next = pair.node
		}
		prev = pair.node
		if pair.node.Left != nil {
			stack = append(stack, Pair{node: pair.node.Left, level: pair.level + 1})
		}
		if pair.node.Right != nil {
			stack = append(stack, Pair{node: pair.node.Right, level: pair.level + 1})
		}
	}
	return root
}
