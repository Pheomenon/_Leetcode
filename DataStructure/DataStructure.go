package DataStructure

type LinkedList struct {
	root Node
	len  int
}

type Node struct {
	next, prev *Node
	ListNode   *LinkedList
	Value      interface{}
}

func (l *LinkedList) Init() *LinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *LinkedList {
	return new(LinkedList).Init()
}

func (l *LinkedList) Len() int {
	return l.len
}

//func NewListNode(arr []int) *LinkedList {
//	res := &LinkedList{
//		Val:  arr[len(arr)-1],
//		Next: nil,
//	}
//	for i := len(arr) - 2; i >= 0; i-- {
//		node := &LinkedList{
//			Val:  arr[i],
//			Next: res,
//		}
//		res = node
//	}
//	return res
//}
