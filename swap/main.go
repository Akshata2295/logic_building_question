package main

import "fmt"

func swap() [] int {
	a, b := 10, 12
	b, a = a, b
	return []int{a,b}
}

func main() {
	fmt.Println(swap())
}