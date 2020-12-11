package DataStructure

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) *ListNode {
	res := &ListNode{
		Val:  arr[len(arr)-1],
		Next: nil,
	}
	for i := len(arr) - 2; i >= 0; i-- {
		node := &ListNode{
			Val:  arr[i],
			Next: res,
		}
		res = node
	}
	return res
}
