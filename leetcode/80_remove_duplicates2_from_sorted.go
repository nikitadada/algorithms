package leetcode

func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	notDuplicateIdx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[notDuplicateIdx-2] {
			nums[notDuplicateIdx] = nums[i]
			notDuplicateIdx++
		}
	}

	return notDuplicateIdx
}
