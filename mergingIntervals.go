package main

import "sort"

func Merge(intervalSlice [][]int) [][]int {
	if len(intervalSlice) < 2 {
		return intervalSlice
	}
	// sort by startpoint of interval
	sort.Slice(intervalSlice, func(i, j int) bool {
		return intervalSlice[i][0] < intervalSlice[j][0]
	})

	returnedIntervals := [][]int{}

	tmpIntervalStart := intervalSlice[0][0]
	tmpIntervalEnd := intervalSlice[0][1]

	for _, interval := range intervalSlice[0:] {

		if tmpIntervalEnd > interval[0] && tmpIntervalEnd < interval[1] {
			tmpIntervalEnd = interval[1]
		} else {
			tmpIntervalStart = interval[0]
			tmpIntervalEnd = interval[1]
		}
		returnedIntervals = append(returnedIntervals, []int{tmpIntervalStart, tmpIntervalEnd})
	}

	return returnedIntervals
}
