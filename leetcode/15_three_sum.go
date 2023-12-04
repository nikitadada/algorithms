package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for idx := 0; idx < len(nums)-1; idx++ {
		if idx != 0 && nums[idx-1] == nums[idx] {
			continue
		}

		p1 := idx + 1
		p2 := len(nums) - 1

		for p1 < p2 {
			if nums[idx]+nums[p1]+nums[p2] < 0 {
				p1++
				continue
			}
			if nums[idx]+nums[p1]+nums[p2] > 0 {
				p2--
				continue
			}

			res = append(res, []int{nums[idx], nums[p1], nums[p2]})
			p1++
			for nums[p1] == nums[p1-1] && p1 < p2 {
				p1++
			}
		}
	}

	return res
}
