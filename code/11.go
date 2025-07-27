package code

// 双指针解决
// 容量为 两点坐标距离乘以最小值
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	maxA := -1
	for {
		if l >= r {
			return maxA
		} else {
			temp := (r - l) * min(height[r], height[l])
			if temp > maxA {
				maxA = temp
			}
			//	 移动指针
			if height[r] >= height[l] {
				l++
			} else {
				r--
			}
		}
	}
}
