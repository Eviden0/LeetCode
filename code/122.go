package code

/*
122. 买卖股票的最佳时机 II
每一天都可以买和卖出,贪心即可
*/
func maxProfit2(prices []int) (ans int) {
	ans = 0
	for i := 1; i < len(prices); i++ {
		ans += max(prices[i]-prices[i-1], 0)
	}
	return ans
}
