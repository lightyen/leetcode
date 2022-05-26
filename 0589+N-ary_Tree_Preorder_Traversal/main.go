package main

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var traversal func(node *Node)
	traversal = func(node *Node) {
		result = append(result, node.Val)
		for _, c := range node.Children {
			if c != nil {
				traversal(c)
			}
		}
	}
	traversal(root)
	return result
}
