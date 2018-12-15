package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var twoLetterCount int
var threeLetterCount int

func main() {
	boxes, err := parseBoxIdsFromStdIn()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range boxes {
		for _, b := range boxes {
			matches := findMatchingCharacters(a, b)
			difference := len(a) - len(matches)

			if difference == 1 {
				fmt.Println(strings.Join(matches, ""))
				os.Exit(0)
			}
		}
	}

	fmt.Println("Failed to find to packages with difference of one.")
}

func parseBoxIdsFromStdIn() ([]string, error) {
	boxes := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		boxes = append(boxes, scanner.Text())
	}
	return boxes, scanner.Err()
}

func findMatchingCharacters(a, b string) []string {
	matches := []string{}
	for i := range a {
		if a[i] == b[i] {
			matches = append(matches, string(a[i]))
		}
	}
	return matches
}

func calculateDifference(a, b string) int {
	difference := 0
	for i := range a {
		if a[i] != b[i] {
			difference++
		}
	}
	return difference
}
