package leetcode

//func main() {
//	numbers := []int{4, 5, 6, 7, 0, 1, 2}
//
//	println(search(numbers, 0))
//}

func searchInRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
