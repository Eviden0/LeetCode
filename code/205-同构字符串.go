package code

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t, t2s := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
