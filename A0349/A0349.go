package main

import (
	"fmt"
)

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]

说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
通过次数155,672提交次数211,749

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	num1 := []int{
		61, 24, 20, 58, 95, 53, 17, 32, 45, 85, 70, 20, 83, 62, 35, 89, 5, 95, 12, 86, 58, 77, 30, 64, 46, 13, 5, 92, 67, 40, 20, 38, 31, 18, 89, 85, 7, 30, 67, 34, 62, 35, 47, 98, 3, 41, 53, 26, 66, 40, 54, 44, 57, 46, 70, 60, 4, 63, 82, 42, 65, 59, 17, 98, 29, 72, 1, 96, 82, 66, 98, 6, 92, 31, 43, 81, 88, 60, 10, 55, 66, 82, 0, 79, 11, 81,
	}
	num2 := []int{
		5, 25, 4, 39, 57, 49, 93, 79, 7, 8, 49, 89, 2, 7, 73, 88, 45, 15, 34, 92, 84, 38, 85, 34, 16, 6, 99, 0, 2, 36, 68, 52, 73, 50, 77, 44, 61, 48,
	}

	ar := intersection(num1, num2)
	for _, i2 := range ar {
		fmt.Printf("%d\n", i2)
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	var mark1, mark2 uint64
	for _, v := range nums1 {
		if v > 63 {
			return sd(nums1, nums2)
		}
		tmp := uint64(1)
		tmp <<= v
		mark1 |= tmp
	}

	for _, v := range nums2 {
		if v > 63 {
			return sd(nums1, nums2)
		}
		tmp := uint64(1)
		tmp <<= v
		mark2 |= tmp
	}

	term := mark1 & mark2

	ans := []int{}
	length := 0

	for i := 0; ; i++ {
		if term>>i == 0 {
			break
		} else {
			length++
		}
	}

	for i := 0; i < length; i++ {
		tmp := uint64(1)
		tmp <<= i
		var exist uint64
		exist = tmp & term

		if exist != 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func sd(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}
