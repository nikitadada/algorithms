package leetcode

func findMaxAverage(nums []int, k int) float64 {
	left, sum := 0, 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i]

		if i-left == k {
			sum -= nums[left]
			left++
		}

		res = max(res, sum)
	}

	return float64(res) / float64(k)
}
