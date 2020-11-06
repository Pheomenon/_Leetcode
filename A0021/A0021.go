/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/

通过次数69,736提交次数95,100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ptr1, ptr2 := min(l1, l2)
	ptr := ptr1
	result := ptr
	ptr1 = ptr1.Next
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			ptr.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			ptr.Next = ptr2
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
	}
	if ptr1 == nil {
		ptr.Next = ptr2
		return result
	}
	if ptr2 == nil {
		ptr.Next = ptr1
		return result
	}
	return result
}

func min(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
	if l1.Val < l2.Val {
		return l1, l2
	}
	return l2, l1
}
