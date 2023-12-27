package main

func quicksort(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	var less []int
	for _, n := range nums[1:] {
		if n < pivot {
			less = append(less, n)
		}
	}

	var greater []int
	for _, n := range nums[1:] {
		if n >= pivot {
			greater = append(greater, n)
		}
	}

	return append(append(quicksort(less), pivot), quicksort(greater)...)
}
