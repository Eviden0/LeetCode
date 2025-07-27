package code

/*
买卖股票最佳时机
*/
func MaxProfit(prices []int) int {
	length := len(prices)
	maxPrice := make([]int, length)
	maxPrice[length-1] = 0
	tmpPrice := prices[length-1]
	for i := length - 2; i >= 0; i-- {
		maxPrice[i] = max(tmpPrice-prices[i], 0)
		tmpPrice = max(tmpPrice, prices[i])
	}
	return getMax(maxPrice)
}
func getMax(a []int) int {
	t := -1
	for _, v := range a {
		t = max(v, t)
	}
	return t
}
