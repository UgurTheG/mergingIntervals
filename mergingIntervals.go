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

	return intervalSlice
}
