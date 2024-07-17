package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("problem 2033")
	fmt.Println(minOperations([][]int{
		{931, 128},
		{639, 712},
	}, 73))
}

func minOperations(grid [][]int, x int) int {
	if x <= 0 {
		return -1
	}

	min := grid[0][0]
	mapR, minValue := makeMapping(min, x)
	for _, cells := range grid {
		for _, c := range cells {
			if (c-min)%x != 0 {
				return -1
			}

			for key, value := range mapR {
				mapR[key] = value + int(math.Abs(float64(key-c))/float64(x))
			}
		}
	}

	fmt.Println(mapR)
	minOps := mapR[minValue]
	for _, value := range mapR {
		if value < minOps {
			minOps = value
		}
	}

	return minOps
}

func makeMapping(min int, x int) (map[int]int, int) {
	mapping := map[int]int{}
	for i := 0; i <= 25; i++ {
		mapping[min+x*i] = 0
	}

	return mapping, min + x*0
}
