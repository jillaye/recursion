package main

import "fmt"

// Linked List Nodes
type Node struct {
	data int
	next *Node
}

func add_node(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		return newNode
	}
	// use a ptr to traverse the list to the end
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return head
}

func print_list(head *Node) {
	fmt.Println("Linked list:")
	current := head
	// use a ptr to traverse the list to the end
	for current.next != nil {
		fmt.Printf("%d => ", current.data)
		current = current.next
	}
	// print the last node
	fmt.Printf("%d\n", current.data)
}

func reverse(head *Node) {
	if head.next == nil {
		return
	}
	(head.next).next = head
	reverse(head.next)
}
