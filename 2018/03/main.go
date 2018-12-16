package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	ID         string
	Offset     offset
	Dimensions dimensions
}

type offset struct {
	X int
	Y int
}

type dimensions struct {
	Width  int
	Height int
}

func main() {
	fabric := [1000][1000]int{}
	claims, err := parseClaimsFromStdIn()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range claims {
		for x := c.Offset.X; x < c.Offset.X+c.Dimensions.Width; x++ {
			for y := c.Offset.Y; y < c.Offset.Y+c.Dimensions.Height; y++ {
				fabric[x][y]++
			}
		}
	}

	overlap := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if fabric[x][y] > 1 {
				overlap++
			}
		}
	}

	fmt.Println(overlap)
}

func parseClaimsFromStdIn() ([]claim, error) {
	claims := []claim{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		claim := parseClaim(scanner.Text())
		claims = append(claims, claim)
	}
	return claims, scanner.Err()
}

func parseClaim(raw string) claim {
	segments := strings.Split(raw, " ")
	parsed := claim{
		ID:         parseID(segments[0]),
		Offset:     parseOffset(segments[2]),
		Dimensions: parseDimensions(segments[3]),
	}
	return parsed
}

func parseID(raw string) string {
	return strings.TrimPrefix(raw, "#")
}

func parseOffset(raw string) offset {
	segments := strings.Split(strings.TrimRight(raw, ":"), ",")
	x, _ := strconv.Atoi(segments[0])
	y, _ := strconv.Atoi(segments[1])
	return offset{x, y}
}

func parseDimensions(raw string) dimensions {
	segments := strings.Split(raw, "x")
	width, _ := strconv.Atoi(segments[0])
	height, _ := strconv.Atoi(segments[1])
	return dimensions{width, height}
}
