package code

import "container/list"

/*
有效括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	hash := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack.PushBack(s[i])
		} else {
			if stack.Len() > 0 && hash[s[i]] == stack.Back().Value.(byte) {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		}
	}
	return stack.Len() == 0
}
