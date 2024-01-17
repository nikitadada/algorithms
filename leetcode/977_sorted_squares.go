package leetcode

func sortedSquares(nums []int) []int {
	p1 := 0
	p2 := len(nums) - 1
	res := make([]int, len(nums), len(nums))
	curInd := len(nums) - 1

	for p1 <= p2 {
		leftSquare := nums[p1] * nums[p1]
		rightSquare := nums[p2] * nums[p2]
		if leftSquare > rightSquare {
			res[curInd] = leftSquare
			p1++
		} else {
			res[curInd] = rightSquare
			p2--
		}

		curInd--
	}

	return res
}
