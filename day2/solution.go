package main

import (
	"fmt"
	"strconv"
	"strings"
)

const INPUT_PATH = "./input"

// func main() {
// 	fmt.Println("Day2")
// 	data, err := os.ReadFile(INPUT_PATH)
// 	if err != nil {
// 		panic(err)
// 	}
// 	lines := (strings.Split(string(data), "\n"))
// 	partOne(lines)
// 	partTwo(lines)
// }

func partOne(lines []string) {
	solution := 0
	for _, line := range lines {
		if len(line) > 1 {
			report := make([]int64, 0)
			for _, s := range strings.Split(line, " ") {
				if n, err := strconv.ParseInt(s, 10, 32); err == nil {
					report = append(report, n)
				}
			}
			isValid := true
			if report[0] < report[len(report)-1] {
				for i := 0; i < len(report)-1; i++ {
					diff := report[i+1] - report[i]
					if report[i+1] <= report[i] {
						isValid = false
						break
					}
					if diff < 1 || diff > 3 {
						isValid = false
						break
					}
				}
			} else if report[0] > report[len(report)-1] {
				for i := 0; i < len(report)-1; i++ {
					diff := report[i] - report[i+1]
					if report[i] <= report[i+1] {
						isValid = false
						break
					}
					if diff < 1 || diff > 3 {
						isValid = false
						break
					}
				}
			} else {
				isValid = false
			}

			if isValid {
				solution++
			}
		}
	}

	fmt.Printf("Part One: %d\n", solution)
}

func partTwo(lines []string) {
	solution := 0
	for _, line := range lines {
		if len(line) > 1 {
			report := make([]int64, 0)
			for _, s := range strings.Split(line, " ") {
				if n, err := strconv.ParseInt(s, 10, 32); err == nil {
					report = append(report, n)
				}
			}
			isValid := false
			i, j := checkValidReport(report)
			if i != -1 && j != -1 {
				l1 := make([]int64, 0, len(report)-1)
				l1 = append(l1, report[:i]...)
				l1 = append(l1, report[i+1:]...)
				x, y := checkValidReport(l1)
				if x == -1 && y == -1 {
					isValid = true
				} else {
					l2 := make([]int64, 0, len(report)-1)
					l2 = append(l2, report[:j]...)
					l2 = append(l2, report[j+1:]...)
					a, b := checkValidReport(l2)
					if a == -1 && b == -1 {
						isValid = true
					}
				}

			} else {
				isValid = true
			}

			if isValid {
				solution++
			}
		}
	}

	fmt.Printf("Part Two: %d\n", solution)
}

func checkValidReport(report []int64) (int, int) {
	if report[0] < report[len(report)-1] {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			if report[i+1] <= report[i] {
				return i, i + 1
			}
			if diff < 1 || diff > 3 {
				return i, i + 1
			}
		}
	} else if report[0] > report[len(report)-1] {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]
			if report[i] <= report[i+1] {
				return i, i + 1
			}
			if diff < 1 || diff > 3 {
				return i, i + 1
			}
		}
	} else if report[0] == report[len(report)-1] {
		return 0, len(report) - 1
	}

	return -1, -1
}
