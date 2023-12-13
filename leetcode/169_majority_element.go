package leetcode

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}

	for val, count := range m {
		if count > len(nums)/2 {
			return val
		}
	}

	return 0
}
