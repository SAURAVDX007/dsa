package main

import "fmt"

func maxProfit(prices []int) int {
    ans := 0
    p1, p2 := 0, 1
    for p2 < len(prices){
        if prices[p2] - prices[p1] > 0{
            ans += prices[p2] - prices[p1]
        }
        p1++
        p2++
    }
    return ans
}

func main(){
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{7,1,5,13,3,6,4,19}))
}