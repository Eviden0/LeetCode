package util

import "fmt"

// 反转链表
func ReverseNodeList(head *ListNode) *ListNode {

	var preNode *ListNode
	nowNode := head
	for nowNode != nil {
		nextNode := nowNode.Next //get the nextNode pointnil
		nowNode.Next = preNode   //reverse
		preNode = nowNode        //reverse
		nowNode = nextNode       //continue next node
	}
	return preNode
}
func ReverseNodeListTest() {
	head := ListNode{
		Var: 1,
	}
	head.Next = &ListNode{
		Var: 2,
	}
	head.Next.Next = &ListNode{
		Var: 3,
	}
	var nowNode *ListNode
	nowNode = &head
	fmt.Printf("初始链表为:   ")

	for nowNode != nil {
		fmt.Printf("\t%d", nowNode.Var)
		nowNode = nowNode.Next
	}
	fmt.Println("\n===========================")
	nowNode = ReverseNodeList(&head)
	fmt.Printf("反转后的链表为: ")
	for nowNode != nil {
		fmt.Printf("\t%d", nowNode.Var)
		nowNode = nowNode.Next
	}
}
