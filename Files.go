package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "Values.txt"
const rowLine = "Enter the values of the rows instead of the blanks, each seperated by a semicolon(;)"
const colLine = "\nEnter the values of the columns instead of the blanks, each seperated by a semicolon(;)"

//creates a new values.txt file if none exists, otherwise overwrites the existing one
func createValuesFile() {
	f, err := os.Create(fileName)
	errorHandler(err)
	f.Close()
}

func formatValuesFile(boardLength int) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644) //opening the values file with writing and reading permission
	errorHandler(err)
	_, err = f.WriteString(rowLine + "\n") //writing the colLine constant to the file
	errorHandler(err)
	for i := 0; i < boardLength; i++ { //writing the row number and value brackets to the line
		j := "row " + strconv.Itoa(i+1) + " {" + strings.Repeat("_;", lengthCalc(boardLength)-1) + "_}\n"
		_, err := f.WriteString(j)
		errorHandler(err)
	}
	_, err = f.WriteString(colLine + "\n") //writing the colLine constant to the file
	errorHandler(err)
	for i := 0; i < boardLength; i++ { //writing the col number and value brackets to the line
		j := "col " + strconv.Itoa(i+1) + " {" + strings.Repeat("_;", lengthCalc(boardLength)-1) + "_}\n"
		_, err := f.WriteString(j)
		errorHandler(err)
	}
	f.Close()
}

//handles the errors of the os functions to maintain readability
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//calcs max number of possible filled fields on the board
func lengthCalc(boardLength int) int {
	maxSize := 0
	if boardLength%2 == 0 {
		maxSize = boardLength / 2
	} else {
		maxSize = boardLength/2 + 1
	}
	return maxSize
}

//FileReader reads a line in the file to form a slice of the passed values
func fileReader(line int, rowOrCol string) []int {
	searchFor := rowOrCol + " " + strconv.Itoa(line) + " "
	file, err := os.Open(fileName)
	errorHandler(err)
	sc := bufio.NewScanner(file)
	sc.Scan()
	for true {
		sc.Scan()
		s := sc.Text()
		if strings.Contains(s, searchFor) {
			break
		}
	}

	intermediateResult := sc.Text()
	intermediateResult = strings.Replace(intermediateResult, searchFor, "", 1)
	intermediateResult = strings.Trim(intermediateResult, "{}")
	resultAsStringSlice := strings.Split(intermediateResult, ";")
	resultAsIntSlice := stringSliceToIntSlice(resultAsStringSlice)
	return resultAsIntSlice
}

//converts all characters of a string slice to an int slice
func stringSliceToIntSlice(inputs []string) []int {
	output := make([]int, 0)
	for _, input := range inputs {
		if a, err := strconv.Atoi(input); err == nil {
			output = append(output, a)
		}
	}
	return output
}
