package main

import "fmt"

// converts decimal to binary
func rConvert(dec uint32) string {
	if dec == 1 || dec == 0 {
		return fmt.Sprintf("%d", dec)
	}
	return rConvert(dec/2) + fmt.Sprintf("%d", dec%2)
}

// sums all numbers up to and including num
func rSumAll(num uint32) uint32 {
	if num <= 0 {
		return 0
	}
	return num + rSumAll(num-1)
}
