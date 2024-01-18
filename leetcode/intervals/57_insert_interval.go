package intervals

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	mergedIntervals := make([][]int, 0, len(intervals)+1)

	currIdx := 0

	for currIdx < len(intervals) && newInterval[0] > intervals[currIdx][1] {
		mergedIntervals = append(mergedIntervals, intervals[currIdx])
		currIdx++
	}

	for currIdx < len(intervals) && intervals[currIdx][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[currIdx][0])
		newInterval[1] = max(newInterval[1], intervals[currIdx][1])
		currIdx++
	}

	mergedIntervals = append(mergedIntervals, newInterval)

	for currIdx < len(intervals) {
		mergedIntervals = append(mergedIntervals, intervals[currIdx])
		currIdx++
	}

	return mergedIntervals
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
