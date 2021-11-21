/*
 * @Author: zijian.su
 * @Date: 2021-11-21 19:47:14 
 * @Last Modified by:   zijian.su
 * @Last Modified time: 2021-11-21 19:47:14 
 */
package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var N int
	fmt.Scanln(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	length, maxval, minval := 0, 0, 0
	// 以i开头，以j结尾
	for i := 0; i < N; i++ {
		minval = int(^uint(0) >> 1)
		maxval = -minval
		hashMap := map[int]int{}

		for j := i; j < N; j++ {
			if _, ok := hashMap[arr[j]]; ok {
				break
			}
			hashMap[arr[j]] = 0
			minval = min(arr[j], minval)
			maxval = max(arr[j], maxval)
			if j-i == maxval-minval {
				length = max(length, j-i+1)
			}
		}
	}
	fmt.Println(length)
}
