package main

import (
	"ACMLeetCode/DataStructure"
	"fmt"
)

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

通过次数387,800提交次数545,466

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := DataStructure.NewListNode(arr)
	reverseList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *DataStructure.ListNode) *DataStructure.ListNode {
	var prev *DataStructure.ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}
