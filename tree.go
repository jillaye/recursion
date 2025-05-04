package main

import "fmt"

// Tree Nodes
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// AddTreeNode uses recursion to insert a new node with data into
// the binary search tree whose root = root.
// If data is already in the tree, it returns the original tree.
func AddTreeNode(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{data: data}
	}
	// traverse tree
	if root.data < data {
		root.right = AddTreeNode(root.right, data)
	} else if root.data > data {
		root.left = AddTreeNode(root.left, data)
	}
	return root
}

func SumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Printf("Node: %d\n", root.data)
	return root.data + SumTree(root.left) + SumTree(root.right)
}

func PrintLeafNodes(root *TreeNode) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		fmt.Printf("\nLeaf Node: %d", root.data)
	}
	PrintLeafNodes(root.left)
	PrintLeafNodes(root.right)
}
