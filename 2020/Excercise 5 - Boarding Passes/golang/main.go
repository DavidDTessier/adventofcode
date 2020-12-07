package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type BoardingPass struct {
	PassNumber string
	Row        int64
	Column     int64
}

func (b BoardingPass) GetSeatId() int64 {
	return b.Row*8 + b.Column
}

func decodeSeatRowCol(pass string) (int64, int64) {
	rr, _ := regexp.Compile("[BR]")
	cr, _ := regexp.Compile("[FL]")
	runes := []rune(pass)
	var strRow = string(runes[0:7])
	var strCol = string(runes[7:10])

	var sRow = string(rr.ReplaceAllString(strRow, "1"))
	sRow = string(cr.ReplaceAllString(sRow, "0"))
	var sCol = string(rr.ReplaceAllString(strCol, "1"))
	sCol = string(cr.ReplaceAllString(sCol, "0"))

	var row, _ = strconv.ParseInt(sRow, 2, 0)
	var col, _ = strconv.ParseInt(sCol, 2, 0)

	return row, col
}

func LoadBoardingPasses(path string) (bps []BoardingPass) {

	lines := readLines(path)
	for _, l := range lines {
		text := strings.Trim(l, "\n\r")

		row, col := decodeSeatRowCol(text)
		b := BoardingPass{
			PassNumber: text,
			Row:        row,
			Column:     col,
		}

		bps = append(bps, b)

	}

	return bps
}

func readLines(path string) (lines []string) {
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

		lines = append(lines, line)

		if err == io.EOF {
			break
		}
	}

	return lines

}

func main() {
	os.Chdir("../input")
	curPath, err := os.Getwd()

	if err != nil {
		log.Fatal("main: %s", err)
	}

	bps := LoadBoardingPasses(curPath + "/input.dat")

	var seatIds []int
	for _, b := range bps {
		var m int64 = b.GetSeatId() << 0
		n := int(m)
		seatIds = append(seatIds, n)
	}

	sort.Ints(seatIds)
	highest := seatIds[len(seatIds)-1]

	fmt.Println("Total Valid Passports %d", highest)
}
