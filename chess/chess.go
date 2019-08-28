package chess

import (
	"fmt"
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
	b.PlayerPosition = Position{b.PlayerPosition.X + m.HorizontalDistance, b.PlayerPosition.Y + m.VerticalDistance}
	b.PastPositions = append(b.PastPositions, b.PlayerPosition)
	return true
}

func (b Board) IsSolved() bool {
	return len(b.PastPositions) == b.Size.X * b.Size.Y
}

func (b Board) String() string {
	// Create a positions array
	positions := make([][]int, b.Size.X)		// position[x][y]
	for i := 0; i < b.Size.X; i++ {
		positions[i] = make([]int, b.Size.Y)
	}

	// Set each respective array element to the number
	for n, pos := range b.PastPositions {
		positions[b.Size.Y - pos.Y][pos.X - 1] = n+1 // Set the element to n+1, which is the move number
	}

	boardStr := ""

	for _, i := range positions {
		boardStr += "\n|"
		for _, j := range i {
			if j == 0 {		// If the element is empty
				boardStr += "  "
			} else {
				if j <= 9 {			// Buffer the number appropriately	// TODO: Make this work for higher numbers
					boardStr += fmt.Sprintf("%d ", j)
				} else {
					boardStr += fmt.Sprintf("%d", j)
				}
			}
			boardStr += "|"
		}
	}

	return boardStr
}