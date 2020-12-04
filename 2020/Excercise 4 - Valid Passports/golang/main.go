package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
)

type Passport struct {
	BirthYear string
	IssueYear string
	ExpirationYear string
	Height string
	HairColor string
	EyeColor string
	PassportId string
	CountryId string
}

func (p Passport) IsValidPassport() (bool) {
	return p.IsAllFieldsRequired() || p.Is7FieldsRequired()
}
func (p Passport) IsAllFieldsRequired() (bool) {
	return (p.BirthYear != "" &&  p.IssueYear != "" && p.ExpirationYear != "" && p.Height != "" && p.EyeColor != "" && p.PassportId != "" && p.CountryId != "" && p.HairColor != "")
}

func (p Passport) Is7FieldsRequired() (bool) {
	return (p.BirthYear != "" &&  p.IssueYear != "" && p.ExpirationYear != "" && p.Height != "" && p.EyeColor != "" && p.PassportId != ""  && p.HairColor != "")
}

func LoadPassports(path string) ([]Passport) {
	
	lines := readLines(path)
	
	var passports []Passport
	var sb strings.Builder

	for _, l := range lines {
		text := strings.Trim(l, "\n\r")
		if text != "" {	
			sb.WriteString(text)
			sb.WriteString(" ")
		} else {
			var str = sb.String()
			strArray := strings.Split(str," ")

			
			p := Passport {
				BirthYear: Find(strArray, "byr"),
				IssueYear: Find(strArray, "iyr"),
				ExpirationYear: Find(strArray, "eyr"),
				Height: Find(strArray, "hgt"),
				HairColor: Find(strArray, "hcl"),
				EyeColor: Find(strArray, "ecl"),
				PassportId: Find(strArray, "pid"),
				CountryId: Find(strArray, "cid"),
			}
			passports = append(passports, p)
			sb.Reset()
		}
	}

	if sb.String() != "" {
		var str = sb.String()
			strArray := strings.Split(str," ")

			
			p := Passport {
				BirthYear: Find(strArray, "byr"),
				IssueYear: Find(strArray, "iyr"),
				ExpirationYear: Find(strArray, "eyr"),
				Height: Find(strArray, "hgt"),
				HairColor: Find(strArray, "hcl"),
				EyeColor: Find(strArray, "ecl"),
				PassportId: Find(strArray, "pid"),
				CountryId: Find(strArray, "cid"),
			}
			passports = append(passports, p)
			sb.Reset()
	}

	return passports
}


// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) string {
	for _, n := range a {
		if strings.Contains(n, x) {
			strA := strings.Split(n, ":")
			return strA[1]
		}
	}
	
	return ""
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

	passports := LoadPassports(curPath + "\\day4input.dat")

	var totalValidPassports int
	for _, p := range passports {
		if (p.IsValidPassport()) {
			totalValidPassports++
		}
	}

	fmt.Println("Total Valid Passports %d", totalValidPassports)

	
}