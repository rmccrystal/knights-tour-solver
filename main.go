package main

import (
	"./chess"
	"./solver"
)

func main() {
	b := chess.CreateBoard(chess.Position{8, 8}, chess.Position{1, 1})
	solver.SolveBoard(b)
}