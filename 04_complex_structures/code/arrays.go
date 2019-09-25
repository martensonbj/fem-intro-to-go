package main

import "fmt"

func main() {

	scores := [5]float64{1, 2, 3, 4, 5}
	for _, value := range scores {
		fmt.Println(value)
	}
}
