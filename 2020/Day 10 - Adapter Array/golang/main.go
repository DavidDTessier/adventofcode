package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var bagAdapters []int

func main() {
	bagAdapters = loadData()
	max := getMaxAdapterJoltage()
	bagAdapters = append(bagAdapters, max)
	bagAdapters = append(bagAdapters, 0)
	sort.Ints(bagAdapters)
	DoPart1()
	DoPart2()
}

func getMaxAdapterJoltage() (max int) {
	adapters := bagAdapters
	sort.Ints(adapters)
	foundMax := adapters[len(adapters)-1]
	max = foundMax + 3
	return max
}

func DoPart1() {

	diff1Jolt := 0
	diff3Jolt := 0

	for idx, val := range bagAdapters {

		nextIdx := idx + 1
		if nextIdx == len(bagAdapters) {
			break
		}
		var nextVal = bagAdapters[nextIdx]
		var diff = diff(val, nextVal)
		if diff == 1 {
			diff1Jolt++
		} else if diff == 3 {
			diff3Jolt++
		}
	}

	fmt.Println("1-jolt differences %i", diff1Jolt)
	fmt.Println("3-jolt differences %i", diff3Jolt)
	fmt.Println("The number of 1-jolt differences multiplied by the number of 3-jolt differences : %i", diff1Jolt*diff3Jolt)
}

func DoPart2() {
	pathCounts := make([]int, len(bagAdapters))
	pathCounts[0] = 1
	fmt.Println(pathCounts)
	backchecks := []int{1, 2, 3, 4}
	fmt.Println(backchecks)
	for idx := 1; idx < len(bagAdapters); idx++ {

		val := bagAdapters[idx]
		for idx1 := 1; idx1 < len(backchecks); idx1++ {
			checkIdx := idx - idx1
			if checkIdx < 0 || bagAdapters[checkIdx] < val-3 {
				continue
			}
			pathCounts[idx] += pathCounts[checkIdx]
		}
	}

	fmt.Println(pathCounts)
	fmt.Println("Number of potential paths is :", pathCounts[len(pathCounts)-1])
}

func diff(val1 int, val2 int) int {
	if val1 < val2 {
		return val2 - val1
	}

	return val1 - val2
}

func loadData() (output []int) {
	os.Chdir("../input")
	curPath, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}
	filePath := curPath + string(os.PathSeparator) + "input.dat"

	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Cannot open file: %v", err))
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("Cannot read line: %v", err))
		}

		l := strings.Trim(line, "\n\r")

		i, _ := strconv.Atoi(l)
		output = append(output, i)

		if err == io.EOF {
			break
		}
	}

	sort.Ints(output)
	return output
}
