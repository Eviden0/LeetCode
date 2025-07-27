package code

/*
跳跃游戏
*/
func CanJump(nums []int) bool {

	maxDis := nums[0]
	l := len(nums)
	if l == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	for k, v := range nums[1:] {
		if maxDis >= l-1 {
			return true
		}
		//fmt.Print(k, "\t", v, "\n")
		//连最近的点都无法到达,必不能到达最后的
		if maxDis < k+1 {
			return false
		}
		//维护最大的可达距离
		if k+v+1 > maxDis {
			maxDis = k + v + 1
		}
	}
	return false
}
