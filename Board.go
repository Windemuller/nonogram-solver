package main

import (
	"fmt"
)

const emptyField = "☐"
const positiveField = "☑"
const negativeField = "☒"

//Board contains the board data
type Board struct {
	boardLength int
	board       [][]string
	valuesRows  [][]int
	valuesCols  [][]int
}

//Extracts the row values from the values.txt file using the filereader function
func makeRowValues(boardLength int) [][]int {
	rowValues := make([][]int, 0)
	for i := 1; i <= boardLength; i++ {
		rowValues = append(rowValues, fileReader(i, "row"))
	}
	return rowValues
}

//Extracts the column values from the values.txt file using the filereader function
func makeColValues(boardLength int) [][]int {
	colValues := make([][]int, 0)
	for i := 1; i <= boardLength; i++ {
		colValues = append(colValues, fileReader(i, "col"))
	}
	return colValues
}

//makes a new, empty nonogram board
func makeBoard(boardLength int) Board {
	board := make([][]string, boardLength)
	for i := 0; i < boardLength; i++ {
		row := make([]string, boardLength)
		for j := 0; j < boardLength; j++ {
			row[j] = emptyField
		}
		board[i] = row
	}
	return Board{boardLength, board, makeRowValues(boardLength), makeColValues(boardLength)}
}

//prints the board
func printBoard(board Board) {
	for i := 0; i < board.boardLength; i++ {
		line := make([]string, board.boardLength)
		for j := 0; j < board.boardLength; j++ {
			line[j] = board.board[i][j]
		}
		fmt.Println(line)
	}
}

//fills in a space with a given symbol
func fillInSpace(board *Board, symbol string, positionRow int, positionCol int) {
	board.board[positionRow][positionCol] = symbol
}

//fills in a whole row or column of fields with a symbol, depending on input
func fillInAllSpaces(board *Board, symbol string, numRowOrCol int, rowOrCol string) {
	if rowOrCol == "row" {
		for i := 0; i < board.boardLength; i++ {
			board.board[numRowOrCol][i] = symbol
		}
	} else {
		for i := 0; i < board.boardLength; i++ {
			board.board[i][numRowOrCol] = symbol
		}
	}
}

/*
solveEasyFields loops over the value slices of the board at the pointer, looks if there are easily solvable ones and fills them in
for each row/col the function firstly looks at the rows/cols where no fields have to be filled in,
secondly looks at the rows/cols where no fields have to be filled in,
finishes by looking at the rows/cols where the values, and the fields in between the values add up to the boardlength
*/
func solveEasyFields(board *Board) {
	for i := 0; i < board.boardLength; i++ {
		valuesRow := board.valuesRows[i]
		valuesCol := board.valuesCols[i]

		if len(valuesRow) == 0 { //fills in the empty rows
			fillInAllSpaces(board, negativeField, i, "row")
		} else if valuesRow[0] == board.boardLength { //fills in the fully filled rows
			fillInAllSpaces(board, positiveField, i, "row")
		}
		solveEasyFieldsStep3(board, valuesRow, "row", i)
		solveEasyFieldsStep4(board, valuesRow, "row", i)
		if len(valuesRow) == 2 && sumArray(valuesRow) > board.boardLength/2 {
			solveEasyFieldsStep5(board, valuesRow, "row", i)
		}

		if len(valuesCol) == 0 { //fills in the empty columns
			fillInAllSpaces(board, negativeField, i, "col")
		} else if valuesCol[0] == board.boardLength { //fills in the fully filled columns
			fillInAllSpaces(board, positiveField, i, "col")
		}
		solveEasyFieldsStep3(board, valuesCol, "col", i)
		solveEasyFieldsStep4(board, valuesCol, "col", i)
		if len(valuesCol) == 2 && sumArray(valuesCol) > board.boardLength/2 {
			solveEasyFieldsStep5(board, valuesCol, "col", i)
		}
	}
}

