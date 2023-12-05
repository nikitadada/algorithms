package leetcode

func validMountainArray(arr []int) bool {
	p1, p2 := 0, len(arr)-1
	for p1 < p2 {
		if arr[p1+1] > arr[p1] {
			p1++
			continue
		}

		if arr[p2-1] > arr[p2] {
			p2--
			continue
		}

		return false
	}

	if p1 == 0 || p2 == len(arr)-1 {
		return false
	}

	return true
}
