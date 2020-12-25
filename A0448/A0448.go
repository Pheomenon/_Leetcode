package main

/*
给定一个范围在 1 ≤ a[i] ≤ n (n = 数组大小 ) 的整型数组，数组中的元素一些出现了两次，另一些只出现一次。
找到所有在 [1, n] 范围之间没有出现在数组中的数字。
您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:
输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]
通过次数65,314提交次数107,130

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 9, 9, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))
}

/*
分析:
	根据题目特点，可以把数组中的元素与索引建立一一对应的关系。因为索引是确定的0到n-1,一个也不缺，而数组的元素不确定，少了哪个也不知道。
	既然两者是一一对应的关系，那么我们对数组中的每个元素对应的索引做个标记；
	然后再对索引进行一次遍历，那么不存的元素就不会对它对应的索引进行比较，由此可查找出这些不存在的元素。

思路:
	遍历每个元素，对索引进行标记
	将对应索引位置的值变为负数；
	遍历下索引，看看哪些索引位置上的数不是负数的。
	位置上不是负数的索引，对应的元素就是不存在的。
*/

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			v *= -1
		}
		if nums[v-1] < 0 {
			continue
		}
		nums[v-1] *= -1
	}
	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
