package sliding_window

func numSubarrayProductLessThanK(nums []int, k int) int {
	windowStart, res := 0, 0
	curr := 1
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		curr *= nums[windowEnd]

		for windowStart <= windowEnd && curr >= k {
			curr /= nums[windowStart]
			windowStart++
		}

		res += windowEnd - windowStart + 1
	}

	return res
}
