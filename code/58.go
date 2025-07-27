package code

func LengthOfLastWord(s string) int {

	if len(s) == 1 {
		return 1
	}
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		} else {
			//	找到了往前走即可
			for i >= 0 && s[i] != ' ' {
				i--
				cnt++
			}
			break
		}
	}
	return cnt
}
