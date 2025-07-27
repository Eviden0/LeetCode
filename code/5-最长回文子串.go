package code

/*
给一个字符串,找出s中最长的回文子串
用动态规划进行解决

(1). 一个字符本身就是回文
(2). 长度为2的字符串首尾相同就是回文
(3). 长度>2如果首尾相同，那么去掉首尾后是否是回文就决定了它是否是回文;
如果首尾不同，则一定不是回文
2. 选用哪种算法解题?
抛去前两种case，第三种case是典型的的可以使用动态规划解题的特征。即后面的判断可以以前面的结算结果作为依据。
*/
func longestPalindrome(s string) string {
	sLen := len(s)
	dp := make([][]bool, sLen)
	for i, _ := range s {
		dp[i] = make([]bool, sLen)
	}
	result := s[0:1]
	for l := 0; l < sLen; l++ {
		for start := 0; start+l < sLen; start++ {
			end := start + l
			switch l {
			// Any single character is a palindrome
			case 0:
				//single char must be palindromic
				dp[start][end] = true
				//Two characters are a palindrome if they are the same
			case 1:
				//start and end must be same,can extend
				dp[start][end] = (s[start] == s[end])
				//dp from the front
			default:
				dp[start][end] = (s[start] == s[end] && dp[start+1][end-1])
			}
			//update result if a longer palindrome is found
			if dp[start][end] && l+1 > len(result) {
				result = s[start : end+1]
			}
		}
	}
	return result
}
