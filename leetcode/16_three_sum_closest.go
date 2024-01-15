package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		leftPointer := i + 1
		rightPointer := len(nums) - 1

		for leftPointer < rightPointer {
			currSum := nums[i] + nums[leftPointer] + nums[rightPointer]
			if currSum == target {
				return currSum
			}

			if absInt(target-currSum) < absInt(target-closest) {
				closest = currSum
			}

			if currSum > target {
				rightPointer--
			} else if currSum < target {
				leftPointer++
			}
		}
	}

	return closest
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
