package Util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func HandleOpenCsvFile(csvFileName *string) [][]string {
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s \n", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV")
	}
	return lines
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
