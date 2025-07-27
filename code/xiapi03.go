package code

/*
将一万以下的中文数字转为阿拉伯数字
比如 九千九百九十九 -> 9999
*/
type chineseChar rune

func Convert_num(number_in_chinese string) int {
	//dict1 := map[chineseChar]int{
	//	'十': 10,
	//	'百': 100,
	//	'千': 1000,
	//}
	//dict2 := map[chineseChar]int{
	//	'一': 1,
	//	'二': 2,
	//	'三': 3,
	//	'四': 4,
	//	'五': 5,
	//	'六': 6,
	//	'七': 7,
	//	'八': 8,
	//	'九': 9,
	//}
	runes := []rune(number_in_chinese) //正确的做法是先转成rune再进行遍历
	result := 0
	for i := 0; i < len(runes); i++ {

	}

	return result
}
