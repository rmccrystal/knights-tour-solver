package main

import (
	"fmt"
	"strings"
)

var Moves = []Move{
	Move{1, 2},
	Move{2, 1},
	Move{2, -1},
	Move{1, -2},
	Move{-1, -2},
	Move{-2, -1},
	Move{-2, 1},
	Move{-1, 2},
}

// Starts at 1
type Position struct {
	X int
	Y int
}

type Move struct {
	HorizontalDistance int
	VerticalDistance   int
}

type Board struct {
	Size             Position
	PlayerPosition   Position
	PastPositions    []Position
	StartingPosition Position
}

func CreateBoard(size Position, startingLocation Position) Board {
	return Board{
		Size:             size,
		PlayerPosition:   startingLocation,
		StartingPosition: startingLocation,
		PastPositions:    []Position{startingLocation},
	}
}

func (b Board) CanMove(move Move) bool {
	// The new location after moving
	moveLocation := Position{b.PlayerPosition.X + move.HorizontalDistance, b.PlayerPosition.Y + move.VerticalDistance}

	if (moveLocation.X > b.Size.X || moveLocation.Y > b.Size.Y) || (moveLocation.X < 1 || moveLocation.Y < 1) {
		// Move is not in the board space
		return false
	}

	// Iterate through all past locations and check if the new location equals that
	for _, pos := range b.PastPositions {
		if pos == moveLocation {
			return false
		}
	}
	return true
}

func (b Board) PossibleMoves() []Move {
	var moves []Move
	for _, move := range Moves {
		if b.CanMove(move) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (b *Board) Move(m Move) bool {
	if !b.CanMove(m) {
		return false
	}
	b.PlayerPosition = Position{b.PlayerPosition.X + m.VerticalDistance, b.PlayerPosition.Y + m.HorizontalDistance}
	b.PastPositions = append(b.PastPositions, b.PlayerPosition)
	return true
}

func (b Board) String() string {
	boardStr := ""
	boardStr = boardStr + strings.Repeat("-", (b.Size.X * 2) + 1) + "\n|"
	var positions [][]int
	for n, pos := range b.PastPositions {

	}
	return boardStr
}

func main() {
	b := CreateBoard(Position{8, 8}, Position{4, 4})
	b.Move(Moves[1])
	b.Move(Moves[1])
	fmt.Println(b)
}
