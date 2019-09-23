package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("something bad happened")
	}
	return nil
}

func openFile() error {
	f, err := os.Open("missingFile.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {
	// err := isGreaterThanTen(9)
	// if err != nil {
	// fmt.Errorf("%d is NOT GREATER THAN TEN", num)
	// 	// panic(err)
	// 	// log.Fatalln(err)
	// }

	err := openFile()

	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		// or more descriptively:
		// fmt.Println(fmt.Errorf("Error opening file: %+v", err))
	}
}
