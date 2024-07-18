package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem 974")
	fmt.Println("result:", subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

func subarraysDivByK(nums []int, k int) int {
	return solution2(nums, k)
}

// time limit exceeded
func solution1(nums []int, k int) int {
	count := 0
	for i, v := range nums {
		sum := v
		if sum%k == 0 {
			count++
		}

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum%k == 0 {
				count++
			}
		}
	}
	return count
}

func solution2(nums []int, k int) int {
	remainders := make(map[int]int)
	remainders[0] = 1 // Initialize with 0 remainder to handle subarrays starting from index 0

	currentSum := 0
	cnt := 0

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i] // Update the cumulative sum
		remainder := currentSum % k

		// Adjust the remainder to be positive
		if remainder < 0 {
			remainder += k
		}

		// If the remainder is already in the map, it means we found a valid subarray
		cnt += remainders[remainder]

		// Update the frequency of the current remainder
		remainders[remainder]++
	}

	return cnt

}

// func recursion(start int, end int, nums []int, k int, count int) int {
// 	if start == end {
// 		if nums[start]%k == 0 {
// 			count++
// 		}
// 		return count
// 	}

// 	sum := nums[start]
// 	subCount := 0
// 	count += subRecursion(start+1, end, sum, nums, k, subCount)

// 	return recursion(start+1, end, nums, k, count)
// }

// func subRecursion(start int, end int, sum int, nums []int, k int, count int) int {
// 	if start == end {
// 		if nums[start]%k == 0 {
// 			count++
// 		}

// 		return count
// 	}

// 	sum += nums[start]
// 	return subRecursion(start+1, end, sum, nums, k, count)
// }
