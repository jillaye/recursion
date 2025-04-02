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

// Non-recursive reverse
func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var prev *Node
	current := head
	// use a ptr to traverse the list to the end
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
	return prev
}

// Recursive reverse
func r_reverse(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}
	head := r_reverse(node.next)
	head.next = node
	return head
}

func merge_lists(a *Node, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var head *Node
	var tail *Node
	for a != nil && b != nil {
		if a.data <= b.data {
			if head == nil {
				head = a
				tail = head
			} else {
				tail.next = a
				tail = tail.next
			}
			a = a.next
		} else {
			if head == nil {
				head = b
				tail = head
			} else {
				tail.next = b
				tail = tail.next
			}
			b = b.next
		}
	}
	if a != nil {
		tail.next = a
	} else if b != nil {
		tail.next = b
	}
	return head
}
