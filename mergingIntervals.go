package mergingIntervals

import "sort"

const (
	// FirstIntervalElement is used to index the first interval in the slice of intervals.
	FirstIntervalElement = 0
	// StartOfInterval is used to index the startpoint of an interval.
	StartOfInterval = 0
	// EndOfInterval is used to index the endpoint of an interval.
	EndOfInterval = 1
	// IntervalLengthCheck is used to check the length of an interval-slice.
	IntervalLengthCheck = 2
)

// Interval provides functions to merge and sort a slice of intervals.
type Interval struct{}

// Merge invokes helper-functions in order to sort and return a slice of merged intervals.
func (i Interval) Merge(intervalSlice [][]int) [][]int {

	// If there are less than two Intervals in the slice of Intervals we can return the slice of intervals without doing anything.
	if len(intervalSlice) < IntervalLengthCheck {
		return intervalSlice
	}
	i.sortIntervalsByStart(intervalSlice)
	return i.mergeIntervals(intervalSlice)

}

// mergeIntervals returns a slice of merged intervals.
func (i Interval) mergeIntervals(intervals [][]int) [][]int {

	returnedIntervals := [][]int{}
	// the start and endpoint of the first interval in the slice is used as a comparison for the following intervals.
	tmpIntervalStart := intervals[FirstIntervalElement][StartOfInterval]
	tmpIntervalEnd := intervals[FirstIntervalElement][EndOfInterval]

	for _, intervalElement := range intervals {
		// if: If the current interval can be merged with the previous interval, the endpoint of the previous interval is updated with the endpoint of the current interval.
		if tmpIntervalEnd >= intervalElement[StartOfInterval] && tmpIntervalEnd < intervalElement[EndOfInterval] {
			tmpIntervalEnd = intervalElement[EndOfInterval]
			// elseif: If the current interval is unmergeable with the previous interval, append it to the returnedIntervals slice.
			// Next, the start and endpoint of the current interval is used for comparison as tmpIntervalStart and tmpIntervalEnd.
		} else if tmpIntervalEnd < intervalElement[StartOfInterval] {
			returnedIntervals = append(returnedIntervals, []int{tmpIntervalStart, tmpIntervalEnd})
			tmpIntervalStart = intervalElement[StartOfInterval]
			tmpIntervalEnd = intervalElement[EndOfInterval]
		}
	}

	returnedIntervals = append(returnedIntervals, []int{tmpIntervalStart, tmpIntervalEnd})
	return returnedIntervals
}

// sortIntervalsByStart sorts the slice of intervals by its startpoints.
func (i Interval) sortIntervalsByStart(intervals [][]int) {
	sort.Slice(intervals, func(firstIntervalIndex, secondIntervalIndex int) bool {
		return intervals[firstIntervalIndex][StartOfInterval] < intervals[secondIntervalIndex][StartOfInterval]
	})
}
