package code

import "sort"

func majorityElement(nums []int) int {
	//给定一个大小为 n 的数组 nums ，返回其中的多数元素。
	//多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
	sort.Ints(nums)

	return nums[len(nums)/2]
}
