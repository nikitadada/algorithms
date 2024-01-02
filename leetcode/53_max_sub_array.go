package leetcode

func maxSubArray(nums []int) int {
	res := nums[0]
	current := res
	for i := 1; i < len(nums); i++ {
		current = max(nums[i], nums[i]+current)
		res = max(res, current)
	}

	return res
}
