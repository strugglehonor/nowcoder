package main

func longestPalindrome(s string) string {
    n := len(s)
    dp := make([][]bool, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
	// 最少也要返回一个长度为1的字符串，例如“a”就是回文串，所以maxlen为1
	maxlen, begin, end := 1, 0, 0
    // 枚举长度
    for length:=2; length<n+1; length++ {
        // 枚举左边界
        for left:=0; left<n; left++ {
            // 右边界
            right := length+left-1
            // 右边界越界则退出
            if right >= n {
                break
            }

            if s[left] == s[right] {
                dp[left][right] = dp[left+1][right-1]
            }else {
                dp[left][right] = false
            }

            if s[left] == s[right] && right-left < 3 {
                dp[left][right] = true
            }
			// 注意right-left+1才是长度
			if right-left+1 > maxlen && dp[left][right] {
				maxlen = right-left+1
				begin = left
                end = right
			}
        }
    }
    return s[begin:end+1]
}