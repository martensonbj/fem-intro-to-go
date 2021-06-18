package main

import "fmt"

// func main() {
// 	mySentence := "Hello this is a sentence."

// 	for _, value := range mySentence {
// 		fmt.Println(string(value))
// 	}
// }

func main() {
	sentence := "Hello this is a sentence."

	for index, value := range sentence {
		if index % 2 != 0 {
			fmt.Printf("%s\n", string(value))
		}
	}
}
