package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func GetLowHigh(value string) (int, int) {
	st := strings.Split(value, "-")
	low := st[0]
	high := st[1]

	one, err := strconv.Atoi(low)
	two, err1 := strconv.Atoi(high)

	if err != nil {
		log.Fatal("main: %s", err)
	}
	if err1 != nil {
		log.Fatal("main: %s", err1)
	}

	return one, two
}

func ValidPassword(low int, high int, value string, matchChar string) (bool) {
	lowerStr := strings.ToLower(value)
	char := strings.ToLower(matchChar)

	charCount := strings.Count(lowerStr, char)

	if charCount >= low && charCount <= high {
		return true
	} 
	
	return false
}

func isValidCharAt(s string, idx int, matchchar string) (bool){
	zeroIdx := idx - 1
	
	strs := strings.Split(s, "")

	return (zeroIdx >= 0 && zeroIdx <= (len(strs) - 1) && strs[zeroIdx] == matchchar)


}

func ValidPasswordIndex(firstIdx int, lastIdx int, value string, matchChar string) (bool) {
	lowerStr := strings.ToLower(value)
	char := strings.ToLower(matchChar)
	
	isFirstIdxValid := isValidCharAt(lowerStr, firstIdx, char)
	isLastIdxValid := isValidCharAt(lowerStr, lastIdx, char)

	if isFirstIdxValid != isLastIdxValid {
		return true
	} 
	
	return false
}


func DoWork(lines []string) {
	var validPasswords int
	var validPasswordIdxs int

	for i, line := range lines {
		fmt.Println(i, line)
		s := strings.Split(line, " ")

		low, high := GetLowHigh(s[0])
		val := strings.Trim(s[1], ":")
		pass := s[2]

		valid := ValidPassword(low, high, pass, val)
		validIdx := ValidPasswordIndex(low, high, pass, val)

		if valid == true {
			validPasswords += 1
		}

		if validIdx == true {
			validPasswordIdxs += 1
		}
	}

	fmt.Println(validPasswords)
	fmt.Println(validPasswordIdxs)
}

func main() {
	os.Chdir("../input")
	curPath, err := os.Getwd()
	
	if err != nil {
		log.Fatal("main: %s", err)
	}

	lines, err := readLines(curPath + "\\day2input.txt")

	if err != nil {
		log.Fatal("readLines: %s", err)
	}

	fmt.Println("Part 1: ----")
	DoWork(lines)
	fmt.Println("\nEnd Part 1: ----")

}