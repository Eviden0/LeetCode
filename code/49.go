package code

import "sort"

//func groupAnagrams(strs []string) [][]string {
//	// 1. 遍历字符串数组，将每个字符串排序后作为key，原字符串作为value存入map
//	anagrams := make(map[string][]string)
//	for _, str := range strs {
//		sortedStr := sortString(str)
//		anagrams[sortedStr] = append(anagrams[sortedStr], str)
//	}
//
//	// 2. 将map的值存入结果切片
//	result := make([][]string, 0, len(anagrams))
//	for _, group := range anagrams {
//		result = append(result, group)
//	}
//
//	return result
//}
//
//func sortString(str string) string {
//	// 1. 将字符串转换为字符切片
//	runes := []rune(str)
//
//	// 2. 对字符切片进行排序
//	sort.Slice(runes, func(i, j int) bool {
//		return runes[i] < runes[j]
//	})
//
//	// 3. 将排序后的字符切片转换回字符串
//	return string(runes)
//}

func groupAnagrams(strs []string) [][]string {
	//	先排序,相同的单词一定会相同
	//然后将所得的单词作为key,原单词作为value存入map,map的value是一个[]string
	// 遍历map,取出每一个value,将value放入结果切片
	oneSlice := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		oneSlice[sortedStr] = append(oneSlice[sortedStr], str)
	}
	result := make([][]string, 0, len(oneSlice))
	for _, v := range oneSlice {
		result = append(result, v)
	}
	return result
}
func sortString(str string) string {
	// 1. 将字符串转换为字符切片
	runes := []rune(str)

	// 2. 对字符切片进行排序
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// 3. 将排序后的字符切片转换回字符串
	return string(runes)
}
