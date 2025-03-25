package main

// merges 2 sorted slices into one (sorted) slice
func merge(a []int, b []int) []int {
	var result = make([]int, 0, (len(a) + len(b)))
	a_idx := 0
	b_idx := 0
	for len(result) <= cap(result) {
		if a_idx >= len(a) || b_idx >= len(b) {
			break
		}
		if a[a_idx] <= b[b_idx] {
			result = append(result, a[a_idx])
			a_idx++
		} else {
			result = append(result, b[b_idx])
			b_idx++
		}
	}
	// append result with remaining elements of a or b
	if a_idx >= len(a) {
		result = append(result, b[b_idx:]...)
	} else {
		result = append(result, a[a_idx:]...)
	}

	return result
}

func mergeSort(data []int) []int {
	// base case
	if len(data) <= 1 {
		return data
	}
	idx := len(data) / 2
	return merge(mergeSort(data[:idx]), mergeSort(data[idx:]))
}
