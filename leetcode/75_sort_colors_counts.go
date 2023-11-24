package leetcode

// TODO переделать через sliding window. Сейчас count sort подход
func sortColors(nums []int) {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	ind := 0
	for i := ind; i < m[0]; i++ {
		nums[ind] = 0
		ind++
	}

	for i := 0; i < m[1]; i++ {
		nums[ind] = 1
		ind++
	}

	for i := 0; i < m[2]; i++ {
		nums[ind] = 2
		ind++
	}
}
