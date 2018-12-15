package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error
	var frequency int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		frequency, err = processInputLine(frequency, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(frequency)
}

func processInputLine(frequency int64, input string) (int64, error) {
	i, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return 0, err
	}

	return frequency + i, nil
}
