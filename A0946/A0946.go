package main

import "fmt"

/*
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false。

示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

提示：

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed是popped的排列。
通过次数17,298提交次数28,696

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-stack-sequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	pushed := []int{
		0,
	}
	poped := []int{
		0,
	}
	fmt.Println(validateStackSequences(pushed, poped))
}

type stack struct {
	context []int
	ptr     int
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}
	pop := 0
	stack := &stack{
		context: make([]int, 0),
		ptr:     -1,
	}
	//stack.push(pushed[0])
	for i := 0; i < len(pushed); i++ {
		stack.push(pushed[i])
		if stack.peek() != popped[pop] {
			continue
		}
		for !stack.empty() && stack.peek() == popped[pop] {
			stack.pop()
			pop++
		}
	}
	if !stack.empty() {
		return false
	}
	return true
}

func (s *stack) empty() bool {
	return s.ptr < 0
}

func (s *stack) peek() int {
	return s.context[s.ptr]
}

func (s *stack) pop() (item int) {
	item = s.context[s.ptr]
	s.context = s.context[:s.ptr]
	s.ptr--
	return
}

func (s *stack) push(item int) {
	s.ptr++
	s.context = append(s.context, item)
}
