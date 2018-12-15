package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var twoLetterCount int
var threeLetterCount int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checksumBoxId(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(twoLetterCount * threeLetterCount)
}

func checksumBoxId(id string) {
	charCount := map[rune]int{}
	for _, r := range id {
		charCount[r] = charCount[r] + 1
	}

	hasExactlyTwo := false
	hasExactlyThree := false
	for _, count := range charCount {
		if count == 2 {
			hasExactlyTwo = true
		} else if count == 3 {
			hasExactlyThree = true
		}
	}

	if hasExactlyTwo {
		twoLetterCount++
	}
	if hasExactlyThree {
		threeLetterCount++
	}
}
