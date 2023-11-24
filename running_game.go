package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Turn(B *Board, p Player, cpu Player) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose your move:")
		B.drawBoard()
		scanner.Scan()
		positionText := scanner.Text()
		intPosition, err := validatePosition(positionText)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		B.insertChar(intPosition, p.char)
		playerWon := IsWinner(B.positions, p.char)
		if playerWon {
			B.drawBoard()
			fmt.Println("You win!")
			break
		}

		if B.numAvailablePositions == 0 {
			B.drawBoard()
			fmt.Println("It's a tie")
			break
		}

		cpuPos := cpuTurn(B, cpu, p)
		B.insertChar(cpuPos, cpu.char)
		fmt.Println("Machine moves: ", cpuPos)
		cpuWon := IsWinner(B.positions, cpu.char)
		if cpuWon {
			B.drawBoard()
			fmt.Println("You lose!")
			break
		}

		if B.numAvailablePositions == 0 {
			B.drawBoard()
			fmt.Println("It's a tie")
			break
		}

	}

	restart(scanner)
	
}

func cpuTurn(B *Board, cpu Player, p Player) int {
	availableMoves := B.availablePositions()

	// Block player from winning
	for _, option := range availableMoves {
		copiedBoard := B.duplicateBoard()
		copiedBoard[option] = p.char
		if IsWinner(copiedBoard, p.char) {
			return option
		}
	}

	// Look for the winning move
	for _, pos := range availableMoves {
		copiedBoard := B.duplicateBoard()
		copiedBoard[pos] = cpu.char
		if IsWinner(copiedBoard, cpu.char) {
			return pos
		}
	}

	middleAndCorners := []int{1, 3, 5, 7, 9}
	// Look for the middle and corners
	for _, move := range availableMoves{
		for _, place := range middleAndCorners {
			if move == place {
				return move
			}
		}
	}

	// Select randomly any available position
	randomIndex := rand.Intn(len(availableMoves))

	return availableMoves[randomIndex]

}

func validatePosition(pos string) (int, error) {
	intPos, err := strconv.Atoi(pos)
	if err != nil || intPos < 1 || intPos > 9 {
		return 0, errors.New("invalid position!")
	}

	return intPos, nil
}

func validateChoice(choice string) (bool, error) {
	lowerChoice := strings.ToLower(choice)
	if lowerChoice == "yes" {
		return true, nil
	} else if lowerChoice == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid option!")
	}
}

func restart(scanner *bufio.Scanner) {
	for {
		fmt.Println("Play again? (yes or no)")
		scanner.Scan()
		choice := scanner.Text()
		result, err := validateChoice(choice)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if result {
			break
		} else {
			os.Exit(0)
		}
	}
}