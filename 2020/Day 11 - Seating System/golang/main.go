package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const occupiedSeat = "#"
const emptySeat = "L"
const floor = "."

func readLines() (lines []string) {
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

		line = strings.Trim(line, "\n\r")
		if line != "" {
			lines = append(lines, line)
		}

		if err == io.EOF {
			break
		}
	}

	return lines
}

func getAdjacentSeats(seatMap [][]string, rowID int, seatID int) (adjacentSeats []string) {
	var rowNum = rowID
	var seatNum = seatID

	//leftSeat
	if seatNum > 0 {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum][seatNum-1])
	}
	//topLeftDiagonal
	if seatNum > 0 && rowNum > 0 {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum-1][seatNum-1])
	}
	//bottomLeftDiagonal
	if rowNum+1 < len(seatMap) && seatNum > 0 {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum+1][seatNum-1])
	}
	//topSeat
	if rowNum > 0 {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum-1][seatNum])
	}
	//bottomSeat
	if rowNum+1 < len(seatMap) {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum+1][seatNum])
	}
	//right
	if seatNum+1 < len(seatMap[rowNum]) {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum][seatNum+1])
	}
	//topRightDiagonal
	if seatNum+1 < len(seatMap[rowNum]) && rowNum > 0 {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum-1][seatNum+1])
	}
	//bottomRightDiagonal
	if rowNum+1 < len(seatMap) && seatNum+1 < len(seatMap[rowNum]) {
		adjacentSeats = append(adjacentSeats, seatMap[rowNum+1][seatNum+1])
	}

	return adjacentSeats
}

func getSeatsInView(seatMap [][]string, rowID int, seatID int) (seatsInView []string) {

	//Seats left
	i := seatID - 1
	for i >= 0 {
		if seatMap[rowID][i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID][i])
			break
		}

		i--
	}

	//Seast topLeftDiagonal
	i = 1
	for rowID-i >= 0 && seatID-i >= 0 {
		if seatMap[rowID-i][seatID-i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID-i][seatID-i])
			break
		}

		i++
	}

	//Seats bottomLeft
	i = 1
	for rowID+i < len(seatMap) && seatID-i >= 0 {
		if seatMap[rowID+i][seatID-i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID+i][seatID-i])
			break
		}

		i++
	}

	//Seats Above
	i = rowID - 1
	for i >= 0 {
		if seatMap[i][seatID] != floor {
			seatsInView = append(seatsInView, seatMap[i][seatID])
			break
		}

		i--
	}

	//seats below
	if rowID+1 < len(seatMap) {
		i = rowID + 1
		for i < len(seatMap) {
			if seatMap[i][seatID] != floor {
				seatsInView = append(seatsInView, seatMap[i][seatID])
				break

			}
			i++
		}
	}

	//Seats right
	i = seatID + 1
	for i < len(seatMap[rowID]) {
		if seatMap[rowID][i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID][i])
			break
		}

		i++
	}

	//Seats topRight
	i = 1
	for seatID+i < len(seatMap[rowID]) && rowID-i >= 0 {
		if seatMap[rowID-i][seatID+i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID-i][seatID+i])
			break
		}

		i++
	}

	//bottomRightDiagonal
	i = 1
	for rowID+i < len(seatMap) && seatID+i < len(seatMap[rowID]) {
		if seatMap[rowID+i][seatID+i] != floor {
			seatsInView = append(seatsInView, seatMap[rowID+i][seatID+i])
			break
		}

		i++
	}

	return seatsInView
}

func generateNewSeatingPlan(seatMap [][]string, seatLimit int) [][]string {
	//var maxOccupancy = 4
	var newSeatingPlan = DeepCopy(seatMap)

	for rdx, row := range seatMap {
		for sdx, seat := range row {
			var seatsCloseBy []string
			if seatLimit == 4 {
				seatsCloseBy = getAdjacentSeats(seatMap, rdx, sdx)
			} else {
				seatsCloseBy = getSeatsInView(seatMap, rdx, sdx)
			}
			if seat == emptySeat && CountSeatType(occupiedSeat, seatsCloseBy) == 0 {
				newSeatingPlan[rdx][sdx] = occupiedSeat
			} else if seat == occupiedSeat && CountSeatType(occupiedSeat, seatsCloseBy) >= seatLimit {
				newSeatingPlan[rdx][sdx] = emptySeat
			}
		}
	}

	return newSeatingPlan

}

func DeepCopy(arr1 [][]string) [][]string {
	newArr := make([][]string, len(arr1))
	for rdx, row := range arr1 {
		var values []string
		for _, val := range row {
			values = append(values, string(val))
		}

		newArr[rdx] = values
	}

	return newArr
}

func DoWork(seatMap [][]string, seatLimit int) {

	var newSeatMap = generateNewSeatingPlan(seatMap, seatLimit)
	fmt.Println(reflect.DeepEqual(seatMap, newSeatMap))
	for reflect.DeepEqual(seatMap, newSeatMap) == false {
		seatMap = DeepCopy(newSeatMap)
		newSeatMap = generateNewSeatingPlan(seatMap, seatLimit)
	}
	var occupiedSeatCount = 0
	for _, row := range seatMap {
		for _, seat := range row {
			if seat == occupiedSeat {
				occupiedSeatCount++
			}
		}
	}

	fmt.Println("Occupied Seat Count: %i", occupiedSeatCount)

}

func CountSeatType(seatType string, seats []string) (count int) {
	count = 0
	for _, seat := range seats {
		if seat == seatType {
			count++
		}
	}
	return count
}

func generateSeatMap(rows []string) [][]string {

	smap := make([][]string, len(rows))
	for idx, row := range rows {
		var seats []string
		for _, val := range row {
			seats = append(seats, string(val))
		}

		smap[idx] = seats
	}

	return smap
}

func main() {
	var rows = readLines()
	var seatMap = generateSeatMap(rows)
	DoWork(seatMap, 4)
	DoWork(seatMap, 5)

}
