package main

import (
	"fmt"
	"sync"
)

/*
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/binary-search/xeog5j/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

var wg sync.WaitGroup

func search(nums []int, target int) int {
	partition0, partition1 := -1, -1
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pivot = i
			break
		}
	}
	wg.Add(2)
	go binarySearch(nums[0:pivot], target, &partition0)
	go binarySearch(nums[pivot:], target, &partition1)
	wg.Wait()
	if partition1 == -1 {
		return partition0
	}
	return partition1 + pivot
}

//func search(nums []int, target int) int {
//	pivot := 0
//	if target == nums[0] {
//		return 0
//	} else {
//		for i := 1; i < len(nums); i++ {
//			if nums[i] == target {
//				return i
//			}
//			if nums[i-1] > nums[i] {
//				pivot = i
//				break
//			}
//		}
//	}
//	tmp := binarySearch(nums[pivot:], target)
//	if tmp == -1 {
//		return tmp
//	} else {
//		return tmp + pivot
//	}
//}

func binarySearch(nums []int, target int, partition *int) {
	defer wg.Done()
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			*partition = mid
			return
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}

	*partition = -1
}
