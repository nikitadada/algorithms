package leetcode

//func moveZeroes(nums []int) {
//	countIterations := len(nums)
//	for i := 0; i < countIterations; {
//		if nums[i] == 0 {
//			nums = append(nums[:i], nums[i+1:]...)
//			nums = append(nums, 0)
//			countIterations--
//			continue
//		}
//		i++
//	}
//}

func moveZeroes(nums []int) {
	l := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[l] = nums[i]
			l++
		}
	}

	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
}

//func moveZeroes(nums []int) {
//	l := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != 0 {
//			nums[i], nums[l] = nums[l], nums[i]
//			l++
//		}
//	}
//}
