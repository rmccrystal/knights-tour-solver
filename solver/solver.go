package solver

import (
	"../chess"
	"errors"
	"fmt"
)

func SolveBoard(b chess.Board) {
	for {
		bestMove, isSolved, err := FindBestMove(b)
		if isSolved {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		b.Move(bestMove)
		fmt.Println(b)
	}
	fmt.Println("Solved Board: \n%s", b.String())
}

// This function will find the best move by using the move with the least amount of subsequent moves
func FindBestMove(b chess.Board) (bestMove chess.Move, isSolved bool, err error) {
	if b.IsSolved() {
		return chess.Move{}, true, nil
	}
	possibleCurrentMoves := b.PossibleMoves()
	if len(possibleCurrentMoves) == 0 {
		return chess.Move{}, false, errors.New("No more available moves")		// If there are no possible moves return nothing
	}
	return FindMoveWithLeastSubsequentMoves(b), false, nil
}

// TODO: b.PossibleMoves will be called twice here
func FindMoveWithLeastSubsequentMoves(b chess.Board) chess.Move {
	var currentBestMoves []chess.Move
	var lowestSubsequentMoves int = 10
	for _, move := range b.PossibleMoves() {
		tempBoard := b
		tempBoard.Move(move)
		if tempBoard.IsSolved() {
			return move
		}
		subsequentMoves := tempBoard.PossibleMoves()
		if len(subsequentMoves) == 0 {
			continue
		}
		if len(subsequentMoves) < lowestSubsequentMoves {
			currentBestMoves = nil		// Clear array, we found a new best
			currentBestMoves = append(currentBestMoves, move)
			lowestSubsequentMoves = len(subsequentMoves)
		} else if len(subsequentMoves) == lowestSubsequentMoves {
			currentBestMoves = append(currentBestMoves, move)
		}
	}
	return currentBestMoves[0]
}