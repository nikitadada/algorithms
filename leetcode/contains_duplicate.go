package leetcode

func containsDuplicate(nums []int) bool {
	existsElements := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := existsElements[v]; ok {
			return true
		}

		existsElements[v] = struct{}{}
	}

	return false
}
