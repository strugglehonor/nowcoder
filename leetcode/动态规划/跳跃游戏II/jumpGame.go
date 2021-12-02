package main

func jump(nums []int) int {
    n := len(nums)
    maxPosition := 0 // 能跳跃的最大位置
    steps := 0
    end := 0 // 到达的最右边的位置
    // 这里是n-1，因为最后一次不算
    for i:=0; i<n-1; i++ {
        maxPosition = max(i+nums[i], maxPosition)
        if i == end {
            steps++
            end = maxPosition
        }
    }
    return steps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}