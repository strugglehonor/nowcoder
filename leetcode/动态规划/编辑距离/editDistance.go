package main

func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)
    dp := make([][]int, n1+1)
    for i:=0; i<n1+1; i++ {
        dp[i] = make([]int, n2+1)
    }
    for j:=0; j<n1+1; j++ {
        dp[j][0] = j
    }
    for k:=0; k<n2+1; k++ {
        dp[0][k] = k
    }

    for i:=1; i<n1+1; i++ {
        for j:=1; j<n2+1; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            }else {
                // 到这里已经进行了一次操作，需要+1
                dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
            }
        }
    }
    return dp[n1][n2]
}

func min(nums ...int) int {
    minval := nums[0]
    for i:=0; i<len(nums); i++ {
        if nums[i] < minval {
            minval = nums[i]
        }
    }
    return minval
}