/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

package main

import "fmt"

func main() {
	r := twoSum([]int{3, 3}, 6)
	fmt.Println(r)
}

func twoSum(arr []int, target int) []int {
	for i, v := range arr {
		if v+arr[i+1] == target {
			return []int{i, i + 1}
		}
	}
	return []int{}
}

/*
func twoSumEfficient(arr []int, target int) []int {
	mapa := make(map[int]int)
	for i, v := range arr {
		complemento := target - v
		if mapa[i] == complemento {
			return mapa[]{complemento}
		}
	}
}

/*