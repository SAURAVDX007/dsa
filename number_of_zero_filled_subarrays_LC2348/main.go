package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	currentZeroCount := 0
	ans := 0
	for _, num := range nums {
		if num == 0 {
			currentZeroCount++
		} else {
			ans += helper(currentZeroCount)
			currentZeroCount = 0
		}
	}
	if currentZeroCount != 0 {
		ans += helper(currentZeroCount)
	}
	return int64(ans)
}

func helper(n int) int {
	return ((n + 1) * n) / 2
}

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println(zeroFilledSubarray([]int{2, 10, 2019}))
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(zeroFilledSubarray([]int{0}))
	fmt.Println(zeroFilledSubarray([]int{-1, -19, 2, 0, 0, 0}))
}
