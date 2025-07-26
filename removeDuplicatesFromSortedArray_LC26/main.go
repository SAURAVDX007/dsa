package main

import "fmt"

func removeDuplicates(nums []int) int {
    pos := 1
    for i:=1; i<len(nums); i++{
        if nums[i] != nums[pos-1]{
            nums[pos] = nums[i]
            pos++
        }
    }
    return pos
}

func main(){
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	fmt.Println(removeDuplicates([]int{0,0,0,0,0}))
}