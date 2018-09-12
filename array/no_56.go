package array

/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.
*/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval struct {
	Start int
	End   int
}

func Tianyun_merge(intervals []Interval) []Interval {

	Tianyun_sort(intervals)

	for i := 0; i < len(intervals)-1; {
		if intervals[i].End >= intervals[i+1].Start {
			var newInter Interval
			if intervals[i].End <= intervals[i+1].End {
				newInter.Start = intervals[i].Start
				newInter.End = intervals[i+1].End
			} else {

				newInter.Start = intervals[i].Start
				newInter.End = intervals[i].End

			}
			intervals[i] = newInter
			intervals = append(intervals[:i+1], intervals[i+2:]...)

		} else {
			i++
		}
	}
	return intervals

}


func Tianyun_sort(intervals []Interval){
	quickSortHelper(intervals,0,len(intervals)-1)
}


func quickSortHelper(intervals []Interval,left,right int){
	if left>right{
		return
	}
	base:=partition(intervals,left,right)

	quickSortHelper(intervals,left,base-1)

	quickSortHelper(intervals,base+1,right)
}


func partition(intervals []Interval,left,right int)int{
    base:=intervals[left].Start

    j:=left
    for i:=j+1;i<=right;i++{
    	if intervals[i].Start<base{
    		intervals[j+1],intervals[i]=intervals[i],intervals[j+1]
    		j++
		}
	}
	intervals[j],intervals[left]=intervals[left],intervals[j]

	return j
}
