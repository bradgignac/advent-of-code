package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/tools/container/intsets"
)

func main() {
	var frequency int
	var previous intsets.Sparse

	previous.Insert(frequency)

	inputs, err := parseAdjustmentsFromStdin()
	if err != nil {
		log.Fatal(err)
	}

	for true {
		for _, adjustment := range inputs {
			frequency += adjustment
			isUnique := previous.Insert(frequency)

			if !isUnique {
				fmt.Println(frequency)
				os.Exit(0)
			}
		}
	}
}

func parseAdjustmentsFromStdin() ([]int, error) {
	inputs := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, line)
	}

	return inputs, scanner.Err()
}
