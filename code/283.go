package code

func MoveZeros(nums []int) {
	//将0移到数组末尾
	//双指针法
	//快指针遍历数组,慢指针记录0的位置
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		//当仅当快指针不为0时交换即可
		//默认慢指针指向的元素为0
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
