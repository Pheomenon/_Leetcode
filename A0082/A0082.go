package main

import "math"

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中没有重复出现的数字。

示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5

示例 2:
输入: 1->1->1->2->3
输出: 2->3
通过次数81,951提交次数163,865

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	var dummy *ListNode
	dummy = &ListNode{
		Val:  math.MinInt64,
		Next: nil,
	}
	var flag bool
	ans := dummy
	for curr != nil {
		next := curr.Next
		for next != nil && curr.Val == next.Val {
			next = next.Next
			flag = true
		}
		if flag {
			curr = next
			flag = false
		} else {
			dummy.Next = curr
			dummy = curr
			curr = curr.Next
		}
	}
	dummy.Next = curr
	return ans.Next
}
