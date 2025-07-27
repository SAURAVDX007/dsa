package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
	rotate([]int{-1}, 4)
}
