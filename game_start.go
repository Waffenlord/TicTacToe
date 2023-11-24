package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


type Player struct {
	char string
}

const option1 = "X"
const option2 = "O"

func Start() (Player, Player) {
	scanner := bufio.NewScanner(os.Stdin)

	player := Player{}
	cpu := Player{}
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("=======================")

	for {

		fmt.Println("Please select x or o to start playing (type x or o)...")
		scanner.Scan()

		choice := scanner.Text()
		cleanchoice, err := cleanChoice(choice)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if cleanchoice == option1 {
			player.char = cleanchoice
			cpu.char = option2
		} else {
			player.char = cleanchoice
			cpu.char = option1
		}

		return player, cpu
	}
}

// Validate user input
func cleanChoice(choice string) (string, error) {
	choiceCapitalize := strings.ToUpper(choice)
	options := []string{option1, option2}
	for _, option := range options {
		if choiceCapitalize == option {
			return choiceCapitalize, nil
		}
	}

	return "", errors.New("invalid option!")
}