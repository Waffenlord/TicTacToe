package main



func main() {
	// Main loop
	for {

		// Extract the structs containing the player and machine characters
		player, cpu := Start()

		// Initialize the empty board
		board := Board{
			positions: []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			numAvailablePositions: 9,
		}

		Turn(&board, player, cpu)
	}
}