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

func makeRowValues(boardLength int) [][]int {
	rowValues := make([][]int, 0)
	for i := 1; i <= boardLength; i++ {
		rowValues = append(rowValues, fileReader(i, "row"))
	}
	return rowValues
}

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

func fillInSpace(board *Board, symbol string, positionRow int, positionCol int) {
	board.board[positionRow][positionCol] = symbol
}

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
		} else if sumArray(valuesRow)+len(valuesRow)-1 == board.boardLength {
			index := 0
			for _, element := range valuesRow {
				for k := 0; k < element; k++ { //for every element in valuesRow fills in that many "positive fields"
					fillInSpace(board, positiveField, i, index)
					index++
				}
				if index != board.boardLength-1 { //adds a negative field after each element if the indexed field is not off the board
					fillInSpace(board, negativeField, i, index)
					index++
				}
			}
		}

		if len(valuesCol) == 0 { //fills in the empty columns
			fillInAllSpaces(board, negativeField, i, "col")
		} else if valuesCol[0] == board.boardLength { //fills in the fully filled columns
			fillInAllSpaces(board, positiveField, i, "col")
		} else if sumArray(valuesCol)+len(valuesCol)-1 == board.boardLength {
			index := 0
			for _, element := range valuesCol {
				for k := 0; k < element; k++ { //for every element in valuesCol fills in that many "positive fields"
					fillInSpace(board, positiveField, index, i)
					index++
				}
				if index != board.boardLength-1 { //adds a negative field after each element if the indexed field is not off the board
					fillInSpace(board, negativeField, index, i)
					index++
				}
			}
		} else if valuesCol[0] >= board.boardLength/2 {
			//index
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

func sumArray(array []int) int {
	output := 0
	for _, element := range array {
		output += element
	}
	return output
}

func main() {
}
