package code

import (
	"Leetcode/util"
	"sort"
)

// 如果两个单词是异位词的话
// 那么它们排序完应该就是一样的
func GroupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 {
		return [][]string{strs}
	}
	resultMap := make(map[string][]string)

	for i, str := range strs {
		r := []rune(str)
		sort.Sort(util.SortRunes(r))
		sr := string(r)
		resultMap[sr] = append(resultMap[sr], strs[i])
	}
	result := make([][]string, 0, len(resultMap))
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
