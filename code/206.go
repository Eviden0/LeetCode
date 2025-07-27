package code

/*
反转链表
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	/*
			反向存储即可
			本来是 1->2->3->4
			变为  1<-2<-3<-4
		也即交换每一个相邻节点的引用即可
	*/
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next //先取到下一个指针进行存储,防止丢失
		curr.Next = prev  //置为上一个prev,头指针赋为nil即可
		prev = curr       //取到当前节点指针,作为下一次的反转进行存储
		curr = next       //递增到下一个节点
	}
	return prev
}