/*
Step 3 of solveEasyFields looks at the rows where the combined value of the row or column restraints combined with the number of them,
minus 1 is equal to the boardlength. This makes the whole row or column fillable.
*/
func solveEasyFieldsStep3(board *Board, values []int, rowOrCol string, rowOrColNumber int) {
	if sumArray(values)+len(values)-1 == board.boardLength && rowOrCol == "col" {
		index := 0
		for _, element := range values {
			for k := 0; k < element; k++ { //for every element in values fills in that many "positive fields"
				fillInSpace(board, positiveField, index, rowOrColNumber)
				index++
			}
			if index != board.boardLength-1 { //adds a negative field after each element if the indexed field is not off the board
				fillInSpace(board, negativeField, index, rowOrColNumber)
				index++
			}
		}
	} else if sumArray(values)+len(values)-1 == board.boardLength {
		index := 0
		for _, element := range values {
			for k := 0; k < element; k++ { //for every element in values fills in that many "positive fields"
				fillInSpace(board, positiveField, rowOrColNumber, index)
				index++
			}
			if index != board.boardLength-1 { //adds a negative field after each element if the indexed field is not off the board
				fillInSpace(board, negativeField, rowOrColNumber, index)
				index++
			}
		}
	}
}

/*
step 4 of solveEasyFields looks at the rows and columns where the restraint value is more than half of the boardlength, this makes it so that by
logic some fields in that row or column are solvable
*/
func solveEasyFieldsStep4(board *Board, values []int, rowOrCol string, rowOrColNumber int) {
	if values[0] >= board.boardLength/2+1 && rowOrCol == "col" {
		startIndex := board.boardLength - 1 - values[0]
		endIndex := 0 + values[0]
		for index := startIndex; index <= endIndex; index++ {
			fillInSpace(board, positiveField, index, rowOrColNumber)
		}

	} else if values[0] >= board.boardLength/2+1 {
		startIndex := board.boardLength - 1 - values[0]
		endIndex := 0 + values[0]
		for index := startIndex; index <= endIndex; index++ {
			fillInSpace(board, positiveField, rowOrColNumber, index)
		}
	}
}

/*
step 5 of solveEasyFields looks at the rows and columns with 2 values where some fields would be positive in all situations and fills them in
by making a map of all possible indexes in a row or column. The function then makes a comparison map for each possible combination
of constraints and blocks, and deletes the values in the possible map that are not in the comparisonmap.
*/
func solveEasyFieldsStep5(board *Board, values []int, rowOrCol string, rowOrColNumber int) {
	possibleFields := make(map[int]bool)
	lastFieldIndex := 0
	for index := range possibleFields {
		possibleFields[index] = true
	}
	for startIndex := 0; startIndex+sumArray(values)+1 <= board.boardLength; startIndex++ {
		for gap := 1; lastFieldIndex != board.boardLength-1; gap++ {
			compareMap := make(map[int]bool)
			for i := 0; i < values[0]; i++ {
				compareMap[startIndex+i] = true
			}
			for i := 0; i < values[1]; i++ {
				compareMap[startIndex+values[0]+gap+i] = true
				lastFieldIndex = startIndex + values[0] + gap + i
			}
			for k := range possibleFields {
				if _, ok := compareMap[k]; ok == false {
					delete(possibleFields, k)
				}
			}
		}
	}
	for k := range possibleFields {
		if rowOrCol == "row" {
			fillInSpace(board, positiveField, rowOrColNumber, k)
		} else {
			fillInSpace(board, positiveField, k, rowOrColNumber)
		}
	}
}

func step1(board *Board) {

}

func step2(board *Board) {

}

func step3(board *Board) {

}

func solveBoard(board Board) {
	solveEasyFields(&board)
	for isBoardSolved(&board) == false {

	}
}

//Checks if the board is solved by looking if all the fields have been filled in
func isBoardSolved(board *Board) bool {
	for i := 0; i < board.boardLength; i++ {
		for j := 0; j < board.boardLength; j++ {
			if board.board[i][j] == emptyField {
				return false
			}
		}
	}
	return true
}

//creates a sum of all values in an integer array
func sumArray(array []int) int {
	output := 0
	for _, element := range array {
		output += element
	}
	return output
}

func main() {
}
