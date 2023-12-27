package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := 0
	p2 := 0
	res := make([]int, 0, m+n)

	for p1 != m || p2 != n {
		if p1 != m && (p2 == n || nums1[p1] < nums2[p2]) {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}
	}

	copy(nums1, res)
}
