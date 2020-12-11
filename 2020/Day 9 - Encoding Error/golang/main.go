package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	os.Chdir("../input")
	curPath, err := os.Getwd()

	if err != nil {
		log.Fatal("main: %s", err)
	}

	file := curPath + string(os.PathSeparator) + "input.dat"
	numbers := readLines(file)
	res := DoPart1(numbers)
	fmt.Println("First Invalid Number is: %i", res)
	weakness := DoPart2(numbers, res)
	fmt.Println("Weakness %i:", weakness)


}

func DoPart1(numbers []int) (int) {
	var preambleSize = 25
	var result = 0
	for idx := preambleSize; idx < len(numbers); idx++ {
		var preambleSet = numbers[idx-preambleSize:idx]
		var currentNumber = numbers[idx]
		if FindInvalidNumber(preambleSet, currentNumber) == false {
			result = currentNumber
			break
		}

	}

	return result
}

func DoPart2(numbers []int, targetNumber int) (int) {
	var result = 0
	for idx := 0; idx < len(numbers); idx++ {
		if numbers[idx] == targetNumber {
			continue
		}

		var sum = 0
		var nextIdx = idx
		for sum < targetNumber {
			sum += numbers[nextIdx]
			nextIdx++
		} 

		if sum == targetNumber {
			sets := numbers[idx:nextIdx]
			sort.Ints(sets)
			min := sets[0]
			max := sets[len(sets)-1]
			fmt.Println("Weakness at idx %i and idx %i", min, max)
			result = min+max

			break
		}
	}

	return result
}

func FindInvalidNumber(numbers []int, value int) (bool) {
	var sum []int
	for _, i := range numbers {
		for _, j := range numbers {
			if i != j {
				sum = append(sum, i+j);
			}
		}
	}

	return Contains(sum, value) 
}


func Contains(array []int, value int) (bool) {
	for _, item := range array {
		if (item == value) {
			return true
		}
	}

	return false
}

func readLines(path string) (output []int) {
	file, err := os.Open(path)
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
		
		i,_ := strconv.Atoi(l);
		output = append(output,  i)

		if err == io.EOF {
			break
		}
	}

	
	return output

}