# /*
#  * @Author: zijian.su 
#  * @Date: 2021-11-23 17:32:18 
#  * @Last Modified by:   zijian.su 
#  * @Last Modified time: 2021-11-23 17:32:18 
#  */
n, m = map(int, input().split())
matrix = []
for _ in range(n):
    arr = input()
    nums = [int(x) for x in arr.split()]
    matrix.append(nums)
    
dp = [[0]*m for _ in range(n)] # n行m列
# 第0行的元素路径和，以及第0列元素的路径和是确定的
dp[0][0] = matrix[0][0]
for i in range(1, n):
    dp[i][0] = dp[i-1][0] + matrix[i][0]
for j in range(1, m):
    dp[0][j] = dp[0][j-1] + matrix[0][j]
    
for i in range(1, n):
    for j in range(1, m):
        dp[i][j] = min(dp[i-1][j], dp[i][j-1])+matrix[i][j]
print(dp[n-1][m-1])
