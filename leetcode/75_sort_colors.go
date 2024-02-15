package leetcode

//func sortColors(nums []int) {
//	m := make(map[int]int)
//
//	for _, n := range nums {
//		m[n]++
//	}
//
//	ind := 0
//	for i := ind; i < m[0]; i++ {
//		nums[ind] = 0
//		ind++
//	}
//
//	for i := 0; i < m[1]; i++ {
//		nums[ind] = 1
//		ind++
//	}
//
//	for i := 0; i < m[2]; i++ {
//		nums[ind] = 2
//		ind++
//	}
//}

// Наиболее оптимальное решение
func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
}
