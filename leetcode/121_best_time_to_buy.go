package leetcode

func maxProfit2(prices []int) int {
	left, res := 0, 0

	for right := 1; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else {
			res = max(res, prices[right]-prices[left])
		}
	}

	return res
}
