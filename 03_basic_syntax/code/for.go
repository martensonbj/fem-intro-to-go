// Uncomment the entire file

package main

import "fmt"

func main() {

	// 	// ****************************

	// i := 1

	// for i := 1; i <= 100; i++ {
	// 	fmt.println(i)
	// }

	// 	// ****************************

	// i := 1

	// for i <= 100 {
	// 	fmt.println(i)
	// 	// this will behave like a while loop
	// 	i += 1
	// }

	// 	// ****************************

	var mySentence = "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", string(letter))
	}
}
