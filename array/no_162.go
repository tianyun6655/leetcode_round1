package array


func Tianyun_findPeakElement(nums []int) int {

	//O(n)
	//if len(nums)<=1{
	//	return 0
	//}
	//
	//for i:=0;i<len(nums)-1;i++{
	//	if i>=1{
	//		if nums[i]>nums[i-1]&&nums[i]>nums[i+1]{
	//			return i
	//		}
	//	}else if i==0{
	//		if nums[i]>nums[i+1]{
	//			return i
	//		}
	//	}
	//}
	//
	//if nums[len(nums)-1]>nums[len(nums)-2]{
	//	return len(nums)-1
	//}
	//
	//return 0

	left:=0
	right:=len(nums)-1

	if len(nums)<=1{
		return 0
	}

	for left+1<right{
		mid:=(left+right)/2
		if nums[mid]>nums[mid+1]{
			right = mid
		}else{
			left=mid+1
		}
	}
	if left==len(nums)-1{
		return left
	}
	if nums[left]>nums[left+1]{
		return left
	}else{
		return left+1
	}


	//testgit
}