package main

import "fmt"

func rotate(nums []int, k int) {
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))

}
func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(n, 3)
	fmt.Println(n)
}