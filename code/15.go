package code

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

*/

func ThreeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	//-4 -1 -1 0 1 2
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		// 对a去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for r > l {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				//	找到答案
				result = append(result, sort.IntSlice{nums[i], nums[l], nums[r]})
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				for r > l && nums[l] == nums[l+1] {
					l++
				}
				r--
				l++
			}
		}
	}
	return result
}
