package main

import "fmt"

func findDuplicates(nums []int) []int {
	lenNums := len(nums)

	for i :=0; i<lenNums; i++ {
		index := (nums[i] - 1) % lenNums
		nums[index] = lenNums + nums[index]
	}

	k := 0

	for i := 0; i < lenNums; i++ {
		if nums[i]> 2*lenNums {
			nums[k] = i + 1
			k++
		}
	}

	return nums[0:k]
}

func main() {
	output := findDuplicates([]int{1,2,3,2,4,3,5,5})
	fmt.Println(output)
}