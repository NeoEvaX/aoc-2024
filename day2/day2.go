package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := splitData()
	part1(lines)
	part2(lines)
}

func splitData() []string {
	filename := "input"
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	// Split content into lines
	lines := strings.Split(string(content), "\n")

	return lines
}

func part1(lines []string) {
	final := 0

	// Process each number
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}
		// Split line into two numbers
		nums := strings.Split(line, " ")

		numbers := []int{}
		for _, num := range nums {

			convertedNumber, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting number: %v\n", err)
				continue
			}

			numbers = append(numbers, convertedNumber)
		}
		ascending, descending, correct := false, false, true
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] && descending == false {
				descending = true
			}

			if numbers[i-1] < numbers[i] && ascending == false {
				ascending = true
			}

			difference := abs(numbers[i-1] - numbers[i])
			if (difference < 1 || difference > 3) && correct {
				correct = false
			}

		}
		if (ascending != descending) && correct {
			// fmt.Println(numbers)
			final += 1
		}
	}

	fmt.Printf("Part 1: Final sum: %d\n", final)
}

func part2(lines []string) {
	final := 0

	// Process each number
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}
		// Split line into two numbers
		nums := strings.Split(line, " ")

		numbers := []int{}
		for _, num := range nums {

			convertedNumber, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting number: %v\n", err)
				continue
			}

			numbers = append(numbers, convertedNumber)
		}
		ascending, descending, bad := 0, 0, 0
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				if bad == 0 && ascending == 0 {
					descending += 1
				} else if bad == 0 && ascending == 1 {
					bad += 1
					descending += 1
				} else if bad == 1 && ascending >= 1 {
					bad += 1
					descending += 1
				}
			}

			if numbers[i-1] < numbers[i] {
				if bad == 0 && descending == 0 {
					ascending += 1
				} else if bad == 0 && descending == 1 {
					bad += 1
					ascending += 1
				} else if bad == 1 && descending >= 1 {
					bad += 1
					ascending += 1
				}
			}

			difference := abs(numbers[i-1] - numbers[i])
			if difference < 1 || difference > 3 {
				bad += 1
			}

		}

		if bad == 1 {
			if (ascending >= len(numbers)-2) || (descending >= len(numbers)-2) {
				fmt.Printf("ascending: %d, desc: %d, bad: %d\n", ascending, descending, bad)
				fmt.Printf("good: %d\n", numbers)
				final += 1
			} else {
				fmt.Printf("ascending: %d, desc: %d, bad: %d\n", ascending, descending, bad)
				fmt.Printf("bad: %d\n", numbers)
			}
		} else if bad == 0 {
			if (ascending >= len(numbers)-2) || (descending >= len(numbers)-2) {
				fmt.Printf("ascending: %d, desc: %d, bad: %d\n", ascending, descending, bad)
				fmt.Printf("good: %d\n", numbers)
				final += 1

			} else {
				fmt.Printf("ascending: %d, desc: %d, bad: %d\n", ascending, descending, bad)
				fmt.Printf("bad: %d\n", numbers)
			}
		} else {
			fmt.Printf("2+ Bad: ascending: %d, desc: %d, bad: %d\n", ascending, descending, bad)
		}
	}

	fmt.Printf("Part 2: Final sum: %d\n", final)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
