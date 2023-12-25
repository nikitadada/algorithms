package leetcode

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if ind, ok := m[v]; ok {
			return []int{i, ind}
		}

		m[target-v] = i
	}

	return []int{}
}
