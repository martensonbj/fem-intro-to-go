package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func openFile() error {
// 	f, err := os.Open("missingFile.txt")
// 	if err != nil {
// 		unwrapped := errors.Unwrap(fmt.Errorf("File is Missing: %w", err))
// 		fmt.Println(unwrapped)
// 		return unwrapped
// 	}
// 	defer f.Close()
// 	return nil
// }

// func main() {
// 	err := openFile()

// 	if err != nil {
// 		if errors.Is(err, os.ErrExist) {
// 			fmt.Println("File doesn't exist")
// 		} else {
// 			fmt.Println("something else happened")
// 		}
// 		fmt.Println(err)
// 	}
// }
