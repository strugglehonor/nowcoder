package main

func wordBreak(s string, wordDict []string) bool {
	// 存储结果，记忆化递归。表示从begin开头的位置，能否匹配到worddict
    resMap := map[int]bool{}  
	// 哈希表思想
    wordMap := map[string]bool{}
    for _, v := range wordDict {
        wordMap[v] = true
    }
    return dfs(0, wordMap, resMap, s)
}

// dfs函数表示从s开始，是否能够匹配到wordDict，返回布尔值
func dfs(begin int, wordMap map[string]bool, resMap map[int]bool, s string) bool {
    n := len(s)
    // 递归退出条件
    if begin == n {
        return true
    }
    if val, ok := resMap[begin]; ok {
        return val
    }
    for i:=begin+1; i<=n; i++ {
        preS := s[begin:i]
        if wordMap[preS] && dfs(i, wordMap, resMap, s){
            resMap[begin] = true
            return true
        }
    }
    resMap[begin] = false
    return false
}