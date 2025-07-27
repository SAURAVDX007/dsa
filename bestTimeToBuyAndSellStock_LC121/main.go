package main

import "fmt"

func maxProfit(prices []int) int {
    maxProfit := 0
    buyPrice := prices[0]
    for i:=1; i<len(prices); i++{
        if prices[i] < buyPrice{
            buyPrice = prices[i]
        }else{
            maxProfit = max(maxProfit, prices[i]-buyPrice)
        }
    }
    return maxProfit
}

func main(){
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,5,3,1}))
}