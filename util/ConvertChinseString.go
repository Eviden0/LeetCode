package util

func ConvertChinseString(s interface{}) []rune {
	switch v := s.(type) {
	case string:
		return []rune(v)
	case []byte:
		return []rune(string(v))
	default:
		return nil
	}
}
