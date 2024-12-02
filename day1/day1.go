package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	filename := "input.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	// Split content into lines
	lines := strings.Split(string(content), "\n")

	numbers1, numbers2 := []int{}, []int{}

	// Process each number
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}
		// Split line into two numbers
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		// Convert strings to numbers
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error converting first number: %v\n", err)
			continue
		}
		numbers1 = append(numbers1, num1)

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error converting second number: %v\n", err)
			continue
		}
		numbers2 = append(numbers2, num2)

	}
	// Sort the numbers
	sort.Ints(numbers1)
	sort.Ints(numbers2)

	// fmt.Println(numbers1)
	// fmt.Println(numbers2)

	if len(numbers1) != len(numbers2) {
		fmt.Println("Lengths do not match")
		return
	}

	final := 0
	for i := 0; i < len(numbers1); i++ {
		final += abs(numbers1[i] - numbers2[i])
	}
	fmt.Printf("Part 1: Final sum: %d\n", final)
}

func part2() {
	filename := "input.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	// Split content into lines
	lines := strings.Split(string(content), "\n")

	numbers1, numbers2 := []int{}, []int{}

	// Process each number
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}
		// Split line into two numbers
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		// Convert strings to numbers
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error converting first number: %v\n", err)
			continue
		}
		numbers1 = append(numbers1, num1)

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error converting second number: %v\n", err)
			continue
		}
		numbers2 = append(numbers2, num2)

	}
	if len(numbers1) != len(numbers2) {
		fmt.Println("Lengths do not match")
		return
	}

	final := 0
	for _, num := range numbers1 {

		count := 0
		for _, target := range numbers2 {
			if num == target {
				count++
			}
		}

		final += num * count
	}

	fmt.Printf("Part 2: Final sum: %d\n", final)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
