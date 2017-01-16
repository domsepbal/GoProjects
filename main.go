package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World! First test!")
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printEven(num...)
}

func printEven(number ...int) {
	for _, n := range number {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
