package main

import (
	"fmt"
	"strings"
)

func main() {
	game := [][]string{
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
	}

	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "X"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}
