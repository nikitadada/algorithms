package sliding_window

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	windowStart, minSubArraySize := 0, math.MaxInt
	currentWindowSum := 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		currentWindowSum += nums[windowEnd]

		for currentWindowSum >= target {
			minSubArraySize = min(minSubArraySize, windowEnd-windowStart+1)

			currentWindowSum -= nums[windowStart]
			windowStart++
		}
	}

	if minSubArraySize == math.MaxInt {
		return 0
	}

	return minSubArraySize
}
