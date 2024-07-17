package main

import "fmt"

func main() {
	fmt.Println(numberOfWeek([]int{5, 9, 4, 4, 8, 9, 9, 8, 7, 3}))
}

func numberOfWeek(milestones []int) int64 {
	if len(milestones) == 0 {
		return 0
	}

	max := milestones[0]
	total := 0
	for _, m := range milestones {
		if m > max {
			max = m
		}
		total += m
	}

	total -= max

	if total == max {
		return int64(total) * 2
	} else if total < max {
		return int64(total)*2 + 1
	}

	return int64(total) + int64(max)
}
