package main

import "fmt"

func main() {
	fmt.Println(mySqrt(2))
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	low := 0
	high := x
	mid := (low + high) / 2
	for high-low > 1 {
		if mid*mid > x {
			high = mid
		} else {
			low = mid
		}
		mid = (high + low) / 2
	}
	return mid
}
