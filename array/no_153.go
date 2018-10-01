package array


/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.
*/


func Tianyun_FindMin(nums []int) int {
    if len(nums)<=0{
    	return -1
	}
	left :=0
	right:= len(nums)-1

	if nums[left]<nums[right]{
		return nums[left]
	}
	for left +1 <right{
		//如果已经排好序
		if nums[left]<nums[right]{
			return nums[left]
		}

		mid:= (left+right)/2
		if nums[mid]>nums[left]{
			left = mid+1
		}else{
			right = mid
		}
	}
	if nums[left]<nums[right]{
		return nums[left]
	}else {
		return nums[right]
	}
}