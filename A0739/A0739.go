package main

/*
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

通过次数118,471提交次数181,035

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	temperatures := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	dailyTemperatures(temperatures)
}

type stack struct {
	content []int
	ptr     int
}

func dailyTemperatures(T []int) []int {
	stack := &stack{
		content: make([]int, 1),
		ptr:     0,
	}
	stack.content[0] = 0
	ans := make([]int, len(T))

	for i := 1; i < len(T); i++ {
		for !stack.empty() && T[i] > T[stack.peek()] {
			ans[stack.peek()] = i - stack.peek()
			stack.pop()
		}
		stack.push(i)
	}
	return ans
}

func (s *stack) peek() int {
	return s.content[s.ptr]
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

func (s *stack) push(item int) {
	s.ptr++
	s.content = append(s.content, item)
}
