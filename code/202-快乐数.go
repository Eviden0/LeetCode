package code

/*
最终会形成一个循环：
1. 最终得到1，直接退出
2. 进入循环，直接退出
3. 值越来越大，接近无限大

* 思考第三种情况如何判断呢？
这里有个规律
1位 9 81
2位 99 162
3位 999 243
4位 9999 324
。。。
13位 999999... 1053

因此最大的情况也是 13位->4位
然后就全落入了4位以内
只需判断是否出现重复数即可
*/
func IsHappy(n int) bool {
	isExist := make(map[int]bool)

	for {
		if isExist[n] {
			return false
		}
		isExist[n] = true
		n = step(n)
		if n == 1 {
			return true
		}
	}
}
func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
