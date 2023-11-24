package main

import (
	"errors"
	"fmt"
)

type Board struct {
	positions []string
	numAvailablePositions int
}

// Print the board in the console
func (B *Board) drawBoard() {
	fmt.Printf(" %s | %s | %s \n", B.positions[7], B.positions[8], B.positions[9])
	fmt.Println("   |   |   ")
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", B.positions[4], B.positions[5], B.positions[6])
	fmt.Println("   |   |   ")
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", B.positions[1], B.positions[2], B.positions[3])
	fmt.Println("   |   |   ")
	fmt.Println()
}

// Inser the character in the position selected by the player or the machine
func (B *Board) insertChar(pos int, char string) error {
	// Insert the character in the selected position
	if r := B.isFree(pos); !r {
		return errors.New("invalid position!")	
	}

	B.positions[pos] = char
	B.numAvailablePositions --
	return nil

}

// Check if a specific position is available
func (B *Board) isFree(pos int) bool{
	// Check if the position is available
	return B.positions[pos] == " "
}

// Create a copy of the current board state to test possible moves
func (B *Board) duplicateBoard() []string {
	copiedBoard := make([]string, len(B.positions))
	copy(copiedBoard, B.positions)
	return copiedBoard
}

// Get a list of all the current available position in the board
func (B *Board) availablePositions() []int {
	empty := []int{}
	for i := 1; i < len(B.positions); i++ {
		if result := B.isFree(i); result {
			empty = append(empty, i)
		}
	}

	return empty
}

// Check if the player or the machine has won
func IsWinner(board []string, char string) bool {
	// Consider all winning cases
	//Horizontal
	if board[1] == char && board[2] == char && board[3] == char {
		return true
	} else if board[4] == char && board[5] == char && board[6] == char {
		return true
	} else if board[7] == char && board[8] == char && board[9] == char {
		return true
	}

	// Vertical
	if board[1] == char && board[4] == char && board[7] == char {
		return true
	} else if board[2] == char && board[5] == char && board[8] == char {
		return true
	} else if board[3] == char && board[6] == char && board[9] == char {
		return true
	}

	// Diagonal
	if board[1] == char && board[5] == char && board[9] == char {
		return true
	} else if board[7] == char && board[5] == char && board[3] == char {
		return true
	}
	
	return false
}
