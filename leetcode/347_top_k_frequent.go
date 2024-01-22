package leetcode

// O(n log n)
//func topKFrequent(nums []int, k int) []int {
//	countMap := make(map[int]int)
//	res := make([]int, 0, len(nums))
//
//	for _, v := range nums {
//		if _, ok := countMap[v]; !ok {
//			res = append(res, v)
//		}
//
//		countMap[v]++
//	}
//
//	sort.Slice(res, func(i int, j int) bool {
//		return countMap[res[i]] > countMap[res[j]]
//	})
//
//	return res[:k]
//}

// O(n) - bucket sort
func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, value := range nums {
		frequency[value]++
	}

	buckets := make([][]int, len(nums)+1, len(nums)+1)

	for number, freq := range frequency {
		buckets[freq] = append(buckets[freq], number)
	}

	res := make([]int, 0, len(nums))
	for i := len(buckets) - 1; i > 0; i-- {
		if len(buckets[i]) != 0 {
			res = append(res, buckets[i]...)
		}
	}

	return res[:k]
}
