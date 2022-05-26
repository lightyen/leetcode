package main

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var traversal func(node *Node)
	traversal = func(node *Node) {
		for _, c := range node.Children {
			if c != nil {
				traversal(c)
			}
		}
		result = append(result, node.Val)
	}
	traversal(root)
	return result
}
