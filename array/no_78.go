package array

import "fmt"

/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

给定给一个整型数组，返回这个数组所有的子集（不重复）

Link: https://leetcode.com/problems/subsets/descri,weption/

*/

//09-10-2018（PDT）,暂时有bug, 未解决
func Tianyun_Subsets(nums []int) [][]int {
	result:=make([][]int,0)
	result=append(result,[]int{})
	for i:=0;i<len(nums);i++{
		length:=len(result)
        for j:=0;j<length;j++{
        	fmt.Println(i," :",result)
          result = append(result,append(result[j],nums[i]))
		}
	}
   return result
}
