package main

import (
	"fmt"
	"math"
)


func majorityElement(nums []int) int {
    // first trying a Hash map based solution. Will optimise later
	// This is linear in time and space solution
	hm := map[int]int{}
	n := len(nums)
	for _, num := range nums{
		hm[num]++
		if hm[num] > int(math.Floor(float64(n)/2.0)){
			return num
		}
	}
	return -1
}

func majorityElementOptimised(nums []int) int{
	// linear in time, constant in space
	majorityElem := nums[0]
	count := 1
	for i:=1; i<len(nums); i++{
		if count == 0{
			majorityElem = nums[i]
		}
		if nums[i] == majorityElem{
			count++
		}else{
			count--
		}
	}
	return  majorityElem
}

func main(){
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
	fmt.Println(majorityElement([]int{66}))

	fmt.Println("===========")

	fmt.Println(majorityElementOptimised([]int{3,2,3}))
	fmt.Println(majorityElementOptimised([]int{2,2,1,1,1,2,2}))
	fmt.Println(majorityElementOptimised([]int{66}))
}