package code

// 暴力解法
func largestRectangleArea1(heights []int) int {

	l := len(heights)
	if l == 1 {
		return heights[0]
	}
	if heights[0] == heights[1] && heights[0] == heights[l-1] {
		return heights[0] * l
	}
	leftToRight, rightToLeft := make([]int, l), make([]int, l)
	leftToRight[l-1] = l - 1
	rightToLeft[0] = 0
	for i := 0; i < l-1; i++ {
		leftToRight[i] = l - 1
		for j := i + 1; j < l; j++ {
			if heights[j] < heights[i] {
				leftToRight[i] = j - 1
				break
			}
		}

	}
	//fmt.Println("从左往右最大下标", leftToRight)
	for i := l - 1; i > 0; i-- {
		rightToLeft[i] = 0
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				rightToLeft[i] = j + 1
				break
			}
		}

	}
	//fmt.Println("从右往左最大下标", rightToLeft)
	maxArea := -1
	for i := 0; i < l; i++ {
		area := (leftToRight[i] - rightToLeft[i] + 1) * (heights[i])
		if int(area) > maxArea {
			maxArea = int(area)
		}
	}
	return maxArea
}

// 单调栈优化解法
func LargestRectangleArea(heights []int) int {

	stack := make([]int, 0)
	stack = append(stack, -1)
	heights = append(heights, 0)
	l := len(heights)
	res := 0
	for i := 0; i < l; i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			res = max(res, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return res
}
