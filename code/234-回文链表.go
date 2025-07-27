package code

/*
判断一个链表是否回文
采取快慢指针的策略

*/
/**
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
//找到中心节点
func FindMidNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

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
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	midNode := FindMidNode(head)
	midNode = ReverseNodeList(midNode)
	nowHead := head
	for midNode != nil && nowHead != nil {
		if midNode.Val != nowHead.Val {
			return false
		}
		midNode = midNode.Next
		nowHead = nowHead.Next
	}
	//复原一下
	ReverseNodeList(midNode)
	return true
}
