package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
通过次数97,284提交次数187,038

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	head := linked(nums, 0)
	m, n := 2, 4
	reverseBetween(head, m, n)
}

func linked(nums []int, idx int) *ListNode {
	if idx > len(nums)-1 {
		return nil
	}
	return &ListNode{
		Val:  nums[idx],
		Next: linked(nums, idx+1),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	slow := head
	var slowPre *ListNode
	for i := 1; i < m; i++ {
		slowPre = slow
		slow = slow.Next
	}

	fast := slow
	for i := 0; i < n-m; i++ {
		fast = fast.Next
	}
	if slowPre != nil {
		slowPre.Next = reverse(slow, fast)
	} else {
		head = reverse(slow, fast)
	}
	return head
}

func reverse(head, tail *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next = curr
	return prev
}
