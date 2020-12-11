package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()

}

func CheckSum2020(one int, two int, three int) (bool) {
	return one + two + three == 2020
}

func Part1(numbers []string) {
	for i, line := range numbers {
		for j, line1 := range numbers {
			if line != line1 {
				one, err := strconv.Atoi(line)
				two, err1 := strconv.Atoi(line1)

				if err != nil {
					log.Fatal("main: %s", err)
				}
				if err1 != nil {
					log.Fatal("main: %s", err1)
				}

				if CheckSum2020(one, two, 0) {
					fmt.Println(i, line)
					fmt.Println(j, line1)
					fmt.Printf("\nNumber %d + %d == 2020", one, two)
					fmt.Printf("\nNumber %d x %d == %d", one, two, (one * two))
				}
			}
			
		}
		
	}
}

func Part2(numbers []string) {
	for i, line := range numbers {
		for j, line1 := range numbers {
			for x, line2 := range numbers {
					one, err := strconv.Atoi(line)
					two, err1 := strconv.Atoi(line1)
					three, err2 := strconv.Atoi(line2)

					if err != nil {
						log.Fatal("main: %s", err)
					}
					if err1 != nil {
						log.Fatal("main: %s", err1)
					}

					if err2 != nil {
						log.Fatal("main: %s", err1)
					}
	
					if CheckSum2020(one, two, three) {
						fmt.Println(i, line)
						fmt.Println(j, line1)
						fmt.Println(x, line2)
						fmt.Printf("\nNumber %d + %d + %d == 2020", one, two, three)
						fmt.Printf("\nNumber %d x %d x %d == %d", one, two, three, (one * two * three))
					}
				}
			}
		}
}

func main() {
	curPath, err := os.Getwd()
	lines, err := readLines(curPath + "\\input.txt")

	if err != nil {
		log.Fatal("readLines: %s", err)
	}

	fmt.Println("Part 1: ----")
	Part1(lines)
	fmt.Println("\nEnd Part 1: ----")
	fmt.Println("\nPart 2: ----")
	Part2(lines)
	fmt.Println("\nEnd Part 2: ----")
}
