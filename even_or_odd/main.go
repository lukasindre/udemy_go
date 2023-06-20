package main

import "fmt"

func main() {
	var nums []int
	l := 10
	for i := 1; i <= l; i++ {
		nums = append(nums, i)
	}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is even!")
		} else if num%2 == 1 {
			fmt.Println(num, "is odd!")
		} else {
			fmt.Println(num, "is a figment of our imagination!")
		}
	}
}
