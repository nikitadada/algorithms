package leetcode

// Плохое решение O(n^2)
func twoSum(nums []int, target int) []int {
	p1, p2 := 0, 1
	for p1 < len(nums)-1 && p2 < len(nums) {
		if nums[p1]+nums[p2] == target {
			return []int{p1, p2}
		}

		if p2 == len(nums)-1 {
			p1++
			p2 = p1 + 1
			continue
		}

		p2++
	}

	return []int{}
}
