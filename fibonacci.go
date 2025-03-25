package main

import "fmt"

// Function  fibonacci returns the fibonacci series up to n
func fibonacci(n int) []int {
	fmt.Printf("n = %d\n", n)
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	last := fibonacci(n - 1)
	lastSum := last[n-1] + last[n-2]
	result := append(last, lastSum)
	return result
}

// nthFibo returns the nth term of the fibonacci series
func nthFibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (nthFibo(n-1) + nthFibo(n-2))
}
