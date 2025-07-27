package code

func rotate(nums []int, k int) {
	//轮转数组
	tnums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tnums[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tnums[i]
	}
}
