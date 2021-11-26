package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanln(&N)
    arr := make([]int, N)
    // 两个参数left和right可确定一个值，所以采用二维数组
    dpF := make([][]int, N)
    dpS := make([][]int, N)
    for i:=0; i<N; i++ {
        fmt.Scan(&arr[i])
        dpF[i] = make([]int, N)
        dpS[i] = make([]int, N)
    }
    for i:=0; i<N; i++ {
        dpF[i][i] = arr[i]
    }
    // 从对角线起点开始遍历 对角线起点从 0,1 开始 0,2 -> 0, 3
    for col:=1; col<N; col++ {
        left, right := 0, col
        for left < N && right < N {
            dpF[left][right] = max(dpS[left+1][right] + arr[left], dpS[left][right-1] + arr[right])
            dpS[left][right] = min(dpF[left+1][right], dpF[left][right-1])
            left++ 
            right++
        }
        
    }
    fmt.Printf("%d", max(dpF[0][N-1], dpS[0][N-1]))
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}