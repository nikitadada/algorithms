package leetcode

func rotate(nums []int, k int) {
	res := make([]int, len(nums), len(nums))
	l := len(nums)

	if k > l {
		k = k % l
	}
	firstElIndex := l - k
	curIndex := 0

	for i := firstElIndex; i < l; i++ {
		res[curIndex] = nums[i]
		curIndex++
	}
	for i := 0; i < firstElIndex; i++ {
		res[curIndex] = nums[i]
		curIndex++
	}

	copy(nums, res)
}

//func rotate(nums []int, k int) {
//	n := len(nums)
//	k %= n
//	reverse(nums, 0, n - 1)
//	reverse(nums, 0, k - 1)
//	reverse(nums, k, n - 1)
//}
//
//func reverse(nums []int, start int, end int) {
//	for start < end {
//		nums[start], nums[end] = nums[end], nums[start]
//		start++
//		end--
//	}
//}
