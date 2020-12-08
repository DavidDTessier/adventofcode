package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
)

type Group struct {
	Answer string
	Size int
}

func (g Group) GetUnanimousAnswerCount() (totalUnanomousAnswers int) {
	seen := make(map[string]int) // a Go set 
	var uniq []int
	for _, ch := range g.Answer {
		str := string(ch)
		cnt := strings.Count(g.Answer, str) 
		if (cnt > 1) {
			if _, ok := seen[str]; !ok { 
				seen[str] = cnt 
			}
		}	
	} 

	for _, i := range seen {
		if i == g.Size {
			uniq = append(uniq, i)
		}
	}

	return len(uniq)
}

func stripDuplicateAnswers(input string) (string) {
	output := strings.Builder{} 
 
	seen := make(map[rune]struct{}) // a Go set 	
 
	//build output 
	for _, ch := range input { 
		if _, ok := seen[ch]; !ok { 
			seen[ch] = struct{}{} 
			output.WriteRune(ch) 
		} 	
	} 
 
	return output.String()
}

func readLines(path string) (groups []Group) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot open file: %v", err))
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	output := strings.Builder{} 
	var groupSize = 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("Cannot read line: %v", err))
		}

		l := strings.Trim(line, "\n\r")
		if (l != "") {
			output.WriteString(l)
			groupSize++
		} else {
			g := Group {
				Answer : output.String(),
				Size : groupSize,
			}
			groups = append(groups, g)
			groupSize = 0
			output.Reset()
		}

		if err == io.EOF {
			break
		}
	}

	if (output.String() != "") {
		g := Group {
			Answer : output.String(),
			Size: groupSize,
		}
		groups = append(groups, g)
		output.Reset()
	}
	
	
	return groups

}

func main()  {
	os.Chdir("../input")
	curPath, err := os.Getwd()
	
	if err != nil {
		log.Fatal("main: %s", err)
	}

	groups := readLines(curPath + "\\input.dat")
	fmt.Println("Part 1: -----\n\r")
	var totalgroupsAnswersCount = 0
	for _, g := range groups {
		uniqueAs := stripDuplicateAnswers(g.Answer)
		totalgroupsAnswersCount += len(uniqueAs)
	}

	fmt.Println("Total Sum Of Group Answers : %i", totalgroupsAnswersCount)
	fmt.Println("\n\rPart 2: -----\n\r")
	var totalUniqueAnswers = 0
	for _, g := range groups {

		totalUniqueAnswers += g.GetUnanimousAnswerCount()
	}

	fmt.Println("Total Unique Of Group Answers : %i", totalUniqueAnswers)
}
