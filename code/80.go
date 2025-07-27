package code

func removeDuplicates(nums []int) int {
	l := len(nums)
	//用stack去理解实现
	if l < 3 {
		return l
	}
	stackSize := 2 //至少先压入两个元素
	for i := 2; i < l; i++ {
		if nums[i] != nums[stackSize-2] {
			nums[stackSize] = nums[i]
			stackSize++
		}
	}
	return stackSize
}
