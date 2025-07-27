package main

import "fmt"

func productExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))
    product := 1
    countOfZero := 0
    for _, num := range nums{
        if num == 0{
            countOfZero++
        }else{
            product *= num
        }    
    }
    if countOfZero > 1{
        return ans
    }
    if countOfZero == 1{
        for i, num := range nums{
            if num == 0{
                ans[i] = product
            }else{
                ans[i] = 0
            }
        }
        return ans
    }
    for i, num := range nums{
        ans[i] = product/num
    }
    return ans
}


func productExceptSelfOptimised(nums []int) []int {
    ans := make([]int, len(nums))
    prefix := make([]int, len(nums))
    suffix := make([]int, len(nums))
    prefix[0] = 1
    suffix[len(nums)-1] = 1
    for i:=1; i<len(nums); i++{
        prefix[i] = prefix[i-1] * nums[i-1]
    }
    for i:=len(nums)-2; i>=0; i--{
        suffix[i] = suffix[i+1]*nums[i+1]
    }
    for i:=0; i<len(nums); i++{
        ans[i] = prefix[i] * suffix[i]
    }
    return ans
}

func productExceptSelfBestSolution(nums []int) []int {
    ans := make([]int, len(nums))
    ans[0] = 1
    for i:=1; i<len(nums); i++{
        ans[i] = ans[i-1] * nums[i-1]
    }
    suffixProd := 1
    for i:=len(nums)-1; i>=0; i--{
        ans[i] = ans[i] * suffixProd
        suffixProd *= nums[i]
    }
    return ans
}

func main(){
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3,0}))

	fmt.Println(productExceptSelfOptimised([]int{1,2,3,4}))
	fmt.Println(productExceptSelfOptimised([]int{-1,1,0,-3,3}))
	fmt.Println(productExceptSelfOptimised([]int{-1,1,0,-3,3,0}))

	fmt.Println(productExceptSelfBestSolution([]int{1,2,3,4}))
	fmt.Println(productExceptSelfBestSolution([]int{-1,1,0,-3,3}))
	fmt.Println(productExceptSelfBestSolution([]int{-1,1,0,-3,3,0}))
}