package intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][]int, 0)
	mergedIntervals = append(mergedIntervals, intervals[0])

	for _, interval := range intervals[1:] {
		currInterval := mergedIntervals[len(mergedIntervals)-1]

		if interval[0] > currInterval[1] {
			mergedIntervals = append(mergedIntervals, interval)
		} else if interval[1] > currInterval[1] {
			currInterval[1] = interval[1]
		}
	}

	return mergedIntervals
}
