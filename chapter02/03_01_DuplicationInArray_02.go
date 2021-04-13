package main

import "fmt"

/* 数组中的重复数字
 * Leetcode题目链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
 */

func findRepeatNumber(nums []int) int {
	const n int = 100000
	var a [n] int
	for _,v := range nums{
		if a[v]==0{
			a[v]++
		}else{
			return v
		}
	}
	return nums[0]
}

func main(){
	var result int
	result = findRepeatNumber([]int {2, 3, 1, 0, 2, 5, 3})
	fmt.Println(result)
}