package main

import (
	"fmt"
)

func drawBoard(board [3][3]string) {
	fmt.Println("  0 1 2")
	for i, line := range board {
		fmt.Print(i)
		for _, cell := range line {
			fmt.Printf(" %s", cell)
		}
		fmt.Println()
	}
}

func checkMarch(board [3][3]string, icon string) bool {
	if (board[0][0] == icon && board[0][1] == icon && board[0][2] == icon) ||
		(board[1][0] == icon && board[1][1] == icon && board[1][2] == icon) ||
		(board[2][0] == icon && board[2][1] == icon && board[2][2] == icon) ||
		(board[0][0] == icon && board[1][0] == icon && board[2][0] == icon) ||
		(board[0][1] == icon && board[1][1] == icon && board[2][1] == icon) ||
		(board[0][2] == icon && board[1][2] == icon && board[2][2] == icon) ||
		(board[0][0] == icon && board[1][1] == icon && board[2][2] == icon) ||
		(board[0][2] == icon && board[1][1] == icon && board[2][1] == icon) {

		return true
	} else {
		return false
	}
}

func main() {
	board := [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	icon := [2]string{"○", "☓"}
	drawBoard(board)

	var line int
	var column int

	/*
		0: First attack
		1: Second attack
	*/
	var turn = 0

	fmt.Print("First: ")

	for {
		n, err := fmt.Scan(&line, &column)
		fmt.Println()

		if err != nil || n > 3 || line > 3 || column > 3 || board[line][column] != "-" {
			fmt.Println("Please input number from 0 to 3.")
			continue
		}

		if turn == 0 {
			board[line][column] = icon[turn]
			drawBoard(board)

			if checkMarch(board, icon[turn]) {
				fmt.Println("Winner First!")
				break
			} else {
				turn = 1
				fmt.Print("Second: ")
			}

		} else {
			board[line][column] = icon[turn]
			drawBoard(board)

			if checkMarch(board, icon[turn]) {
				fmt.Println("Winner Second!")
				break
			} else {
				turn = 0
				fmt.Print("First: ")
			}
		}
	}
}
