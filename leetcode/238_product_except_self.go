package leetcode

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l, l)
	right := make([]int, l, l)

	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		left[i] = nums[i-1] * left[i-1]

		j := l - i - 1
		right[j] = nums[j+1] * right[j+1]
	}

	for i := 0; i < l; i++ {
		nums[i] = left[i] * right[i]
	}

	return nums
}
