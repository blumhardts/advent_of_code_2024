package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lefts, rights := make([]int, 0), make([]int, 0)

	// Read file
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())

		// Push IDs into lefts and rights as integers
		left, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	// Sort lefts and rights
	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})

	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	// Calculate total distance
	total_distance := 0
	for i := 0; i < len(lefts); i++ {
		distance := lefts[i] - rights[i]
		if distance < 0 {
			distance = -distance
		}

		total_distance += distance
	}

	fmt.Printf("Total distance: %d\n", total_distance)

	// Compile number of occurrences in rights
	occurrences := make(map[int]int)
	for _, id := range rights {
		occurrences[id] += 1
	}

	// Calculate total similarity
	total_similarity := 0
	for _, id := range lefts {
		total_similarity += id * occurrences[id]
	}

	fmt.Printf("Total similarity: %d\n", total_similarity)
}
