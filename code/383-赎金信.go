package code

func CanConstruct(ransomNote string, magazine string) bool {
	cntChar := make(map[byte]int)
	// 统计magazine中每个字符出现的次数
	for i := 0; i < len(magazine); i++ {
		cntChar[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		cntChar[ransomNote[i]]--
		if cntChar[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
