package code

import "sort"

/*
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
返回前k个高频元素
对map的value 做排序
*/
type Pair struct {
	Key int
	Val int
}

func topKFrequent(nums []int, k int) []int {
	var boolFunc func(i, j int) bool
	boolFunc = func(i, j int) bool {
		return i > j
	}

	hashMap := make(map[int]int)
	for _, n := range nums {
		hashMap[n]++
	}
	pairs := make([]Pair, len(hashMap))
	i := 0
	for k, v := range hashMap {
		pairs[i] = Pair{k, v}
		i++
	}
	sort.Slice(pairs, boolFunc)
	result := make([]int, 0, k)
	for i = 0; i < k; i++ {
		result = append(result, pairs[i].Key)
	}
	return result
}
