package leetcode

func removeDuplicates(nums []int) int {
	notDuplicateIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[notDuplicateIdx] = nums[i]
			notDuplicateIdx++
		}
	}

	return notDuplicateIdx
}
