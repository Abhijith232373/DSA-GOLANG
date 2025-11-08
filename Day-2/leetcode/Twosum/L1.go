package main

import (
	"fmt"

)

func sum(nums[]int ,target int)bool{
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if nums[i]+nums[j]==target{
			return true
		}
		}
	}
	return false
}
func main() {
	n:=sum([]int{1,2,3,4,5},9)
	fmt.Println(n)
}