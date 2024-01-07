package leetcode

//func main() {
//	numbers := []int{1, 2, 5, 6, 7, 8, 9, 10, 11, 14, 15, 16}
//
//	println(binarySearch(numbers, 23))
//}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}
