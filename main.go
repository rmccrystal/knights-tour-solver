package main

import (
	"fmt"
)

func main() {
	b := chess.CreateBoard(Position{8, 8}, Position{4, 4})
	b.Move(Moves[1])
	b.Move(Moves[1])
	fmt.Println(b)
}