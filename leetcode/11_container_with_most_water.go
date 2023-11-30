package leetcode

func maxArea(height []int) int {
	var maxArea int
	p1, p2 := 0, len(height)-1

	for p1 < p2 {
		curMaxArea := min(height[p1], height[p2]) * (p2 - p1)
		maxArea = max(maxArea, curMaxArea)

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}
