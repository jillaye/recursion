package main

import "fmt"

// bibSearch performs NON-RECURSIVE binary search
func binSearch(nums []int, target int) int {
	var left int = 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + (right-left)/2)
		fmt.Printf("left = %d, right = %d, middle = %d\n", left, right, middle)
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

// rbinSearch performs RECURSIVE binary search
func rbinSearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	var mid int = left + (right-left)/2
	var midval int = nums[mid]
	if midval == target {
		return mid
	}
	if midval > target {
		return rbinSearch(nums, left, (mid - 1), target)
	}
	return rbinSearch(nums, (mid + 1), right, target)
}
