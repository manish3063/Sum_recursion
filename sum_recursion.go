package main

//Sum of natural numbers using recursion
// Input : 3
// Output : 6
// Explanation : 1 + 2 + 3 = 6

import "fmt"

func main() {
	a := Display(5)
	fmt.Println(a)

}

func Display(num int) int {
	var sum int = 0
	if num == 1 {
		return num
	} else {
		sum = num + Display(num-1)
		return sum
	}

}
