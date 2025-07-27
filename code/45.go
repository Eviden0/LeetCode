package code

import "fmt"

/*
45. 跳跃游戏 II
给的数据一定能跳到终点,给出所需的最少步数
还是贪心,不过这次没必要每个点都去跳了
可以找到右边界,如果可以直达右边界那么直接过去即可
*/
func Jump(nums []int) int {
	maxDis := 0
	l := len(nums)
	if l == 1 {
		return 0
	}
	cnt := 0
	deadRight := 0
	for k, v := range nums[0 : l-1] {
		//维护最大的可达的点
		if k+v > maxDis {
			fmt.Print(k, "\t", v, "\t", maxDis, "\n")
			maxDis = k + v
		}
		//直接右移边界
		if k == deadRight {
			deadRight = maxDis
			cnt++
		}
	}
	return cnt
}
