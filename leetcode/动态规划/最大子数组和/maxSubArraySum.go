package main

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 子数组 是数组中的一个连续部分。
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := nums[0]
	for i:=1; i<n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		}else {
			dp[i] = nums[i] + dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}