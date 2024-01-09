package leetcode

// Первое решение, выглядит сложнее
func searchMatrix2(matrix [][]int, target int) bool {
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1

	for top <= bottom {
		midRow := (top + bottom) / 2
		midColumn := (left + right) / 2

		if target >= matrix[midRow][left] && target <= matrix[midRow][right] {
			if matrix[midRow][midColumn] == target {
				return true
			}

			if matrix[midRow][midColumn] > target {
				right = midColumn - 1
			} else {
				left = midColumn + 1
			}
		} else {
			if matrix[midRow][left] > target {
				bottom = midRow - 1
			} else {
				top = midRow + 1
			}
		}
	}

	return false
}

// Решение более простое для понимания. Сначала находим строку, потом делаем обычный бинарный поиск по строке.
func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1

	// find row
	var row []int
	for top <= bottom {
		midRow := (top + bottom) / 2
		row = matrix[midRow]

		if target < row[0] {
			bottom = midRow - 1
		} else if target > row[len(row)-1] {
			top = midRow + 1
		} else {
			break
		}
	}

	if len(row) == 0 {
		return false
	}

	// binary search in row
	left, right := 0, len(row)-1

	for left <= right {
		mid := (left + right) / 2
		if row[mid] == target {
			return true
		}

		if row[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
