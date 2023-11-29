package leetcode

func twoSum3(nums []int, target int) []int {
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		if nums[p1]+nums[p2] > target {
			p2--
			continue
		}
		if nums[p1]+nums[p2] < target {
			p1++
			continue
		}

		return []int{p1 + 1, p2 + 1}
	}

	return []int{}
}
