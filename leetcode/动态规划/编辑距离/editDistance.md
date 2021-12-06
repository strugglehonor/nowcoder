## 编辑距离
https://leetcode-cn.com/problems/edit-distance 

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

## 示例
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

## 解题
1. 想好怎么尝试，枚举操作，这里一共有3种操作，插入，删除，替换
2. 使用动态规划解题，这里两个参数确定问题答案，所以应该使用二维dp表，dp[i][j]表示从word1[:i]转化为word2[:j]所需的最少操作
3. 草稿画出动态规划表，推导如何填表。
4. 当word1[i-1]!=word2[i-1]，dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])，这三个操作分别表示删除，插入，替换
可参考：https://leetcode-cn.com/problems/edit-distance/solution/zi-di-xiang-shang-he-zi-ding-xiang-xia-by-powcai-3/