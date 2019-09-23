package main

import (
	"fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := utils.Add(2, 2)
	return totalValue
}

func main() {
	total := calculateImportantData()
	fmt.Println(total)
}
