package main

import "fmt"

func moveZeroes(nums []int)  {
    pos := 0
    for i, num := range nums{
        if num != 0{
            nums[i], nums[pos] = nums[pos], nums[i]
            pos++
        }
    }
	fmt.Println(nums)
}

func main(){
	moveZeroes([]int{0,1,0,3,12})
	moveZeroes([]int{0})
	moveZeroes([]int{1,2,3,1,2})
}