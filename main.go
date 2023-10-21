package main

import (
	"chess/chess/board"
	"chess/chess/color"
	"fmt"
)

func main() {
	fmt.Println(board.New().View(color.Black))
}
