package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		leftPointer := i + 1
		rightPointer := len(nums) - 1

		for leftPointer < rightPointer {
			currSum := nums[i] + nums[leftPointer] + nums[rightPointer]
			if currSum > 0 {
				rightPointer--
			} else if currSum < 0 {
				leftPointer++
			} else {
				res = append(res, []int{nums[i], nums[leftPointer], nums[rightPointer]})
				break
			}
		}
	}

	return res
}
