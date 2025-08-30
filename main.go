package main

import (
	"Leetcode/code"
	"fmt"
)

/*
实现常见数据结构
*/
// 函数变量（包级作用域）
var add = func(a, b int) int {
	return a + b
}

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	fmt.Println(code.ContainsNearbyDuplicate(nums, k))
}
