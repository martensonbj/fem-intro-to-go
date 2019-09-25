package main

import "fmt"

func printAge(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
}

func main() {
	printAge(16, 21, 30)
}
