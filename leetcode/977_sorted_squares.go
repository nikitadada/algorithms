package leetcode

func sortedSquares(nums []int) []int {
	l := len(nums)
	leftPointer, rightPointer := 0, l-1
	squares := make([]int, l, l)
	currSquaresIdx := l - 1

	for leftPointer <= rightPointer {
		leftSquare := nums[leftPointer] * nums[leftPointer]
		rightSquare := nums[rightPointer] * nums[rightPointer]

		if leftSquare > rightSquare {
			squares[currSquaresIdx] = leftSquare
			leftPointer++
		} else {
			squares[currSquaresIdx] = rightSquare
			rightPointer--
		}

		currSquaresIdx--
	}

	return squares
}
