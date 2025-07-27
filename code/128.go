package code

/*
找出最长连续序列的长度
不要求在数组里连续
不能排序 要求时间复杂度在O(n)
*/
func longestConsecutive(nums []int) int {
	//构建一个hashmap,将数字都存入map中
	hashMap := make(map[int]bool)
	for _, n := range nums {
		hashMap[n] = true //构建hashMap
	}
	maxLength := 0
	for n, _ := range hashMap {
		if !hashMap[n-1] {
			//如果前一个数字不存在,则从当前数字开始,前一个数字存在则跳过
			length := 1
			// 开始滑动窗口
			for hashMap[n+1] {
				length++
				n++
			}
			//每次取最大值
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}
