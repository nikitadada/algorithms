package leetcode

import "sort"

// Реализация за n * log(n) без дополнительной аллокации памяти
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		}
	}

	return res
}

// Реализация за линейное время O(n), но с дополнительной аллокацией памяти
func intersect2(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int

	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if count, ok := m[v]; ok {
			if count > 0 {
				res = append(res, v)
				m[v]--
			}
		}
	}

	return res
}
