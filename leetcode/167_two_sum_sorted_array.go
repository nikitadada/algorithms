package leetcode

func twoSum3(nums []int, target int) []int {
	leftPointer, rightPointer := 0, len(nums)-1

	for leftPointer < rightPointer {
		currSum := nums[leftPointer] + nums[rightPointer]
		if currSum > target {
			rightPointer--
		} else if currSum < target {
			leftPointer++
		} else {
			return []int{leftPointer + 1, rightPointer + 1}
		}
	}

	return []int{}
}
