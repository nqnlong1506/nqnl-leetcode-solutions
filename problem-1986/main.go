package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("problem 1986")
	fmt.Println("result:", minSessions([]int{2, 3, 3, 4, 4, 4, 5, 6, 7, 10}, 12))
}

// func minSessions(tasks []int, sessionTime int) int {
// 	subSessions := []int{}

// 	for _, t := range tasks {
// 		flag := false
// 		for i, s := range subSessions {
// 			if t+s <= sessionTime {
// 				subSessions[i] = s + t
// 				flag = true
// 				break
// 			}
// 		}

// 		if !flag {
// 			subSessions = append(subSessions, t)
// 		}
// 	}

// 	fmt.Println(subSessions)

// 	return len(subSessions)
// }

func minSessions(tasks []int, sessionTime int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	bitmask := 0
	for i := range tasks {
		bitmask ^= 1 << i
	}
	dp := make(map[int]int)
	dp[0] = 0
	return dynamicProgramming(tasks, dp, sessionTime, bitmask)
}

func dfs(tasks []int, index, res, bitmask int) (int, int) {
	for i := index; i < len(tasks); i++ {
		if bitmask&(1<<i) == 0 || tasks[i] > res {
			continue
		}
		h1, b1 := dfs(tasks, i+1, res-tasks[i], bitmask^(1<<i))
		h2, b2 := dfs(tasks, i+1, res, bitmask)
		if h2 < h1 {
			return h2, b2
		}
		return h1, b1
	}
	return res, bitmask
}

func dynamicProgramming(tasks []int, dp map[int]int, sessionTime, bitmask int) int {
	if v, exist := dp[bitmask]; exist {
		return v
	}
	minValue := math.MaxInt32
	for i := range tasks {
		if bitmask&(1<<i) == 0 {
			continue
		}
		_, nextBitmask := dfs(tasks, i+1, sessionTime-tasks[i], bitmask^(1<<i))
		nextValue := dynamicProgramming(tasks, dp, sessionTime, nextBitmask)
		if nextValue+1 < minValue {
			minValue = nextValue + 1
		}
	}
	dp[bitmask] = minValue
	return dp[bitmask]
}
