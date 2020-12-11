package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const JMPCMD = "jmp"
const ACCCMD = "acc"
const NOOPCMD = "nop"

type Instruction struct {
	action string
	step int
}

func (i Instruction) IsJumpCommand() (bool) {
	return i.action == JMPCMD
}

func (i Instruction) IsAccumulator() (bool) {
	return i.action == ACCCMD
}

func (i Instruction) IsNoOp() (bool) {
	return i.action == NOOPCMD
}





var instructions []Instruction
var accumulator = 0
var executedInstructions []int

func UpdateAccumulator(value int) {
	accumulator += value
}

func main() {
	os.Chdir("../input")
	curPath, err := os.Getwd()

	if err != nil {
		log.Fatal("main: %s", err)
	}

	file := curPath + string(os.PathSeparator) + "input.dat"
	readLines(file)
	DoPart1()


}

func DoPart1() {
	FindHaltingInstruction(0, executedInstructions)
	fmt.Println("Accumulator count before any step is executed a second time : %i", accumulator)
	DoPart2()
	fmt.Println("Accumulator after program terminates  : %i", accumulator)
}

func DoPart2() {
	executedInstructions = []int{}
	accumulator = 0
	var stepIdx = 0
	for stepIdx != len(instructions) {
		var ins = instructions[stepIdx]
		if ins.IsAccumulator() {
			UpdateAccumulator(ins.step)
			executedInstructions = append(executedInstructions, stepIdx)
			stepIdx++
		} else {
			if ins.IsJumpCommand() {
				SwapInstruction(stepIdx, NOOPCMD, ins.step)		
			} else if ins.IsNoOp() {
				SwapInstruction(stepIdx, JMPCMD, ins.step)
			}
			
			stepsRan := executedInstructions
			halted := FindHaltingInstruction(stepIdx, stepsRan)
			
			if halted {
				instructions[stepIdx] = ins
				executedInstructions = append(executedInstructions, stepIdx)
				if ins.IsJumpCommand() {
					stepIdx += ins.step
				} else if ins.IsNoOp() {
					stepIdx++
				}
			
			} else {
				stepFix := instructions[stepIdx]
				fmt.Println("Replace Action step %s with step %s at index %i fixes the flow", ins.action, stepFix.action, stepIdx)
 				break
			}
		}
	}
}

func SwapInstruction(idx int, newCmd string, newStep int) {
	swagCmd := Instruction {
		action: newCmd,
		step: newStep,
	}

	instructions[idx] = swagCmd
}


func FindHaltingInstruction(stepIdx int, executedInstructions []int) (bool) {
	for stepIdx != len(instructions) {
		var ins = instructions[stepIdx]
		if Contains(executedInstructions, stepIdx) {
			return true
		}

		executedInstructions = append(executedInstructions, stepIdx)

		if ins.IsAccumulator() {
			UpdateAccumulator(ins.step)
			stepIdx++
		} else if ins.IsJumpCommand() {
			stepIdx += ins.step
		} else if ins.IsNoOp() {
			stepIdx++
		}
	}

	return false
}

func Contains(array []int, value int) (bool) {
	for _, item := range array {
		if (item == value) {
			return true
		}
	}

	return false
}


func readLines(path string) {
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
		arr := strings.Split(l, " ")
		i,_ := strconv.Atoi(arr[1]);
		
		ins := Instruction {
			action : arr[0],
			step : i,
		}
		
		instructions = append(instructions, ins)
	
		if err == io.EOF {
			break
		}
	}
}