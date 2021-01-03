/**
给出两个非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
通过次数619,636提交次数1,593,923

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	head1 := &ListNode{
		Val:  l1[0],
		Next: nil,
	}
	curr1 := head1
	for i := 1; i < len(l1); i++ {
		newNode := &ListNode{}
		newNode.Val = l1[i]
		curr1.Next = newNode
		curr1 = newNode
	}

	l2 := []int{9, 9, 9, 9}
	head2 := &ListNode{
		Val:  l2[0],
		Next: nil,
	}
	curr2 := head2
	for i := 1; i < len(l2); i++ {
		newNode := &ListNode{}
		newNode.Val = l2[i]
		curr2.Next = newNode
		curr2 = newNode
	}

	fmt.Println(addTwoNumbers(head1, head2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	//head := &ListNode{
	//	Val:  l1.Val + l2.Val,
	//	Next: nil,
	//}
	curr := head
	currL1 := l1
	currL2 := l2

	ret := 0
	for currL1 != nil && currL2 != nil {
		value := currL1.Val + currL2.Val + ret
		ret = 0
		if value >= 10 {
			value -= 10
			ret = 1
		}
		newNode := &ListNode{
			Val:  value,
			Next: nil,
		}
		curr.Next = newNode
		curr = newNode

		currL1 = currL1.Next
		currL2 = currL2.Next
	}

	if currL1 == nil {
		curr.Next = currL2
	}
	if currL2 == nil {
		curr.Next = currL1
	}

	if ret != 0 {
		node := &ListNode{
			Val:  1,
			Next: nil,
		}
		curr.Next = addTwoNumbers(node, curr.Next)
	}

	return head.Next
}
