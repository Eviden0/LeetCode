package code

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	for i, w := range words {
		ch := pattern[i]
		if word2ch[w] > 0 && word2ch[w] != ch || ch2word[ch] != " " && ch2word[ch] != w {
			return false
		}
		word2ch[w] = ch
		ch2word[ch] = w
	}
	return true
}
