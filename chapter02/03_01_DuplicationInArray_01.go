package main

import "fmt"

/* 数组中的重复数字
 * Leetcode题目链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
 */

func findRepeatNumber02(nums []int) int {
	if nums==nil || len(nums)==0 {
		return -1
	}
	var result int = -1
	for i:=0;i<len(nums);i++{
		for nums[i]!=i{
			if nums[i]==nums[nums[i]] {
				result = nums[i]
				return result
			}
			//swap index=i, nums[i]
			t:=nums[i]
			nums[i]=nums[t]
			nums[t]=t

		}
	}
return result

}

func main(){
	var result int
	result = findRepeatNumber02([]int {2, 3, 1, 0, 2, 5, 3})
	fmt.Println(result)
}


