package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"io"
	"strings"
)

const tree = "#"

type Slope struct {
	moveRight int
	moveDown  int
	treeCount int
}

var slopes = []Slope {
	{
		moveRight: 1,
		moveDown: 1,
	},
	{
		moveRight: 3,
		moveDown: 1,
	},
	{
		moveRight: 5,
		moveDown: 1,
	},
	{
		moveRight: 7,
		moveDown: 1,
	},
	{
		moveRight: 1,
		moveDown: 2,
	},
}


func readLines(path string) (treeMap [][]string) {
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

		line = strings.Trim(line, "\n\r")

		if line != "" {
			lineSlice := make([] string, len(line))
			for i, char := range []rune(line) {
				lineSlice[i] = string(char)
			}
			treeMap = append(treeMap, lineSlice)
		}

		if err == io.EOF {
			break
		}
	}
	
	return treeMap

}


func DoWork(treeMap [][]string) {
	
	totalSlopeTreeCount := 1
	/* output each array element's value */
	for _, slope := range slopes {
		var curY = 0
		var curX = 0

		for {

			if treeMap[curY][curX] == tree {
				slope.treeCount++
			}

			curX += slope.moveRight
			curY += slope.moveDown

			if (curY >= len(treeMap)) {
				break
			}

			if (curX >= len(treeMap[curY])) {
				curX -= len(treeMap[curY])
			}
		}

		fmt.Println("Slope tree Count : %d", slope.treeCount)
		totalSlopeTreeCount *= slope.treeCount
	
	}
	
	fmt.Println("Total Slope tree Count : %d", totalSlopeTreeCount)
	
}


func main() {
	os.Chdir("../input")
	curPath, err := os.Getwd()
	
	if err != nil {
		log.Fatal("main: %s", err)
	}

	treeMap := readLines(curPath + "\\day3input.dat")

	if err != nil {
		log.Fatal("readLines: %s", err)
	}

	DoWork(treeMap)

}