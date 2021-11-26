#  @Author: zijian.su 
#  @Date: 2021-11-26 16:50:10 
#  @Last Modified by:   zijian.su 
#  @Last Modified time: 2021-11-26 16:50:10 

# dp solution
def solve(step, a, b):
    # dp[x][y][1~step]
    dp =[[[0]*9 for _ in range(10)] for _ in range(step+1)]

    print(len(dp), len(dp[0]), len(dp[0][0]))
    dp[0][a][b] = 1
    for z in range(1, step+1):
        for i in range(10):
            for j in range(9):
                # 如果数组越界，则应该返回0，所以需要getValue辅助函数
                dp[z][i][j] = \
                    getValue(i+2, j-1, z-1, step, dp)\
                   +getValue(i+2, j+1, z-1, step, dp)\
                   +getValue(i-2, j+1, z-1, step, dp)\
                   +getValue(i-2, j-1, z-1, step, dp)\
                   +getValue(i+1, j+2, z-1, step, dp)\
                   +getValue(i-1, j+2, z-1, step, dp)\
                   +getValue(i+1, j-2, z-1, step, dp)\
                   +getValue(i-1, j-2, z-1, step, dp)
    return dp[step][a][b]

# 辅助函数
def getValue(i, j, z, step, dp):
    # print(step)
    if i>9 or i<0 or j>8 or j<0 or z<0 or z>step:
        return 0
    else:
        return dp[z][i][j]

if __name__ == '__main__':
    x, y, k = map(int, input().split())
    print(solve(k, x, y))