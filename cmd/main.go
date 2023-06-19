package main

import (
	"chess-board-backend/m/v2/board"
	"fmt"
)

func main() {
	brd := board.InitBoard()
	board.PrintBoard(brd)

	fmt.Println("hello chess")
}
