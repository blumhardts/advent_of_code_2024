package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strsToInts(strs []string) []int {
	ints := make([]int, len(strs))

	for i, str := range strs {
		int, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ints[i] = int
	}

	return ints
}

func diff(minuend, subtrahend int) int {
	difference := minuend - subtrahend
	if difference < 0 {
		return -difference
	}
	return difference
}

func unsafe(levels []int, default_to_unsafe bool) bool {
	var sort string
	if levels[1] > levels[0] {
		sort = "increasing"
	} else if levels[1] < levels[0] {
		sort = "decreasing"
	} else {
		return default_to_unsafe || unsafe(levels[1:], true) || unsafe(append(levels[:1], levels[2:]...), true)
	}

	difference := diff(levels[0], levels[1])
	if difference < 1 || 3 < difference {
		return default_to_unsafe || unsafe(levels[1:], true) || unsafe(append(levels[:1], levels[2:]...), true)
	}

	for i, level := range levels[2:] {
		prev := levels[i+1]

		if sort == "increasing" && level < prev {
			return default_to_unsafe || unsafe(append(levels[:i+2], levels[i+2+1:]...), true)
		}

		if sort == "decreasing" && level > prev {
			return default_to_unsafe || unsafe(append(levels[:i+2], levels[i+2+1:]...), true)
		}

		difference := diff(prev, level)
		if difference < 1 || 3 < difference {
			return default_to_unsafe || unsafe(append(levels[:i+2], levels[i+2+1:]...), true)
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	number_of_safe_reports := 0
	line_number := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strsToInts(strings.Fields(scanner.Text()))

		if unsafe(levels, false) {
			continue
		}

		number_of_safe_reports += 1
		line_number += 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d reports are safe\n", number_of_safe_reports)
}
