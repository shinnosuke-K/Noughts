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

/* Frank way
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

*/

// Improvement point
func countCell(board [3][3]string, icon string, x, y, dx, dy, count int) int {
	//fmt.Println(x, y)
	if x >= 0 && x <= 2 && y >= 0 && y <= 2 && board[x][y] == icon {
		count++
		return countCell(board, icon, x+dx, y+dy, dx, dy, count)
	} else {
		return count
	}
}

// Improvement point
func checkMatch(board [3][3]string, icon string, line, column, turn int) bool {
	dr := [4][2]int{{-1, -1}, {0, -1}, {-1, 0}, {-1, 1}}
	dl := [4][2]int{{1, 1}, {0, 1}, {1, 0}, {1, -1}}

	for n := 0; n < 4; n++ {
		count := 0
		count = countCell(board, icon, line+dr[n][0], column+dr[n][1], dr[n][0], dr[n][1], count) +
			countCell(board, icon, line+dl[n][0], column+dl[n][1], dl[n][0], dl[n][1], count)
		if count == 2 {
			switch turn {
			case 0:
				fmt.Println("Winner First!")
				return true
			case 1:
				fmt.Println("Winner Second!")
				return true
			}
		}
	}
	return false
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

		if err != nil || n >= 3 || line >= 3 || column >= 3 || board[line][column] != "-" {
			fmt.Println("Please input number from 0 to 3.")
			continue
		}

		if turn == 0 {
			board[line][column] = icon[turn]
			drawBoard(board)

			if checkMatch(board, icon[turn], line, column, turn) {
				break
			} else {
				turn = 1
				fmt.Print("Second: ")
			}

		} else {
			board[line][column] = icon[turn]
			drawBoard(board)

			if checkMatch(board, icon[turn], line, column, turn) {
				break
			} else {
				turn = 0
				fmt.Print("First: ")
			}
		}
	}
}
