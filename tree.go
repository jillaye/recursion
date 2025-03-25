package main

// Tree Nodes
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func sumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.data + sumTree(root.left) + sumTree(root.right)
}
