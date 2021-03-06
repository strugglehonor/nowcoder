## 排成一条线的纸牌博弈问题

### 题目描述
给定一个整型数组arr，代表数值不同的纸牌排成一条线，玩家A和玩家B依次拿走每张纸牌，规定玩家A先拿，玩家B后拿，但是每个玩家每次只能拿走最左和最右的纸牌，玩家A和玩家B绝顶聪明。请返回最后的获胜者的分数。

输入描述:
输出包括两行，第一行一个整数n，代表数组arr长度，第二行包含n个整数，第i个代表arr[i]。

输出描述:
输出一个整数，代表最后获胜者的分数。

输入
4
1 2 100 4

输出
101

### 解题
这道题可以使用动态规划或者暴力解法（递归）解题，那么关键是如何从暴力解法演变为动态规划
由暴力递归的思路：
- 设定两个函数，一个代表先手方获得的纸牌点数f(arr, left, right)，一个是后手方代表的纸牌点数S(arr, left, right)。
- f()的base case是left==right时，返回arr[left]，而S()的base case是left==right时，返回0（后手只有一个牌的时候，这个牌被先手拿了)
- 非base case的时候，
- f(left, right)为max(S(left+1, right) + arr[left]， S(left, right-1) + arr[right]) 
- S(left, right)为min(f(left+1, right), f(left, right-1))
动态规划套路如下：
- 首先看哪几个参数可以确定我们的结果，根据题意，这里left和right两个参数可以确定结果，所以需要二维的表
- 先手方和后手方分别开一个表
- 把这两个表画出来，结果我们要返回的就是max(dpF[0][N-1], dpS[0][N-1])
- 先初始化这两个表，当left=right时，dpF[i][i] = arr[i], dpS[i][i] = 0
- 转移，根据上述的暴力递归的思路进行动态规划的状态转移。