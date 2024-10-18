package main

import (
	"advent_of_go_2023/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := utils.ReadTxtFile("day_01/input.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(3)
	}

	var total int

	for _, row := range input {
		var first, last string

		for _, char := range row {
			if _, err := strconv.Atoi(char); err == nil {
				if first == "" {
					first = char
				}
				last = char
			}
		}

		code, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(3)
		}

		total += code
	}
	fmt.Println(total)
}
