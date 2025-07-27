package code

/*
除自身以外数组的乘积
1 2 3 4
L: 1 1 2 6 nums[i-1]*L[i-1]
R: 24 12 4 1 nums[i+1]*R[i+1]
优化空间: 可以优化掉L,R数组,用一个ans数组来存放L,然后用一个变量R来迭代更新答案即可
*/
func ProductExceptSelf(nums []int) []int {
	//l := len(nums)
	//// 开两额外空间
	////分别存放当前元素左边元素乘积和右边元素乘积
	//L, R := make([]int, l), make([]int, l)
	//L[0], R[l-1] = 1, 1
	//i := 1
	//j := l - 2
	//for ; i < l; i++ {
	//	L[i] = L[i-1] * nums[i-1]
	//}
	//for ; j >= 0; j-- {
	//	R[j] = R[j+1] * nums[j+1]
	//}
	////fmt.Println(L, R)
	//for k, _ := range nums {
	//	nums[k] = L[k] * R[k]
	//}
	//return nums
	l := len(nums)
	ans := make([]int, l)
	ans[0] = 1
	R := 1
	//先存L数组
	for i := 1; i < l; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	for i := l - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		// 更新R
		R *= nums[i]
	}
	return ans
}
