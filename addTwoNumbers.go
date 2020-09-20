package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1, listLen1 := reverseNumList(l1)
	list2, listLen2 := reverseNumList(l2)
	sumList := list1
	otherList := list2
	if listLen1 < listLen2 {
		sumList = list2
		otherList = list1
	}
	carry := 0
	sumListHead := sumList
	forward := sumList
	for otherList != nil {
		sumList.Val = sumList.Val + otherList.Val + carry
		carry = sumList.Val / 10
		sumList.Val = sumList.Val % 10

		forward = sumList
		otherList = otherList.Next
		sumList = sumList.Next
	}
	for sumList != nil {
		sumList.Val += carry
		carry = sumList.Val / 10
		sumList.Val = sumList.Val % 10
		forward = sumList
		sumList = sumList.Next
	}
	if carry == 1 {
		forward.Next = &ListNode{Val:carry, Next:nil,}
	}
	retList, _ := reverseNumList(sumListHead)
	return retList
}

func reverseNumList(list *ListNode) (*ListNode,int) {
	var forward, curNode *ListNode
	curNode = list
	forward = nil
	listLen := 0
	for curNode.Next != nil {
		nextNode := curNode.Next
		curNode.Next = forward
		forward = curNode
		curNode = nextNode
		listLen ++
	}
	curNode.Next = forward
	listLen ++
	return curNode, listLen
}

func main() {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	arr1 := [3]int{1, 2, 3}
	var node1, node2 *ListNode
	l1Head := l1
	l2Head := l2
	for _, value := range arr1 {
		node1 = &ListNode{Val:value, Next:nil}
		l1.Next = node1
		l1 = l1.Next
		node2 = &ListNode{Val:value+1, Next:nil}
		l2.Next = node2
		l2 = l2.Next
	}
	ret := addTwoNumbers(l1Head, l2Head)
	fmt.Printf("%v", ret)

}
