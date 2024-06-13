package Util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv() {
	fmt.Println("Loading environment variables")
	loadEnvFromFile(".env")
}

func loadEnvFromFile(filename string) error {
	// Open the .env file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // skip invalid lines
		}
		key, value := parts[0], parts[1]
		// Set the environment variable if it's not already set
		if os.Getenv(key) == "" {
			err := os.Setenv(key, value)
			if err != nil {
				return fmt.Errorf("failed to set environment variable %s: %w", key, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	return nil
}
