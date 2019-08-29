package main

import (
	"./chess"
	"./solver"
)

func main() {
	b := chess.CreateBoard(chess.Position{50, 50}, chess.Position{1, 1})
	solver.SolveBoard(b)
}