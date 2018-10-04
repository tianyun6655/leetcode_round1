package review

import "fmt"

func BinarySearch(source []int, target int) int{


	left :=0
	right:=len(source)-1
	for left<right{
         mid := (left+right)/2
         fmt.Println(mid)
         if  source[mid]==target{
			 //tetetetete
			 //tetetetete
			 return mid
		 }else if source[mid]<target{
		 	left=mid+1
		 }else {
		 	right=mid-1
		 }
	}
	return -1
}
