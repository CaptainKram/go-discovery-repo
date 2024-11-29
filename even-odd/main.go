package main

import "fmt"

func main() {
	var numbers [10]int
	for i := range 10 {
		numbers[i] = i
		fmt.Printf("%v ", numbers[i])
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
