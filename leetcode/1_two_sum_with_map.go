package leetcode

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		m[target-v] = i
	}

	for i, v := range nums {
		if k, ok := m[v]; ok {
			if i != k {
				return []int{i, k}
			}
		}
	}

	return []int{}
}
