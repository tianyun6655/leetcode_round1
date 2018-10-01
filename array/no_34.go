package array


/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

*/


func Tianyun_searchRange(nums []int, target int) []int {

	//O(n)
//	startIndex:=-1
//	count:=0
//
//	for i:=0;i<len(nums);i++{
//		if count==0&&nums[i]==target{
//			startIndex = i
//			count++
//
//		}else if nums[i]==target{
//			count++
//		}
//}
//
// if count==0{
// 	return []int{-1,-1}
// }else{
// 	return []int{startIndex,startIndex+count}
// }



//O(logn)
if len(nums)==0{
	return []int{-1,-1}
}
      startIndex:=-1
      left:=0
      right:=len(nums)-1

      for left +1 <right{
      	mid:=(left+right)/2
      	if nums[mid]>=target{
           right = mid
		}else{
			left =mid+1
		}
	  }
	  if nums[left]==target{
	  	startIndex =left
	  }else if nums[right]==target{
	  	startIndex = right
	  }else{
	  	return []int{-1,-1}
	  }


	left=0
	right=len(nums)-1

	for left +1 <right{
		mid:=(left+right)/2
		if nums[mid]<=target{
			left = mid
		}else{
			right = mid
		}
	}

	if nums[right]==target{
		return []int{startIndex,right}
	}else if nums[left]==target{
		return []int{startIndex,left}
	}else{
		return []int{-1,-1}
	}
}