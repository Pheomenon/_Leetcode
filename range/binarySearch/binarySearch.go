package main

import "fmt"

func main() {
	fmt.Println(searchI([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
}

func searchI(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchII(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums) // attention!
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}
