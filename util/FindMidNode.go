package util

import "fmt"

/*
找到一个链表的中心
使用快慢指针
慢指针一次走一步
快指针一次走两步
同时从节点出发,当快指针到达的时候则慢指针为链表中心
*/

// 找到链表中心节点
func FindMidNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func FindMidNodeTest() {
	head := ListNode{
		Var: 1,
	}
	head.Next = &ListNode{
		Var: 2,
	}
	head.Next.Next = &ListNode{
		Var: 3,
	}
	//head.Next.Next.Next = &ListNode{
	//	Var: 4,
	//}
	midNode := FindMidNode(&head)
	fmt.Println(midNode.Var)
}
