package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
	"regexp"
	"strconv"
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

func (p Passport) IsValidPassportStrict() (bool) {
	var validBirthYear = p.BirthYear != "" && (ConvertToInt(p.BirthYear) >= 1920 && ConvertToInt(p.BirthYear) < 2003)
	var validExpYear = p.ExpirationYear != "" && (ConvertToInt(p.ExpirationYear) >= 2020 && ConvertToInt(p.ExpirationYear) < 2031)
	var validIssueYear = p.IssueYear != "" && (ConvertToInt(p.IssueYear) >= 2010 && ConvertToInt(p.IssueYear) < 2021)
	var validHieght = p.Height != "" && (RegexMatch("^(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)$", p.Height))
	var validHairColor = p.HairColor != "" && (RegexMatch("^#[0-9a-f]{6}$", p.HairColor))
	var validEyeColor = p.EyeColor != "" && (Contains([]string { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" }, p.EyeColor))
	var validPassPortId = p.PassportId != "" && (RegexMatch("^[0-9]{9}$" , p.PassportId));

	return validBirthYear && validExpYear && validIssueYear && validHieght && validHairColor && validEyeColor && validPassPortId;
}

func ConvertToInt(value string) (int) {
	intVal, _ := strconv.Atoi(value)
	return intVal
}

func RegexMatch(pattern string, value string) (bool) {
	match, _ := regexp.MatchString(pattern, value)
	return match
}

func Contains(strArray []string, value string) (bool) {
	for _, item := range strArray {
		if (item == value) {
			return true
		}
	}

	return false
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
	fmt.Println("Part 1: -----\n\r")
	var totalValidPassports int
	var totalValidStrictPassports int
	for _, p := range passports {
		if (p.IsValidPassport()) {
			totalValidPassports++
		}

		if (p.IsValidPassportStrict()) {
			totalValidStrictPassports++
		}
	}

	fmt.Println("Total Valid Passports %d", totalValidPassports)
	fmt.Println("\n\rPart 2: -----\n\r")
	fmt.Println("Total Valid Strict Passports %d", totalValidStrictPassports)

	
}

