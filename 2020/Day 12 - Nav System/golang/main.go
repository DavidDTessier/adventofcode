package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const (
	North int = 0
	East      = 90
	South     = 180
	West      = 270
)

type NavigationDirection struct {
	direction string
	position  int
}

type Ship struct {
	direction int
	x_pos     int // east-west
	y_pos     int // nort-south
}

func (s *Ship) MoveEast(position int) {
	s.x_pos += position
}

func (s *Ship) MoveWest(position int) {
	s.x_pos -= position
}

func (s *Ship) MoveNorth(position int) {
	s.y_pos -= position

}

func (s *Ship) MoveSouth(position int) {
	s.y_pos += position
}

func (s *Ship) MoveForward(position int) {
	switch s.direction {
	case East:
		s.MoveEast(position)
		break
	case West:
		s.MoveWest(position)
		break
	case North:
		s.MoveNorth(position)
		break
	case South:
		s.MoveSouth(position)
		break
	}
}

func (s *Ship) ChangeDirection(degrees int) {
	s.direction += degrees
	if s.direction < 0 {
		s.direction += 360
	}

	if s.direction > 359 {
		s.direction -= 360
	}
}

func LoadNavigationCoordinates() (coordinates []NavigationDirection) {
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
			nav := NavigationDirection{
				direction: string(line[0]),
				position:  int(line[1]),
			}
			coordinates = append(coordinates, nav)
		}

		if err == io.EOF {
			break
		}
	}

	return coordinates
}

func main() {
	var coords = LoadNavigationCoordinates()
	fmt.Println(len(coords))
	ship := Ship{East, 0, 0}

	for _, coord := range coords {
		switch coord.direction {
		case "N":
			ship.MoveNorth(coord.position)
			break
		case "S":
			ship.MoveSouth(coord.position)
			break
		case "E":
			ship.MoveEast(coord.position)
			break
		case "W":
			ship.MoveWest(coord.position)
			break
		case "F":
			ship.MoveForward(coord.position)
			break
		case "L":
			ship.ChangeDirection(-coord.position)
			break
		case "R":
			ship.ChangeDirection(coord.position)
			break
		}
	}

	fmt.Println(ship.x_pos)
	fmt.Println(ship.y_pos)
	//fmt.Println(direction)
	var manhattan_distance = int(math.Abs(float64(ship.x_pos))) + int(math.Abs(float64(ship.y_pos)))
	fmt.Println("Ship Manhatten Distance", manhattan_distance)
}
