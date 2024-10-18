package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadTxtFile(path string) ([][]string, error) {

	// Open the file
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	file, err := os.Open(filepath.Join(workingDir, path))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Initialize a  2D slice to store the characters
	var matrix [][]string

	// Read the file line by line
	for scanner.Scan() {
		var lineArray []string
		line := scanner.Text()

		// Convert each character in the line to a string and add it to the array
		for _, char := range line {
			lineArray = append(lineArray, string(char))
		}

		// Append the array of characters (line) to the 2D slice
		matrix = append(matrix, lineArray)
	}

	return matrix, scanner.Err()
}
