package main

import "fmt"

func main() {

	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range num {
		if value%2 == 0 {
			fmt.Println(value, " is even")
		} else {
			fmt.Println(value, " is odd")
		}

	}
}
