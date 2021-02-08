package main

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

示例 1：
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

示例 2：
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
通过次数128,988提交次数177,776

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	queue := Constructor()
	queue.DeleteHead()
	queue.AppendTail(5)
	queue.AppendTail(2)
	queue.DeleteHead()
	queue.DeleteHead()
}

type CQueue struct {
	SA *stack
	SB *stack
}

type stack struct {
	content []int
}

func (s *stack) pop() int {
	res := s.content[len(s.content)-1]
	s.content = s.content[:len(s.content)-1]
	return res
}

func (s *stack) push(val int) {
	s.content = append(s.content, val)
}

func (s *stack) empty() bool {
	return len(s.content) == 0
}

func Constructor() CQueue {
	return CQueue{
		SA: &stack{
			content: make([]int, 0),
		},
		SB: &stack{
			content: make([]int, 0),
		},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.SA.content = append(this.SA.content, value)
}

func (this *CQueue) DeleteHead() int {
	if this.SB.empty() {
		for !this.SA.empty() {
			this.SB.push(this.SA.pop())
		}
		if this.SB.empty() {
			return -1
		} else {
			return this.SB.pop()
		}
	} else {
		return this.SB.pop()
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
