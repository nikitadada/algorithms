package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	res := make([]int, 0, len(nums))

	for _, v := range nums {
		if _, ok := countMap[v]; !ok {
			res = append(res, v)
		}

		countMap[v]++
	}

	sort.Slice(res, func(i int, j int) bool {
		return countMap[res[i]] > countMap[res[j]]
	})

	return res[:k]
}
