package main

import (
	"Leetcode/code"
	"fmt"
)

/*
实现常见数据结构
*/
// 函数变量（包级作用域）
var add func(int, int) int = func(a, b int) int {
	return a + b
}

func main() {
	test := []int{2, 1, 5, 6, 2, 3}
	area := code.LargestRectangleArea(test)
	fmt.Println(area)
	l := len(test)

	newHeights := make([]int, l+2)
	copy(newHeights[1:l+1], test)
	fmt.Println(newHeights)

}
