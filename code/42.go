package code

/*
height = [0, 1,0,2 ,1,0,1,3, 2,1,2 ,1]
输出：6
坐标i处能接的雨水就是左右两边比当前高的柱子中小者与自己的差
*/
func Trap(height []int) int {
	length := len(height) - 1
	leftMax := make([]int, length+1)
	rightMax := make([]int, length+1)
	//左边最大的就从左边开始扫描,右边最大就从右边开始
	//两边肯定不装水,默认边上最高都为0
	leftMax[0] = 0
	rightMax[length] = 0
	tmpMax := height[0]
	for i := 1; i <= length; i++ {
		leftMax[i] = tmpMax
		if height[i] > tmpMax {
			tmpMax = height[i]
		}
	}
	tmpMax = height[length]
	for i := length - 1; i >= 0; i-- {
		rightMax[i] = tmpMax
		if height[i] > tmpMax {
			tmpMax = height[i]
		}
	}
	//fmt.Println("左边最大的: ", leftMax)
	//fmt.Println("右边最大的: ", rightMax)
	tmpMax = 0
	for i := 1; i <= length; i++ {
		tmpMax += max((min(leftMax[i], rightMax[i]) - height[i]), 0)
	}
	//变量重用~
	return tmpMax
}
