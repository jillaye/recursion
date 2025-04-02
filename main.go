package main

import (
	"fmt"
)

var six = TreeNode{data: 6, left: nil, right: nil}
var five = TreeNode{data: 5, left: nil, right: nil}
var four = TreeNode{data: 4, left: nil, right: nil}
var three = TreeNode{data: 3, left: &five, right: &six}
var root = TreeNode{data: 2, left: &three, right: &four}

/************* recursive functions ***************************/

func main() {
	// Create a slice of bytes
	data := []byte{'H', 'e', 'l', 'l', 'o'}
	fmt.Println(data)

	// Convert a slice of bytes back to a string
	newStr := string(data)
	fmt.Println(newStr)

	fmt.Println("Non-recursive function:")
	numbers := []int{1, 7, 13, 21, 27, 30, 37}
	var left int = 0
	var right int = len(numbers) - 1
	idx := binSearch(numbers, 2)
	fmt.Printf("idx = %d\n\n", idx)
	fmt.Println("Recursive function:")
	rIdx := rbinSearch(numbers, left, right, 2)
	fmt.Printf("rIdx = %d\n\n", rIdx)

	var decimal uint32 = 10
	fmt.Printf("Decimal: %d = Binary: %s\n", decimal, rConvert(decimal))
	fmt.Printf("Decimal: %d = Sum: %d\n\n", decimal, rSumAll(decimal))

	fmt.Printf("Sum of binary tree = %d\n\n", sumTree(&root))

	n := 10
	fmt.Printf("fibo(%d) = %v\n", n, fibonacci(n))
	fmt.Printf("nthFibo(%d) = %v\n", n, nthFibo(n))

	// merge sort ************************************************
	merge_result := merge([]int{0, 1, 2, 3, 4}, []int{-1, 7, 9, 10, 20})
	fmt.Printf("merge result: %v\n", merge_result)
	msr := mergeSort([]int{4, 1, 3, 2, 0, -1, 7, 10, 9, 20})
	fmt.Printf("merge sort result: %v\n", msr)
	msr2 := mergeSort([]int{7, 4, 2, -3, 6, -11, 20, 3})
	fmt.Printf("merge sort result 2: %v\n", msr2)

	// linked list ***********************************************
	even_head := Node{data: 10}
	add_node(&even_head, 12)
	add_node(&even_head, 14)
	add_node(&even_head, 16)
	add_node(&even_head, 18)
	print_list(&even_head)
	odd_head := Node{data: 11}
	add_node(&odd_head, 13)
	add_node(&odd_head, 15)
	add_node(&odd_head, 17)
	add_node(&odd_head, 19)
	print_list(&odd_head)
	// fmt.Println("Non recursive reverse (even):")
	// r_list := reverse(&even_head)
	// print_list(r_list)
	// fmt.Println("Recursive reverse (odd):")
	// rr_list := r_reverse(&odd_head)
	// print_list(rr_list)
	fmt.Println("Recursive merge:")
	m_list := r_merge_lists(&even_head, &odd_head)
	print_list(m_list)
}
