package sliding_window

func numOfSubarrays(arr []int, k int, threshold int) int {
	windowStart := 0
	res := 0
	sum := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]

		if windowEnd-windowStart+1 == k {
			average := sum / k

			if average >= threshold {
				res++
			}

			sum -= arr[windowStart]
			windowStart++
		}
	}

	return res
}
