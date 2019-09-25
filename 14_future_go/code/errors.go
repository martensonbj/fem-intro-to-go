package main

import (
	"fmt"
	"os"
)

func openFile() error {
	f, err := os.Open("missingFile.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {
	err := openFile()

	if err != nil {
		fmt.Println(err)
	}
}
