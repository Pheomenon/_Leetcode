package main

import "fmt"

/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。

通过次数39,654提交次数67,574

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type stack struct {
	content []int
	ptr     int
}

func main() {
	nums := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 100}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	stack := &stack{
		content: make([]int, 1),
		ptr:     0,
	}
	stack.content[0] = 0
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	for i := 1; i < len(nums); i++ {
		for !stack.empty() && nums[stack.peek()] < nums[i] {
			ans[stack.peek()] = nums[i]
			stack.pop()
		}
		stack.push(i)
	}

	for i := 0; i < len(nums); i++ {
		for !stack.empty() && nums[stack.peek()] < nums[i] {
			if ans[stack.peek()] == -1 {
				ans[stack.peek()] = nums[i]
			}
			stack.pop()
		}
		stack.push(i)
	}
	return ans
}

func (s *stack) peek() int {
	return s.content[s.ptr]
}

func (s *stack) push(item int) {
	s.ptr++
	s.content = append(s.content, item)
}

func (s *stack) pop() int {
	ret := s.content[s.ptr]
	s.content = s.content[:s.ptr]
	s.ptr--
	return ret
}

func (s *stack) empty() bool {
	return s.ptr < 0
}
