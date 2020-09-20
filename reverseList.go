package main

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

var listHead *ListNode
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	recurseList(head)
	head.Next = nil
	return listHead
}

func recurseList(node *ListNode) {
	if node.Next == nil {
		listHead = node
		return
	}
	recurseList(node.Next)
	node.Next.Next = node
}
