package code

import "encoding/hex"

/*
将一个字符串以UTF-8形式转为hex编码的字符串
golang好像默认就是UTF-8 也秒了
*/
func Hex_string(str_to_hex string) string {
	// write code here
	//utf8Bytes := []byte(str_to_hex)
	//encodeToString := hex.EncodeToString(utf8Bytes)
	return hex.EncodeToString([]byte(str_to_hex))
}
