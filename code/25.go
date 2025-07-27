package code

/*
k 个一组翻转链表
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
有一种思路就是先全部反转,遇到不足再反转回去的
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var prev *ListNode
	for head != nil && head.Next != nil && isKGroup(head.Next, k) {

	}
	return prev
}

// 判断能否走k步
func isKGroup(head *ListNode, k int) bool {
	cnt := 0
	for head != nil && head.Next != nil {
		cnt++
		head = head.Next
		if cnt >= k {
			return true
		}
	}
	return false
}
